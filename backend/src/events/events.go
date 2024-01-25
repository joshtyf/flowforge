package events

import (
	"github.com/gookit/event"
	"github.com/joshtyf/flowforge/src/database/models"
)

const (
	NewServiceRequestEventName = "NewServiceRequestEvent"
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
