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

func (sre *ServiceRequestEvent) GetStepsLatestEvent(serviceReequestId string) ([]*models.ServiceRequestEventModel, error) {
	queryStr := `
		WITH LatestEvents AS (
			SELECT *, ROW_NUMBER() OVER (PARTITION BY step_name ORDER BY created_at DESC) AS row_num
			FROM service_request_event
			WHERE service_request_id = $1
		)
		SELECT event_id, event_type, service_request_id, step_name, approved_by, created_at FROM LatestEvents
		WHERE row_num = 1;`

	rows, err := sre.db.Query(queryStr, serviceReequestId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var srems []*models.ServiceRequestEventModel
	for rows.Next() {
		srem := &models.ServiceRequestEventModel{}
		err := rows.Scan(
			&srem.EventId,
			&srem.EventType,
			&srem.ServiceRequestId,
			&srem.StepName,
			&srem.ApprovedBy,
			&srem.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		srems = append(srems, srem)
	}
	return srems, nil
}
