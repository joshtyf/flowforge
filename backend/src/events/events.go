package events

import (
	"github.com/gookit/event"
	"github.com/joshtyf/flowforge/src/database/models"
)

const (
	NewServiceRequestEventName = "NewServiceRequestEvent"
	StepCompletedEventName     = "StepCompletedEvent"
	StepFailedEventName        = "StepFailedEvent"
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
	completedStep    string
	serviceRequestId string
	createdBy        string
	results          interface{}
	err              error
}

func NewStepCompletedEvent(completedStep string, serviceRequestId string, createdBy string, results interface{}, err error) *StepCompletedEvent {
	e := &StepCompletedEvent{
		completedStep:    completedStep,
		serviceRequestId: serviceRequestId,
		createdBy:        createdBy,
		results:          results,
		err:              err,
	}
	e.SetName(StepCompletedEventName)
	return e
}

func (e *StepCompletedEvent) CompletedStep() string {
	return e.completedStep
}

func (e *StepCompletedEvent) ServiceRequestId() string {
	return e.serviceRequestId
}

func (e *StepCompletedEvent) CreatedBy() string {
	return e.createdBy
}

func (e *StepCompletedEvent) Results() interface{} {
	return e.results
}

func (e *StepCompletedEvent) Err() error {
	return e.err
}

type StepFailedEvent struct {
	event.BasicEvent
	failedStep     string
	serviceRequest *models.ServiceRequestModel
	createdBy      string
	remarks        string
	err            error
}

func NewStepFailedEvent(failedStep string, serviceRequest *models.ServiceRequestModel, createdBy string, remarks string, err error) *StepFailedEvent {
	e := &StepFailedEvent{
		failedStep:     failedStep,
		serviceRequest: serviceRequest,
		createdBy:      createdBy,
		remarks:        remarks,
		err:            err,
	}
	e.SetName(StepFailedEventName)
	return e
}

func (e *StepFailedEvent) FailedStep() string {
	return e.failedStep
}

func (e *StepFailedEvent) ServiceRequest() *models.ServiceRequestModel {
	return e.serviceRequest
}

func (e *StepFailedEvent) CreatedBy() string {
	return e.createdBy
}

func (e *StepFailedEvent) Remarks() string {
	return e.remarks
}

func (e *StepFailedEvent) Err() error {
	return e.err
}
