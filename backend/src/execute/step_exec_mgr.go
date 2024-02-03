package execute

import (
	"context"
	"fmt"

	"github.com/gookit/event"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/events"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExecutionManager struct {
	mongoClient *mongo.Client
	executors   map[models.PipelineStepType]*stepExecutor
}

type ExecutionManagerConfig func(*ExecutionManager)

func WithStepExecutor(step stepExecutor) ExecutionManagerConfig {
	return func(srm *ExecutionManager) {
		srm.executors[step.getStepType()] = &step
	}
}

func NewStepExecutionManager(configs ...ExecutionManagerConfig) *ExecutionManager {
	mongoClient, err := client.GetMongoClient()
	if err != nil {
		logger.Error("[ServiceRequestManager] Error getting mongo client", map[string]interface{}{"err": err})
	}
	srm := &ExecutionManager{
		executors:   map[models.PipelineStepType]*stepExecutor{},
		mongoClient: mongoClient,
	}
	for _, c := range configs {
		c(srm)
	}
	return srm
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

	err = database.NewServiceRequest(srm.mongoClient).UpdateStatus(serviceRequest.Id.Hex(), models.Running)
	if err != nil {
		logger.Error("[ServiceRequestManager] Error updating service request status", map[string]interface{}{"err": err})
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
	// Execute the current step
	_, err := (*executor).execute(executeCtx)
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
