package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceRequestStatus string

const (
	Pending ServiceRequestStatus = "Pending"
	Running ServiceRequestStatus = "Running"
	Success ServiceRequestStatus = "Success"
	Failure ServiceRequestStatus = "Failure"
)

type ServiceRequestModel struct {
	Id              primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	PipelineId      string               `bson:"pipeline_uuid" json:"pipeline_uuid"`
	PipelineVersion int                  `bson:"pipeline_version" json:"pipeline_version"`
	Status          ServiceRequestStatus `bson:"status" json:"status"`
	CreatedOn       time.Time            `bson:"created_on" json:"created_on"`
	LastUpdated     time.Time            `bson:"last_updated" json:"last_updated"`
	Remarks         string               `bson:"remarks" json:"remarks"`
	// FormData        FormData           `bson:"form_data"`
}
