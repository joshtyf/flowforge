package database

import (
	"context"
	"fmt"

	"github.com/joshtyf/flowforge/src/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceRequest struct {
	c *mongo.Client
}

func NewServiceRequest(c *mongo.Client) *ServiceRequest {
	return &ServiceRequest{c: c}
}

func (sr *ServiceRequest) Create(srm *models.ServiceRequestModel) (*mongo.InsertOneResult, error) {
	res, err := sr.c.Database(DatabaseName).Collection("service_requests").InsertOne(context.Background(), srm)
	return res, err
}

func (sr *ServiceRequest) GetById(id string) (*models.ServiceRequestModel, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	result := sr.c.Database(DatabaseName).Collection("service_requests").FindOne(context.Background(), bson.M{"_id": objectId})
	if result.Err() != nil {
		return nil, result.Err()
	}
	srm := &models.ServiceRequestModel{}
	result.Decode(srm)
	return srm, nil
}

func (sr *ServiceRequest) UpdateById(id string, srm *models.ServiceRequestModel) (*mongo.UpdateResult, error) {
	// TODO: update the method once form data is finalised
	// objectId, _ := primitive.ObjectIDFromHex(id)
	// filter := bson.M{"_id": objectId}
	// update := bson.M{"$set": bson.M{
	// 	"pipeline_id":      srm.PipelineId,
	// 	"pipeline_version": srm.PipelineVersion,
	// 	"last_updated":     srm.LastUpdated,
	// 	"remarks":          srm.Remarks}}

	// res, err := sr.c.Database(DatabaseName).Collection("service_requests").UpdateOne(context.Background(), filter, update)
	// if err != nil {
	// 	return nil, err
	// }
	// return res, nil

	return nil, nil
}

func (sr *ServiceRequest) GetAll(pg Pagination) ([]*models.ServiceRequestModel, error) {
	result, err := sr.c.Database(DatabaseName).Collection("service_requests").Aggregate(
		context.TODO(),
		mongo.Pipeline{},
	)
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

func (sr *ServiceRequest) UpdateStatus(id string, status models.ServiceRequestStatus) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = sr.c.Database(DatabaseName).Collection("service_requests").UpdateOne(
		context.Background(), bson.M{"_id": objectId}, bson.M{"$set": bson.M{"status": status}})
	return err
}

type GetServiceRequestFilters struct {
	UserId   string
	Statuses []string
}

type GetAllServiceRequestByOrgResponse struct {
	Data       []*models.ServiceRequestModel
	TotalCount int
}

func (sr *ServiceRequest) GetAllServiceRequestByOrg(orgId int, filters GetServiceRequestFilters, pg Pagination) (*GetAllServiceRequestByOrgResponse, error) {
	query := bson.M{"org_id": orgId}
	if len(filters.Statuses) > 0 {
		query["status"] = bson.M{"$in": filters.Statuses}
	}
	if filters.UserId != "" {
		query["user_id"] = filters.UserId
	}
	result, err := sr.c.Database(DatabaseName).Collection("service_requests").Aggregate(
		context.Background(),
		mongo.Pipeline{
			{{Key: "$match", Value: query}},
			{{Key: "$facet", Value: bson.D{
				{
					Key:   "totalCount",
					Value: bson.A{bson.D{{Key: "$count", Value: "total"}}},
				},
				{
					Key: "data",
					Value: bson.A{
						bson.D{{Key: "$skip", Value: (pg.Page - 1) * pg.PageSize}},
						bson.D{{Key: "$limit", Value: pg.PageSize}},
					},
				},
			}}},
		},
	)
	if err != nil {
		return nil, err
	}
	defer result.Close(context.Background())

	type dataResp struct {
		Data       []*models.ServiceRequestModel `bson:"data"`
		TotalCount []struct {
			Total int `bson:"total"`
		} `bson:"totalCount"`
	}
	srms := []*models.ServiceRequestModel{}
	var totalCount int
	for result.Next(context.Background()) {
		data := &dataResp{}
		err := result.Decode(data)
		if err != nil {
			return nil, fmt.Errorf("error decoding data: %w", err)
		}
		totalCount = data.TotalCount[0].Total
		srms = append(srms, data.Data...)
	}
	return &GetAllServiceRequestByOrgResponse{
		Data:       srms,
		TotalCount: totalCount,
	}, nil
}
