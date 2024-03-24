package client

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c *mongo.Client

func GetMongoClient() (*mongo.Client, error) {
	if c == nil {
		uri := os.Getenv("MONGO_URI")
		if uri == "" {
			return nil, fmt.Errorf("MONGO_URI environment variable not set")
		}
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
		if err != nil {
			return nil, err
		}
		c = client
	}

	return c, nil
}
