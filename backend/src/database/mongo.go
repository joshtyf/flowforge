package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClientWrapper struct {
	client *mongo.Client
}

var c *mongo.Client

func init() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI environment variable not set") // TODO: replace with custom logger
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("failed to create mongo client, err=%v", err) // TODO: replace with custom logger
	}
	c = client
}

func NewMongoClientWrapper() *MongoClientWrapper {
	return &MongoClientWrapper{client: c}
}

func (mcw *MongoClientWrapper) GetClient() *mongo.Client {
	return mcw.client
}
