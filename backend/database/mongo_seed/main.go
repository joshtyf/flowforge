package main

import (
	"time"

	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	c, err := client.GetMongoClient()
	if err != nil {
		panic(err)
	}

	serviceReqId, err := primitive.ObjectIDFromHex("F2D8E1A73B964C5E7A0F81D9")
	if err != nil {
		panic(err)
	}
	serviceRequest := models.ServiceRequestModel{
		Id:              serviceReqId,
		PipelineId:      "1",
		PipelineVersion: 1,
		Status:          models.Success,
		Remarks:         "This is a test service request.",
		CreatedOn:       time.Date(2024, time.January, 1, 1, 0, 0, 0, time.UTC),
		LastUpdated:     time.Date(2024, time.January, 1, 1, 0, 0, 0, time.UTC),
	}

	res, err := database.NewServiceRequest(c).Create(&serviceRequest)
	if err != nil {
		panic(err)
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		logger.Info("Inserted service request", map[string]interface{}{"id": oid.String()})
	} else {
		panic("Inserted ID is not an ObjectID")
	}
	pipelineUuid, err := primitive.ObjectIDFromHex("8A7F3EBCD951246A5F0E9B87")
	if err != nil {
		panic(err)
	}
	pipeline := models.PipelineModel{
		PipelineName:  "Test Pipeline",
		Id:            pipelineUuid,
		Version:       1,
		FirstStepName: "step1",
		Steps: []models.PipelineStepModel{
			{
				StepName:     "step1",
				StepType:     models.APIStep,
				NextStepName: "",
				PrevStepName: "",
				Parameters: map[string]string{
					"method": "GET",
					"url":    "https://example.com",
				},
				IsTerminalStep: true,
			},
		},
	}

	res, err = database.NewPipeline(c).Create(&pipeline)
	if err != nil {
		panic(err)
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		logger.Info("Inserted pipeline", map[string]interface{}{"id": oid.String()})
	} else {
		panic("Inserted ID is not an ObjectID")
	}
}
