package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceRequestStatus string

const (
	Pending    ServiceRequestStatus = "Pending"
	Running    ServiceRequestStatus = "Running"
	Success    ServiceRequestStatus = "Success"
	Failure    ServiceRequestStatus = "Failure"
	Canceled   ServiceRequestStatus = "Canceled"
	NotStarted ServiceRequestStatus = "Not Started"
)

type FormData map[string]any

type ServiceRequestModel struct {
	Id              primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	PipelineId      string               `bson:"pipeline_id" json:"pipeline_id"` // should we use primitive.ObjectID here?
	PipelineVersion int                  `bson:"pipeline_version" json:"pipeline_version"`
	Status          ServiceRequestStatus `bson:"status" json:"status"`
	CreatedOn       time.Time            `bson:"created_on" json:"created_on"`
	LastUpdated     time.Time            `bson:"last_updated" json:"last_updated"`
	Remarks         string               `bson:"remarks" json:"remarks"`
	FormData        FormData             `bson:"form_data" json:"form_data"`
}
