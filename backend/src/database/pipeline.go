package database

import (
	"context"
	"time"

	"github.com/joshtyf/flowforge/src/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Pipeline struct {
	c *mongo.Client
}

func NewPipeline(c *mongo.Client) *Pipeline {
	return &Pipeline{c: c}
}

func (p *Pipeline) Create(pm *models.PipelineModel) (*mongo.InsertOneResult, error) {
	pm.CreatedOn = time.Now()
	res, err := p.c.Database(DatabaseName).Collection("pipelines").InsertOne(context.Background(), pm)
	return res, err
}

func (p *Pipeline) GetById(id string) (*models.PipelineModel, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res := p.c.Database(DatabaseName).Collection("pipelines").FindOne(context.Background(), bson.M{"_id": objId})
	if res.Err() != nil {
		return nil, res.Err()
	}
	pipeline := &models.PipelineModel{}
	res.Decode(pipeline)
	return pipeline, nil
}
