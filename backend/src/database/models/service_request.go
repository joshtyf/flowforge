package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceRequestModel struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	PipelineId      string             `bson:"pipeline_id"`
	PipelineVersion int                `bson:"pipeline_version"`
	Status          string             `bson:"status"`
	CreatedOn       time.Time          `bson:"created_on"`
	LastUpdated     time.Time          `bson:"last_updated"`
	Remarks         string             `bson:"remarks"`
	// FormData        FormData           `bson:"form_data"`
}
