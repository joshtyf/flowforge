package models

import "time"

type EventType string

const (
	STEP_NOT_STARTED EventType = "Not Started"
	STEP_RUNNING     EventType = "Running"
	STEP_FAILED      EventType = "Failed"
	STEP_CANCELLED   EventType = "Cancelled"
	STEP_COMPLETED   EventType = "Completed"
)

type ServiceRequestEventModel struct {
	EventId          int
	EventType        EventType
	ServiceRequestId string
	StepName         string
	ApprovedBy       string
	CreatedAt        time.Time
}
