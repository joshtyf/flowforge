package execute

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"

	"github.com/gookit/event"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/events"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExecutionManager struct {
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

func NewStepExecutionManager(mongoClient *mongo.Client, psqlClient *sql.DB, configs ...ExecutionManagerConfig) (*ExecutionManager, error) {
	if mongoClient == nil {
		return nil, fmt.Errorf("mongo client is nil")
	}
	if psqlClient == nil {
		return nil, fmt.Errorf("psql client is nil")
	}
	srm := &ExecutionManager{
		executors:   map[models.PipelineStepType]*stepExecutor{},
		mongoClient: mongoClient,
		psqlClient:  psqlClient,
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
	logger.Info("[ServiceRequestManager] Handling new service request event", nil)
	serviceRequest := e.(*events.NewServiceRequestEvent).ServiceRequest()
	if serviceRequest == nil {
		logger.Error("[ServiceRequestManager] Service request is nil", nil)
		return fmt.Errorf("service request is nil")
	}
	// Fetch the pipeline so that we know what steps to execute
	pipeline, err := database.NewPipeline(srm.mongoClient).GetById(serviceRequest.PipelineId)
	if err != nil {
		logger.Error("[ServiceRequestManager] Error getting pipeline", map[string]interface{}{"err": err})
	}

	// Get the first step and its executor
	firstStep := pipeline.GetPipelineStep(pipeline.FirstStepName)
	if firstStep == nil {
		logger.Error("[ServiceRequestManager] No first step found", map[string]interface{}{"step": pipeline.FirstStepName})
		return fmt.Errorf("no first step found")
	}
	currExecutor := srm.executors[firstStep.StepType]
	if currExecutor == nil {
		// TODO: Handle error
		logger.Error("[ServiceRequestManager] No executor found for first step", map[string]interface{}{"step": firstStep.StepName})
		return fmt.Errorf("no executor found for first step")
	}

	// Update the service request status to running
	err = database.NewServiceRequest(srm.mongoClient).UpdateStatus(serviceRequest.Id.Hex(), models.Running)
	if err != nil {
		logger.Error("[ServiceRequestManager] Error updating service request status", map[string]interface{}{"err": err})
		return err
	}

	// Create log directory
	log_dir := fmt.Sprintf("%s/%s", logger.BaseLogDir, serviceRequest.Id.Hex())
	err = os.MkdirAll(log_dir, 0755)
	if err != nil {
		logger.Error("[ServiceRequestManager] Error creating log folder", map[string]interface{}{"err": err})
		return err
	}

	err = srm.execute(serviceRequest, firstStep, currExecutor)
	return err
}

func (srm *ExecutionManager) execute(serviceRequest *models.ServiceRequestModel, step *models.PipelineStepModel, executor *stepExecutor) error {
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
		EventType:        models.STEP_STARTED,
		ServiceRequestId: serviceRequest.Id.Hex(),
		StepName:         step.StepName,
	})
	if err != nil {
		logger.Error("[ServiceRequestManager] Error creating service request event", map[string]interface{}{"err": err})
		return err
	}

	// Create a log file for the current step
	// TODO: file is not persisted in Docker container
	f, err := os.OpenFile(
		fmt.Sprintf("%s/%s/%s.log", logger.BaseLogDir, serviceRequest.Id.Hex(), step.StepName),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		logger.Error("[ServiceRequestManager] Error opening log file", map[string]interface{}{"err": err})
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			logger.Error("[ServiceRequestManager] Error closing log file", map[string]interface{}{"err": err})
		}
	}()
	executor_logger := logger.NewExecutorLogger(io.MultiWriter(os.Stdout, f), step.StepName)

	// Execute the current step
	_, err = (*executor).execute(executeCtx, executor_logger)
	if err != nil {
		logger.Error("[ServiceRequestManager] Error executing step", map[string]interface{}{"step": (*executor).getStepType(), "err": err})
		// TODO: Handle error
		return err
	}

	return nil
}

func (srm *ExecutionManager) handleCompletedStepEvent(e event.Event) error {
	logger.Info("[ServiceRequestManager] Handling step completed event", nil)
	completedStepEvent := e.(*events.StepCompletedEvent)
	completedStep := completedStepEvent.CompletedStep()
	if completedStep == nil {
		logger.Error("[ServiceRequestManager] Completed step is nil", nil)
		return fmt.Errorf("completed step is nil")
	}
	serviceRequest := completedStepEvent.ServiceRequest()
	if serviceRequest == nil {
		logger.Error("[ServiceRequestManager] Service request is nil", nil)
		return fmt.Errorf("service request is nil")
	}

	// Log step completed event
	serviceRequestEvent := database.NewServiceRequestEvent(srm.psqlClient)
	err := serviceRequestEvent.Create(&models.ServiceRequestEventModel{
		EventType:        models.STEP_COMPLETED,
		ServiceRequestId: serviceRequest.Id.Hex(),
		StepName:         completedStep.StepName,
	})
	if err != nil {
		// TODO: not sure if we should return here. We need to handle the error better
		logger.Error("[ServiceRequestManager] Error creating service request event", map[string]interface{}{"err": err})
		return err
	}

	if completedStep.IsTerminalStep {
		err := database.NewServiceRequest(srm.mongoClient).UpdateStatus(serviceRequest.Id.Hex(), models.Success)
		if err != nil {
			// TODO: Handle error
			// Need to ensure idempotency or figure out a rollback solution
			logger.Error("[ServiceRequestManager] Error updating service request status", map[string]interface{}{"err": err})
		}
		return nil
	}
	pipeline, err := database.NewPipeline(srm.mongoClient).GetById(serviceRequest.PipelineId)
	if err != nil {
		logger.Error("[ServiceRequestManager] Error getting pipeline", map[string]interface{}{"err": err})
		return err
	}

	// Set the current executor to the next executor
	nextStep := pipeline.GetPipelineStep(completedStep.NextStepName)
	if nextStep == nil {
		logger.Error("[ServiceRequestManager] No next step found", map[string]interface{}{"step": completedStep.NextStepName})
		return fmt.Errorf("no next step found")
	}
	nextExecutor := srm.executors[nextStep.StepType]
	if nextExecutor == nil {
		// TODO: Handle error
		logger.Error("[ServiceRequestManager] No executor found for next step", map[string]interface{}{"step": nextStep.StepName})
		return fmt.Errorf("no executor found for next step")
	}

	srm.execute(serviceRequest, nextStep, nextExecutor)
	return nil
}
