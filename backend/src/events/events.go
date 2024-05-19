package events

import (
	"github.com/gookit/event"
	"github.com/joshtyf/flowforge/src/database/models"
)

const (
	NewServiceRequestEventName = "NewServiceRequestEvent"
	StepCompletedEventName     = "StepCompletedEvent"
)

type NewServiceRequestEvent struct {
	event.BasicEvent
	serviceRequest *models.ServiceRequestModel
}

func NewNewServiceRequestEvent(serviceRequest *models.ServiceRequestModel) *NewServiceRequestEvent {
	e := &NewServiceRequestEvent{
		serviceRequest: serviceRequest,
	}
	e.SetName(NewServiceRequestEventName)
	return e
}

func (e *NewServiceRequestEvent) ServiceRequest() *models.ServiceRequestModel {
	return e.serviceRequest
}

type StepCompletedEvent struct {
	event.BasicEvent
	completedStep  string
	serviceRequest *models.ServiceRequestModel
	results        interface{}
	err            error
}

func NewStepCompletedEvent(completedStep string, serviceRequest *models.ServiceRequestModel, results interface{}, err error) *StepCompletedEvent {
	e := &StepCompletedEvent{
		completedStep:  completedStep,
		serviceRequest: serviceRequest,
		results:        results,
		err:            err,
	}
	e.SetName(StepCompletedEventName)
	return e
}

func (e *StepCompletedEvent) CompletedStep() string {
	return e.completedStep
}

func (e *StepCompletedEvent) ServiceRequest() *models.ServiceRequestModel {
	return e.serviceRequest
}

func (e *StepCompletedEvent) Results() interface{} {
	return e.results
}

func (e *StepCompletedEvent) Err() error {
	return e.err
}
