package handlers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HttpError struct {
	Error string
}

type FormData struct {
	UserId           string   `bson:"user_id"`
	UserName         string   `bson:"username"`
	AdditionalFields []string `bson:"additional_fields"`
}

type ServiceRequest struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	PipelineId      string             `bson:"pipeline_id"`
	PipelineVersion float32            `bson:"pipeline_version"`
	Status          string             `bson:"status"`
	CreatedOn       time.Time          `bson:"created_on"`
	LastUpdated     time.Time          `bson:"last_updated"`
	Remarks         string             `bson:"remarks"`
	FormData        FormData           `bson:"form_data"`
}

func NewServiceRequest() *ServiceRequest {
	var time = time.Now()
	return &ServiceRequest{CreatedOn: time, LastUpdated: time, FormData: FormData{}}
}

func NewHttpError(err error) *HttpError {
	return &HttpError{Error: err.Error()}
}
