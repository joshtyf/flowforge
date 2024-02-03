package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PipelineStepType string

const (
	APIStep             PipelineStepType = "API"
	WaitForApprovalStep PipelineStepType = "WAIT_FOR_APPROVAL"
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
	Id            primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"` // unique id for the pipeline
	PipelineName  string              `bson:"pipeline_name" json:"pipeline_name"`
	Version       int                 `bson:"version" json:"version"`
	PrevVersionId primitive.ObjectID  `bson:"prev_version_id" json:"prev_version_id"`
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
