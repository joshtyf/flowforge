package database

import (
	"context"
	"os"

	"github.com/joshtyf/flowforge/src/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c *mongo.Client

func GetMongoClient() (*mongo.Client, error) {
	if c == nil {
		uri := os.Getenv("MONGO_URI")
		if uri == "" {
			logger.Error("MONGO_URI environment variable not set", nil)
		}
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			return nil, err
		}
		c = client
	}

	return c, nil
}
