package main

import (
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
		PipelineUuid:    "1",
		PipelineVersion: 1,
		Status:          models.Success,
		Remarks:         "This is a test service request.",
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
