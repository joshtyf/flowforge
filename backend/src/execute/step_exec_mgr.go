package execute

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/gookit/event"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/events"
	"github.com/joshtyf/flowforge/src/helper"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExecutionManager struct {
	logger      logger.ServerLogger
	mongoClient *mongo.Client
	psqlClient  *sql.DB
	executors   map[models.PipelineStepType]*stepExecutor
}

type ExecutionManagerConfig func(*ExecutionManager)

func WithStepExecutor(step stepExecutor) ExecutionManagerConfig {
	return func(srm *ExecutionManager) {
		srm.executors[step.getStepType()] = &step
	}
}

func NewStepExecutionManager(mongoClient *mongo.Client, psqlClient *sql.DB, logger logger.ServerLogger, configs ...ExecutionManagerConfig) (*ExecutionManager, error) {
	if mongoClient == nil {
		return nil, fmt.Errorf("mongo client is nil")
	}
	if psqlClient == nil {
		return nil, fmt.Errorf("psql client is nil")
	}
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}
	srm := &ExecutionManager{
		executors:   map[models.PipelineStepType]*stepExecutor{},
		mongoClient: mongoClient,
		psqlClient:  psqlClient,
		logger:      logger,
	}
	for _, c := range configs {
		c(srm)
	}
	return srm, nil
}

// Starts the manager by registering event listeners
func (srm *ExecutionManager) Start() {
	event.On(events.NewServiceRequestEventName, event.ListenerFunc(srm.handleNewServiceRequestEvent), event.Normal)
	event.On(events.StepCompletedEventName, event.ListenerFunc(srm.handleCompletedStepEvent), event.Normal)
}

func (srm *ExecutionManager) handleNewServiceRequestEvent(e event.Event) error {
	srm.logger.Info("handling service request event")
	serviceRequest := e.(*events.NewServiceRequestEvent).ServiceRequest()
	if serviceRequest == nil {
		srm.logger.Error(fmt.Sprintf("event %s missing data: %s", e.Name(), "service request"))
		return fmt.Errorf("service request is nil")
	}
	// Fetch the pipeline so that we know what steps to execute
	pipeline, err := database.NewPipeline(srm.mongoClient).GetById(serviceRequest.PipelineId)
	if err != nil {
		srm.logger.Error(fmt.Sprintf("error encountered while handling event: %s", err))
	}

	// Get the first step and its executor
	firstStep := pipeline.GetPipelineStep(pipeline.FirstStepName)
	if firstStep == nil {
		srm.logger.Error(fmt.Sprintf("missing pipeline step: %s", pipeline.FirstStepName))
		return fmt.Errorf("no first step found")
	}
	currExecutor := srm.executors[firstStep.StepType]
	if currExecutor == nil {
		// TODO: Handle error
		srm.logger.Error(fmt.Sprintf("missing executor for step: %s", firstStep.StepName))
		return fmt.Errorf("no executor found for first step")
	}

	// Update the service request status to running
	err = database.NewServiceRequest(srm.mongoClient).UpdateStatus(serviceRequest.Id.Hex(), models.RUNNING)
	if err != nil {
		srm.logger.Error(fmt.Sprintf("failed to run service request %s: %s", serviceRequest.Id.Hex(), err))
		return err
	}

	// Create log directory
	err = logger.CreateExecutorLogDir(serviceRequest.Id.Hex())
	if err != nil {
		srm.logger.Error(fmt.Sprintf("error encountered while handling event: %s", err))
		return err
	}

	err = srm.execute(serviceRequest, firstStep, currExecutor)
	return err
}

func (srm *ExecutionManager) execute(serviceRequest *models.ServiceRequestModel, step *models.PipelineStepModel, executor *stepExecutor) error {
	// Parse and replace step parameters with service request form data
	for key, val := range step.Parameters {
		replaced, err := helper.ReplacePlaceholders(val, serviceRequest.FormData)
		if err != nil {
			srm.logger.Error(fmt.Sprintf("unable to replace placeholder based on form data for %s", key))
			return err
		}
		step.Parameters[key] = replaced
	}

	// Create an execution context with the current step and service request
	executeCtx := context.WithValue(
		context.WithValue(
			context.Background(),
			util.ServiceRequestKey,
			serviceRequest),
		util.StepKey,
		step,
	)
	// Log step started event
	serviceRequestEvent := database.NewServiceRequestEvent(srm.psqlClient)
	err := serviceRequestEvent.Create(&models.ServiceRequestEventModel{
		EventType:        models.STEP_RUNNING,
		ServiceRequestId: serviceRequest.Id.Hex(),
		StepName:         step.StepName,
		StepType:         step.StepType,
	})
	if err != nil {
		srm.logger.Error(fmt.Sprintf("error encountered while handling event: %s", err))
		return err
	}

	// Create a log file for the current step
	// TODO: file is not persisted in Docker container
	f, err := os.OpenFile(
		logger.CreateExecutorLogFilePath(serviceRequest.Id.Hex(), step.StepName),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		srm.logger.Error(fmt.Sprintf("error encountered while handling event: %s", err))
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			srm.logger.Error(fmt.Sprintf("error encountered while handling event: %s", err))
		}
	}()
	executor_logger := logger.NewExecutorLogger(io.MultiWriter(os.Stdout, f), step.StepName)

	// Execute the current step
	_, err = (*executor).execute(executeCtx, executor_logger)
	if err != nil {
		srm.logger.Error(fmt.Sprintf("error encountered while executing step %s: %s", step.StepName, err))
		// TODO: Handle error
		return err
	}

	return nil
}

func (srm *ExecutionManager) handleCompletedStepEvent(e event.Event) error {
	srm.logger.Info("handling step completed event")
	completedStepEvent := e.(*events.StepCompletedEvent)
	completedStep := completedStepEvent.CompletedStep()
	if completedStep == "" {
		srm.logger.Error(fmt.Sprintf("event %s missing data: %s", e.Name(), "completed step"))
		return fmt.Errorf("completed step is not provided")
	}
	serviceRequest := completedStepEvent.ServiceRequest()
	if serviceRequest == nil {
		srm.logger.Error(fmt.Sprintf("event %s missing data: %s", e.Name(), "service request"))
		return fmt.Errorf("service request is nil")
	}

	pipeline, err := database.NewPipeline(srm.mongoClient).GetById(serviceRequest.PipelineId)
	if errors.Is(err, mongo.ErrNoDocuments) {
		srm.logger.Error(fmt.Sprintf("pipeline not found: %s", serviceRequest.PipelineId))
		return err
	}
	if err != nil {
		srm.logger.Error(fmt.Sprintf("error encountered while handling event: %s", err))
		return err
	}
	completedStepModel := pipeline.GetPipelineStep(completedStep)
	if completedStepModel.IsTerminalStep {
		err := database.NewServiceRequest(srm.mongoClient).UpdateStatus(serviceRequest.Id.Hex(), models.COMPLETED)
		if err != nil {
			// TODO: Handle error
			// Need to ensure idempotency or figure out a rollback solution
			srm.logger.Error(fmt.Sprintf("failed to mark service request %s successful: %s", serviceRequest.Id.Hex(), err))
		}
		return nil
	}

	// Log step completed event
	serviceRequestEvent := database.NewServiceRequestEvent(srm.psqlClient)
	err = serviceRequestEvent.Create(&models.ServiceRequestEventModel{
		EventType:        models.STEP_COMPLETED,
		ServiceRequestId: serviceRequest.Id.Hex(),
		StepName:         completedStep,
		StepType:         completedStepModel.StepType,
	})
	if err != nil {
		// TODO: not sure if we should return here. We need to handle the error better
		srm.logger.Error(fmt.Sprintf("error encountered while handling event: %s", err))
		return err
	}

	// Set the current executor to the next executor
	nextStep := pipeline.GetPipelineStep(completedStepModel.NextStepName)
	if nextStep == nil {
		srm.logger.Error(fmt.Sprintf("missing pipeline step: %s", completedStepModel.NextStepName))
		return fmt.Errorf("no next step found")
	}
	nextExecutor := srm.executors[nextStep.StepType]
	if nextExecutor == nil {
		// TODO: Handle error
		srm.logger.Error(fmt.Sprintf("missing executor for step: %s", nextStep.StepName))
		return fmt.Errorf("no executor found for next step")
	}

	srm.execute(serviceRequest, nextStep, nextExecutor)
	return nil
}
