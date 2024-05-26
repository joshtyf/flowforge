package main

import (
	"os"
	"time"

	"github.com/gookit/event"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/events"
	"github.com/joshtyf/flowforge/src/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	logger := logger.NewServerLog(os.Stdout)
	c, err := client.GetMongoClient()
	if err != nil {
		panic(err)
	}

	// pipeline 1
	pipelineIdInHex := "8A7F3EBCD951246A5F0E9B87"
	pipelineName := "Test Pipeline 1"
	pipelineId, err := primitive.ObjectIDFromHex(pipelineIdInHex)
	if err != nil {
		panic(err)
	}
	pipeline1 := models.PipelineModel{
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
					"method": "${method}",
					"url":    "https://httpbin.org/${method}?param=${param}",
					"data":   "{\"key\": \"${value}\", \"${key}\": \"hardcoded_value\"}",
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

	// pipeline 2
	pipelineIdInHex = "7168233D6CED6EC585F3E205"
	pipelineName = "Test Pipeline 2"
	pipelineId, err = primitive.ObjectIDFromHex(pipelineIdInHex)
	if err != nil {
		panic(err)
	}
	pipeline2 := models.PipelineModel{
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
				StepName:     "step2",
				StepType:     models.APIStep,
				NextStepName: "step3",
				PrevStepName: "step1",
				Parameters: map[string]string{
					"method": "GET",
					"url":    "https://example.com?param=${param}",
				},
				IsTerminalStep: false,
			},
			{
				StepName:     "step3",
				StepType:     models.APIStep,
				NextStepName: "",
				PrevStepName: "step2",
				Parameters: map[string]string{
					"method": "GET",
					"url":    "https://example.com?param=${param}",
				},
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

	// pipeline 3
	pipelineIdInHex = "E03C8592CCE73820B56C1F14"
	pipelineName = "Test Pipeline 3"
	pipelineId, err = primitive.ObjectIDFromHex(pipelineIdInHex)
	if err != nil {
		panic(err)
	}
	pipeline3 := models.PipelineModel{
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
				NextStepName:   "step3",
				PrevStepName:   "step1",
				Parameters:     map[string]string{},
				IsTerminalStep: false,
			},
			{
				StepName:     "step3",
				StepType:     models.APIStep,
				NextStepName: "",
				PrevStepName: "step2",
				Parameters: map[string]string{
					"method": "GET",
					"url":    "https://example.com?param=${param}",
				},
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

	// pipeline 4
	pipelineIdInHex = "C6A3978A8723D64C29FEC42E"
	pipelineName = "Test Pipeline 4"
	pipelineId, err = primitive.ObjectIDFromHex(pipelineIdInHex)
	if err != nil {
		panic(err)
	}
	pipeline4 := models.PipelineModel{
		PipelineName:  pipelineName,
		Id:            pipelineId,
		Version:       1,
		FirstStepName: "step1",
		Steps: []models.PipelineStepModel{
			{
				StepName:       "step1",
				StepType:       models.WaitForApprovalStep,
				NextStepName:   "step2",
				PrevStepName:   "",
				Parameters:     map[string]string{},
				IsTerminalStep: false,
			},
			{
				StepName:     "step2",
				StepType:     models.APIStep,
				NextStepName: "",
				PrevStepName: "step1",
				Parameters: map[string]string{
					"method": "GET",
					"url":    "https://example.com?param=${param}",
				},
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

	pipelines := [...]models.PipelineModel{pipeline1, pipeline2, pipeline3, pipeline4}

	for i := 0; i < len(pipelines); i++ {
		res, err := database.NewPipeline(c).Create(&pipelines[i])
		if err != nil {
			panic(err)
		}
		if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
			logger.Info("created pipeline with ID: " + oid.String())
		} else {
			panic("Inserted ID is not an ObjectID")
		}
	}

	// service request 1
	serviceReqId, err := primitive.ObjectIDFromHex("F2D8E1A73B964C5E7A0F81D9")
	if err != nil {
		panic(err)
	}
	serviceRequest1 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|65e9dabff2dab546ed0c231e", // josh's user ID
		PipelineId:      pipeline1.Id.Hex(),
		PipelineName:    pipeline1.PipelineName,
		PipelineVersion: 1,
		Status:          models.NOT_STARTED,
		OrganizationId:  1,
		Remarks:         "This is a test service request.",
		CreatedOn:       time.Date(2024, time.January, 1, 1, 0, 0, 0, time.UTC),
		LastUpdated:     time.Date(2024, time.January, 1, 1, 0, 0, 0, time.UTC),
		FormData: models.FormData{
			"param":  "test_param",
			"method": "get",
			"key":    "test_key",
			"value":  "test_value",
		},
	}

	// service request 2
	serviceReqId, err = primitive.ObjectIDFromHex("662E134616B653509203CB93")
	if err != nil {
		panic(err)
	}
	serviceRequest2 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|65e9dabff2dab546ed0c231e", // josh's user ID
		PipelineId:      pipeline1.Id.Hex(),
		PipelineName:    pipeline1.PipelineName,
		PipelineVersion: 1,
		Status:          models.NOT_STARTED,
		OrganizationId:  1,
		Remarks:         "This is a test service request.",
		CreatedOn:       time.Date(2024, time.January, 1, 1, 0, 0, 0, time.UTC),
		LastUpdated:     time.Date(2024, time.January, 1, 1, 0, 0, 0, time.UTC),
		FormData: models.FormData{
			"param":  "test_param",
			"method": "post",
			"key":    "test_key",
			"value":  "test_value",
		},
	}

	// service request 3
	serviceReqId, err = primitive.ObjectIDFromHex("3192E86FDA27815A7E73DE4D")
	if err != nil {
		panic(err)
	}
	serviceRequest3 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|65e9dabff2dab546ed0c231e", // josh's user ID
		PipelineId:      pipeline2.Id.Hex(),
		PipelineName:    pipeline2.PipelineName,
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

	// service request 4
	serviceReqId, err = primitive.ObjectIDFromHex("0DE778EDE701480806798778")
	if err != nil {
		panic(err)
	}
	serviceRequest4 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|65e9dabff2dab546ed0c231e", // josh's user ID
		PipelineId:      pipeline2.Id.Hex(),
		PipelineName:    pipeline2.PipelineName,
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

	// service request 5
	serviceReqId, err = primitive.ObjectIDFromHex("A31A553999B9ACFED58F5C36")
	if err != nil {
		panic(err)
	}
	serviceRequest5 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|65e9dabff2dab546ed0c231e", // josh's user ID
		PipelineId:      pipeline3.Id.Hex(),
		PipelineName:    pipeline3.PipelineName,
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

	// service request 6
	serviceReqId, err = primitive.ObjectIDFromHex("74349FB2BB485BB06E4AE6D6")
	if err != nil {
		panic(err)
	}
	serviceRequest6 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|65e9dabff2dab546ed0c231e", // josh's user ID
		PipelineId:      pipeline4.Id.Hex(),
		PipelineName:    pipeline4.PipelineName,
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

	// service request 7
	serviceReqId, err = primitive.ObjectIDFromHex("9FEAAA1741E2DA3CA236DAC0")
	if err != nil {
		panic(err)
	}
	serviceRequest7 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|65ffab5c004e8d1620d06a64", // zqzheng.2019@smu.edu.sg
		PipelineId:      pipeline1.Id.Hex(),
		PipelineName:    pipeline1.PipelineName,
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

	// service request 8
	serviceReqId, err = primitive.ObjectIDFromHex("DBDB82098F6FD3B856EE3933")
	if err != nil {
		panic(err)
	}
	serviceRequest8 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|65ffab5c004e8d1620d06a64", // zqzheng.2019@smu.edu.sg
		PipelineId:      pipeline2.Id.Hex(),
		PipelineName:    pipeline2.PipelineName,
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

	// service request 9
	serviceReqId, err = primitive.ObjectIDFromHex("68F0F82F5B3F2432D51BD511")
	if err != nil {
		panic(err)
	}
	serviceRequest9 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|66010ad5095367b237799680", // zheng.zhiqiang49@gmail.com
		PipelineId:      pipeline3.Id.Hex(),
		PipelineName:    pipeline3.PipelineName,
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

	// service request 10
	serviceReqId, err = primitive.ObjectIDFromHex("583B94DDD4F3109A4140E617")
	if err != nil {
		panic(err)
	}
	serviceRequest10 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|66010ad5095367b237799680", // zheng.zhiqiang49@gmail.com
		PipelineId:      pipeline4.Id.Hex(),
		PipelineName:    pipeline4.PipelineName,
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

	// service request 11
	serviceReqId, err = primitive.ObjectIDFromHex("0F7BF7F3C4C5893CD90EF591")
	if err != nil {
		panic(err)
	}
	serviceRequest11 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|66010ad5095367b237799680", // zheng.zhiqiang49@gmail.com
		PipelineId:      pipeline1.Id.Hex(),
		PipelineName:    pipeline1.PipelineName,
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

	// service request 12
	serviceReqId, err = primitive.ObjectIDFromHex("1F8E7E3C9DFF02BE23161C4F")
	if err != nil {
		panic(err)
	}
	serviceRequest12 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|66010ad5095367b237799680", // zheng.zhiqiang49@gmail.com
		PipelineId:      pipeline2.Id.Hex(),
		PipelineName:    pipeline2.PipelineName,
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

	// service request 13
	serviceReqId, err = primitive.ObjectIDFromHex("558C137931D94A6BD686B7FD")
	if err != nil {
		panic(err)
	}
	serviceRequest13 := models.ServiceRequestModel{
		Id:              serviceReqId,
		UserId:          "auth0|6640b87db9de48e5a6243618", // test1@example.com
		PipelineId:      pipeline2.Id.Hex(),
		PipelineName:    pipeline2.PipelineName,
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

	serviceRequests := [...]models.ServiceRequestModel{
		serviceRequest1,
		serviceRequest2,
		serviceRequest3,
		serviceRequest4,
		serviceRequest5,
		serviceRequest6,
		serviceRequest7,
		serviceRequest8,
		serviceRequest9,
		serviceRequest10,
		serviceRequest11,
		serviceRequest12,
		serviceRequest13,
	}

	for i := 0; i < len(serviceRequests); i++ {
		res, err := database.NewServiceRequest(c).Create(&serviceRequests[i])
		if err != nil {
			panic(err)
		}
		if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
			logger.Info("created service request with ID: " + oid.String())
		} else {
			panic("Inserted ID is not an ObjectID")
		}
	}

	// start SRs
	event.FireAsync(events.NewNewServiceRequestEvent(&serviceRequest1))
	event.FireAsync(events.NewNewServiceRequestEvent(&serviceRequest7))
	event.FireAsync(events.NewNewServiceRequestEvent(&serviceRequest3))
	event.FireAsync(events.NewNewServiceRequestEvent(&serviceRequest6))
	event.FireAsync(events.NewNewServiceRequestEvent(&serviceRequest9))
	event.FireAsync(events.NewNewServiceRequestEvent(&serviceRequest12))

}
