package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PipelineStepType string

const (
	APIStep PipelineStepType = "API"
)

type PipelineStepModel struct {
	StepName       string            `bson:"step_name" json:"step_name"`
	StepType       PipelineStepType  `bson:"step_type" json:"step_type"`
	NextStepName   string            `bson:"next_step_name" json:"next_step_name"`
	PrevStepName   string            `bson:"prev_step_name" json:"prev_step_name"`
	Parameters     map[string]string `bson:"parameters" json:"parameters"`
	IsTerminalStep bool              `bson:"is_terminal_step" json:"is_terminal_step"`
}

type PipelineModel struct {
	Uuid          primitive.ObjectID  `bson:"_id,omitempty" json:"uuid,omitempty"` // unique id for the pipeline
	PipelineName  string              `bson:"pipeline_name" json:"pipeline_name"`
	PipelineId    string              `bson:"pipeline_id" json:"pipeline_id"` // id for the pipeline. non-unique and uses `Version` to differentiate between different versions of the pipeline
	Version       int                 `bson:"version" json:"version"`
	FirstStepName string              `bson:"first_step_name" json:"first_step_name"`
	Steps         []PipelineStepModel `bson:"steps" json:"steps"`
	CreatedOn     time.Time           `bson:"created_on" json:"created_on"`
}

func (p *PipelineModel) GetPipelineStep(name string) *PipelineStepModel {
	for _, step := range p.Steps {
		if step.StepName == name {
			return &step
		}
	}
	return nil
}
