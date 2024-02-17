package validation

import (
	"testing"

	"github.com/joshtyf/flowforge/src/database/models"
)

func TestValidatePipeline(t *testing.T) {
	testCases := []struct {
		testDescription string
		pipeline        *models.PipelineModel
		expected        error
	}{
		{
			"Valid pipeline with 1 step",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, IsTerminalStep: true},
				},
				FirstStepName: "step1",
			},
			nil,
		},
		{
			"Valid pipeline with 2 steps",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, NextStepName: "step2"},
					{StepName: "step2", StepType: models.APIStep, IsTerminalStep: true},
				},
				FirstStepName: "step1",
			},
			nil,
		},
		{
			"Valid pipeline with 3 steps",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, NextStepName: "step2"},
					{StepName: "step2", StepType: models.APIStep, NextStepName: "step3"},
					{StepName: "step3", StepType: models.APIStep, IsTerminalStep: true},
				},
				FirstStepName: "step1",
			},
			nil,
		},
		{
			"Pipeline name is empty",
			&models.PipelineModel{
				PipelineName: "",
			},
			NewMissingRequiredFieldError("pipeline_name"),
		},
		{
			"Steps is nil",
			&models.PipelineModel{
				PipelineName: "test",
				Steps:        nil,
			},
			NewMissingRequiredFieldError("steps"),
		},
		{
			"Steps is empty",
			&models.PipelineModel{
				PipelineName: "test",
				Steps:        []models.PipelineStepModel{},
			},
			NewZeroStepsError(),
		},
		{
			"Invalid step type",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: "invalid", IsTerminalStep: true},
				},
				FirstStepName: "step1",
			},
			NewInvalidStepTypeError("step1", "invalid"),
		},
		{
			"First step name is undefined",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: "step"},
				},
			},
			NewMissingRequiredFieldError("first_step_name"),
		},
		{
			"Step name is empty",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "", StepType: models.APIStep},
				},
				FirstStepName: "step1",
			},
			NewMissingRequiredFieldError("step_name"),
		},
		{
			"No next step name for non-terminal step",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep},
				},
				FirstStepName: "step1",
			},
			NewNoNextStepError("step1"),
		},
		{
			"First step contains prev step",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, PrevStepName: "step2", NextStepName: "step2"},
				},
				FirstStepName: "step1",
			},
			NewFirstStepContainsPrevStepError("step1"),
		},
		{
			"Invalid next step name for non-terminal step",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, NextStepName: "step3"},
					{StepName: "step2", StepType: models.APIStep, IsTerminalStep: true},
				},
				FirstStepName: "step1",
			},
			NewNoStepNameFoundError("next_step_name", "step3"),
		},
		{
			"Invalid prev step name",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, NextStepName: "step2"},
					{StepName: "step2", StepType: models.APIStep, PrevStepName: "step3", IsTerminalStep: true}},
				FirstStepName: "step1",
			},
			NewNoStepNameFoundError("prev_step_name", "step3"),
		},
		{
			"Duplicate step name",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, IsTerminalStep: true},
					{StepName: "step1", StepType: models.APIStep, IsTerminalStep: true},
				},
				FirstStepName: "step1",
			},
			NewDuplicateStepNameError("step1"),
		},
		{
			"Invalid reference between steps",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, NextStepName: "step3"},
					{StepName: "step2", StepType: models.APIStep, PrevStepName: "step1", NextStepName: "step3"},
					{StepName: "step3", StepType: models.APIStep, IsTerminalStep: true},
				},
				FirstStepName: "step1",
			},
			NewInvalidStepReferenceError("step1", "step3", "step2", "step1"),
		},
		{
			"Invalid first step reference",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, IsTerminalStep: true},
				},
				FirstStepName: "step2",
			},
			NewInvalidFirstStepReference("step2"),
		},
		{
			"Circular reference with 2 steps",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, NextStepName: "step2"},
					{StepName: "step2", StepType: models.APIStep, NextStepName: "step1"},
				},
				FirstStepName: "step1",
			},
			NewCircularReferenceError("step1", "step2"),
		},
		{
			"Circular reference with 3 steps",
			&models.PipelineModel{
				PipelineName: "test",
				Steps: []models.PipelineStepModel{
					{StepName: "step1", StepType: models.APIStep, NextStepName: "step2"},
					{StepName: "step2", StepType: models.APIStep, NextStepName: "step3"},
					{StepName: "step3", StepType: models.APIStep, NextStepName: "step1"},
				},
				FirstStepName: "step1",
			},
			NewCircularReferenceError("step1", "step3"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testDescription, func(t *testing.T) {
			err := ValidatePipeline(tc.pipeline)
			if err == nil {
				if tc.expected != nil {
					t.Errorf("Expected error %v, got nil", tc.expected)
				}
				return
			}
			if tc.expected == nil {
				t.Errorf("Expected no error, got %v", err)
				return
			}
			if err.Error() != tc.expected.Error() {
				t.Errorf("Expected %v, got %v", tc.expected, err)
			}
		})
	}
}
