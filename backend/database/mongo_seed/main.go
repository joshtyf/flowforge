package main

import (
	"os"
	"time"

	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	logger := logger.NewServerLog(os.Stdout)
	c, err := client.GetMongoClient()
	if err != nil {
		panic(err)
	}

	pipelineIdInHex := "8A7F3EBCD951246A5F0E9B87"
	pipelineName := "Test Pipeline"
	pipelineId, err := primitive.ObjectIDFromHex(pipelineIdInHex)
	if err != nil {
		panic(err)
	}
	pipeline := models.PipelineModel{
		PipelineName:  pipelineName,
		Id:            pipelineId,
		Version:       1,
		FirstStepName: "step1",
		Steps: []models.PipelineStepModel{
			{
				StepName:     "step1",
				StepType:     models.APIStep,
				NextStepName: "step2",
				PrevStepName: "",
				Parameters: map[string]string{
					"method": "GET",
					"url":    "https://example.com?param=${param}",
				},
				IsTerminalStep: false,
			},
			{
				StepName:       "step2",
				StepType:       models.WaitForApprovalStep,
				NextStepName:   "",
				PrevStepName:   "step1",
				Parameters:     map[string]string{},
				IsTerminalStep: true,
			},
		},
		Form: models.Form{Fields: []models.FormField{
			{
				Name:        "field1",
				Title:       "Field 1",
				Type:        models.InputField,
				Required:    true,
				Placeholder: "Enter text...",
				MinLength:   1,
			},
			{
				Name:        "field2",
				Title:       "Field 2",
				Type:        models.SelectField,
				Required:    true,
				Placeholder: "Select an option",
				Options:     []string{"Option 1", "Option 2", "Option 3"},
				Default:     "Option 1",
			},
			{
				Name:    "field3",
				Title:   "Field 3",
				Type:    models.CheckboxField,
				Options: []string{"Option 1", "Option 2", "Option 3"},
			},
		}},
	}

	res, err := database.NewPipeline(c).Create(&pipeline)
	if err != nil {
		panic(err)
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		logger.Info("created pipeline with ID: " + oid.String())
	} else {
		panic("Inserted ID is not an ObjectID")
	}

	serviceReqIdInHex := "F2D8E1A73B964C5E7A0F81D9"
	serviceReqId, err := primitive.ObjectIDFromHex(serviceReqIdInHex)
	if err != nil {
		panic(err)
	}
	serviceRequest := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|65e9dabff2dab546ed0c231e", // josh's user ID
		PipelineId:      pipelineIdInHex,
		PipelineName:    pipelineName,
		PipelineVersion: 1,
		Status:          models.NOT_STARTED,
		OrganizationId:  1,
		Remarks:         "This is a test service request.",
		CreatedOn:       time.Date(2024, time.January, 1, 1, 0, 0, 0, time.UTC),
		LastUpdated:     time.Date(2024, time.January, 1, 1, 0, 0, 0, time.UTC),
		FormData: models.FormData{
			"param": "test_param",
		},
	}

	res, err = database.NewServiceRequest(c).Create(&serviceRequest)
	if err != nil {
		panic(err)
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		logger.Info("created service request with ID: " + oid.String())
	} else {
		panic("Inserted ID is not an ObjectID")
	}
}
