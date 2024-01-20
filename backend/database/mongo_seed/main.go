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

	serviceRequest := models.ServiceRequestModel{
		PipelineId:      "1",
		PipelineVersion: 1,
		Status:          "Pending",
		Remarks:         "This is a test service request.",
		CreatedOn:       time.Now(),
		LastUpdated:     time.Now(),
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
}
