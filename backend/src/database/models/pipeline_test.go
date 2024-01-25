package models

import "testing"

func TestGetPipelineStep(t *testing.T) {
	pipeline := PipelineModel{
		Steps: []PipelineStepModel{{StepName: "step1"}, {StepName: "step2"}},
	}
	t.Run("Get step 1", func(t *testing.T) {
		step := pipeline.GetPipelineStep("step1")
		if step == nil {
			t.Errorf("Expected step1, got nil")
			return
		}
		if step.StepName != "step1" {
			t.Errorf("Expected step1, got %s", step.StepName)
		}
	})
	t.Run("Expect nil result", func(t *testing.T) {
		step := pipeline.GetPipelineStep("step3")
		if step != nil {
			t.Errorf("Expected nil result, got %v", step)
		}
	})
}
