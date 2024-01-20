package database

import (
	"context"

	"github.com/joshtyf/flowforge/src/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceRequest struct {
	c *mongo.Client
}

func NewServiceRequest(c *mongo.Client) *ServiceRequest {
	return &ServiceRequest{c: c}
}

func (sr *ServiceRequest) Create(srm *models.ServiceRequestModel) error {
	_, err := c.Database(DatabaseName).Collection("service_requests").InsertOne(context.Background(), srm)
	return err
}

func (sr *ServiceRequest) GetById(id string) (*models.ServiceRequestModel, error) {
	result := c.Database(DatabaseName).Collection("service_requests").FindOne(context.Background(), bson.M{"_id": id})
	if result.Err() != nil {
		return nil, result.Err()
	}
	srm := &models.ServiceRequestModel{}
	result.Decode(srm)
	return srm, nil
}

func (sr *ServiceRequest) GetAll() ([]*models.ServiceRequestModel, error) {
	result, err := c.Database(DatabaseName).Collection("service_requests").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	srms := []*models.ServiceRequestModel{}
	for result.Next(context.Background()) {
		srm := &models.ServiceRequestModel{}
		result.Decode(srm)
		srms = append(srms, srm)
	}
	return srms, nil
}
