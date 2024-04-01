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

	pipelineIdInHex := "8A7F3EBCD951246A5F0E9B87"
	pipelineId, err := primitive.ObjectIDFromHex(pipelineIdInHex)
	if err != nil {
		panic(err)
	}
	pipeline := models.PipelineModel{
		PipelineName:  "Test Pipeline",
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
					"url":    "https://example.com",
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
		Form: map[string]interface{}{
			"input": models.FormInput{
				Title:       "Input",
				Description: "Input Description with minimum length 1",
				Type:        "input",
				Required:    true,
				MinLength:   1,
				Placeholder: "Input placeholder...",
			},
			"select": models.FormSelect{
				Title:       "Select Option",
				Description: "Dropdown selection with default value as Item 1",
				Type:        "select",
				Required:    true,
				Options:     []string{"Item 1", "Item 2", "Item 3"},
				Placeholder: "Select placeholder...",
			},
			"checkbox": models.FormCheckboxes{
				Title: "Checkboxes",
				Description: "You can select more than 1 item",
				Type: "checkboxes",
				Required:    false,
				Options: []string{"Item 1", "Item 2", "Item 3"},
			},
		},
	}
	res, err := database.NewPipeline(c).Create(&pipeline)
	if err != nil {
		panic(err)
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		logger.Info("Inserted pipeline", map[string]interface{}{"id": oid.String()})
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
		PipelineId:      pipelineIdInHex,
		PipelineVersion: 1,
		Status:          models.Success,
		OrganisationId:  1,
		Remarks:         "This is a test service request.",
		CreatedOn:       time.Date(2024, time.January, 1, 1, 0, 0, 0, time.UTC),
		LastUpdated:     time.Date(2024, time.January, 1, 1, 0, 0, 0, time.UTC),
	}

	res, err = database.NewServiceRequest(c).Create(&serviceRequest)
	if err != nil {
		panic(err)
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		logger.Info("Inserted service request", map[string]interface{}{"id": oid.String()})
	} else {
		panic("Inserted ID is not an ObjectID")
	}
}
