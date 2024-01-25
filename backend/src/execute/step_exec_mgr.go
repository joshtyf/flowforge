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
)

type ExecutionManager struct {
	executors map[models.PipelineStepType]*stepExecutor
}

type ExecutionManagerConfig func(*ExecutionManager)

func WithStepExecutor(step stepExecutor) ExecutionManagerConfig {
	return func(srm *ExecutionManager) {
		srm.executors[step.getStepType()] = &step
	}
}

func NewStepExecutionManager(configs ...ExecutionManagerConfig) *ExecutionManager {
	srm := &ExecutionManager{
		executors: map[models.PipelineStepType]*stepExecutor{},
	}
	for _, c := range configs {
		c(srm)
	}
	return srm
}

// Starts the manager by registering event listeners
func (srm *ExecutionManager) Start() {
	event.On(events.NewServiceRequestEventName, event.ListenerFunc(srm.handleNewServiceRequestEvent), event.Normal)
}

func (srm *ExecutionManager) handleNewServiceRequestEvent(e event.Event) error {
	logger.Info("[ServiceRequestManager] Handling new service request event", nil)
	req := e.(*events.NewServiceRequestEvent).ServiceRequest()
	err := srm.execute(req)
	return err
}

func (srm *ExecutionManager) execute(serviceRequest *models.ServiceRequestModel) error {
	if serviceRequest == nil {
		logger.Error("[ServiceRequestManager] Service request is nil", nil)
		return fmt.Errorf("service request is nil")
	}
	mongoClient, err := client.GetMongoClient()
	if err != nil {
		logger.Error("[ServiceRequestManager] Error getting mongo client", map[string]interface{}{"err": err})
	}
	// Fetch the pipeline so that we know what steps to execute
	pipeline, err := database.NewPipeline(mongoClient).GetById(serviceRequest.PipelineId)
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
	currStep := firstStep

	err = database.NewServiceRequest(mongoClient).UpdateStatus(serviceRequest.Id.Hex(), models.Running)
	if err != nil {
		logger.Error("[ServiceRequestManager] Error updating service request status", map[string]interface{}{"err": err})
		return err
	}

	// Execute the pipeline step by step
	for {

		// Create an execution context with the current step and service request
		executeCtx := context.WithValue(
			context.WithValue(
				context.Background(),
				util.ServiceRequestKey,
				serviceRequest),
			util.StepKey,
			currStep,
		)
		// Execute the current step
		_, err := (*currExecutor).execute(executeCtx)
		if err != nil {
			logger.Error("[ServiceRequestManager] Error executing step", map[string]interface{}{"step": (*currExecutor).getStepType(), "err": err})
			// TODO: Handle error
			return err
		}

		if currStep.IsTerminalStep {
			// If the current step is a terminal step, we're done
			break
		} else {
			// Set the current executor to the next executor
			nextStep := pipeline.GetPipelineStep(currStep.NextStepName)
			if nextStep == nil {
				logger.Error("[ServiceRequestManager] No next step found", map[string]interface{}{"step": currStep.NextStepName})
				return fmt.Errorf("no next step found")
			}
			nextExecutor := srm.executors[nextStep.StepType]
			if nextExecutor == nil {
				// TODO: Handle error
				logger.Error("[ServiceRequestManager] No executor found for next step", map[string]interface{}{"step": nextStep.StepName})
				return fmt.Errorf("no executor found for next step")
			}
			currExecutor = nextExecutor
			currStep = nextStep
		}
	}

	err = database.NewServiceRequest(mongoClient).UpdateStatus(serviceRequest.Id.Hex(), models.Success)
	if err != nil {
		// TODO: Handle error
		// Need to ensure idempotency or figure out a rollback solution
		logger.Error("[ServiceRequestManager] Error updating service request status", map[string]interface{}{"err": err})
	}

	return nil
}
