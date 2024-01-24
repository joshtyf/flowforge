package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceRequestModel struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	PipelineId      string             `bson:"pipeline_id" json:"pipeline_id"`
	PipelineVersion int                `bson:"pipeline_version" json:"pipeline_version"`
	Status          Status             `bson:"status" json:"status"`
	CreatedOn       time.Time          `bson:"created_on" json:"created_on"`
	LastUpdated     time.Time          `bson:"last_updated" json:"last_updated"`
	Remarks         string             `bson:"remarks" json:"remarks"`
	// FormData        FormData           `bson:"form_data"`
}
