package models

import "time"

type EventType string

const (
	STEP_STARTED   EventType = "STEP_STARTED"
	STEP_APPROVED  EventType = "STEP_APPROVED"
	STEP_REJECTED  EventType = "STEP_REJECTED"
	STEP_COMPLETED EventType = "STEP_COMPLETED"
)

type ServiceRequestEventModel struct {
	EventId          int
	EventType        EventType
	ServiceRequestId string
	StepName         string
	ApprovedBy       string
	CreatedAt        time.Time
}
