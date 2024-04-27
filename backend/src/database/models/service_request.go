package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceRequestStatus string

const (
	NOT_STARTED ServiceRequestStatus = "Not Started"
	RUNNING     ServiceRequestStatus = "Running"
	FAILED      ServiceRequestStatus = "Failed"
	CANCELLED   ServiceRequestStatus = "Cancelled"
	COMPLETED   ServiceRequestStatus = "Completed"
	// NOTE: update allServiceRequestStatuses when adding new status
)

var allServiceRequestStatuses = []ServiceRequestStatus{NOT_STARTED, RUNNING, FAILED, CANCELLED, COMPLETED}

func ValidateServiceRequestStatus(status string) bool {
	for _, s := range allServiceRequestStatuses {
		if s == ServiceRequestStatus(status) {
			return true
		}
	}
	return false
}

type FormData map[string]any

type ServiceRequestModel struct {
	Id              primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	UserId          string               `bson:"user_id" json:"user_id"`
	OrganizationId  int                  `bson:"org_id" json:"org_id"`
	PipelineId      string               `bson:"pipeline_id" json:"pipeline_id"` // should we use primitive.ObjectID here?
	PipelineName    string               `bson:"pipeline_name" json:"pipeline_name"`
	PipelineVersion int                  `bson:"pipeline_version" json:"pipeline_version"`
	Status          ServiceRequestStatus `bson:"status" json:"status"`
	CreatedOn       time.Time            `bson:"created_on" json:"created_on"`
	LastUpdated     time.Time            `bson:"last_updated" json:"last_updated"`
	Remarks         string               `bson:"remarks" json:"remarks"`
	FormData        FormData             `bson:"form_data" json:"form_data"`
}
