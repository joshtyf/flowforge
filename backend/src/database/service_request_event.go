package database

import (
	"database/sql"

	"github.com/joshtyf/flowforge/src/database/models"
)

type ServiceRequestEvent struct {
	db *sql.DB
}

func NewServiceRequestEvent(db *sql.DB) *ServiceRequestEvent {
	return &ServiceRequestEvent{
		db: db,
	}
}

func (sre *ServiceRequestEvent) Create(srem *models.ServiceRequestEventModel) error {
	queryStr := "INSERT INTO service_request_event (event_type, service_request_id, step_name, approved_by) VALUES ($1, $2, $3, $4)"
	_, err := sre.db.Exec(
		queryStr,
		srem.EventType,
		srem.ServiceRequestId,
		srem.StepName,
		srem.ApprovedBy,
	)
	return err
}
