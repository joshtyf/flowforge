package validation

import (
	"github.com/joshtyf/flowforge/src/database/models"
)

func ValidatePipeline(pipeline *models.PipelineModel) error {
	if pipeline.PipelineName == "" {
		return NewMissingRequiredFieldError("pipeline_name")
	}
	if pipeline.Steps == nil {
		return NewMissingRequiredFieldError("steps")
	}
	if len(pipeline.Steps) == 0 {
		return NewZeroStepsError()
	}
	if pipeline.FirstStepName == "" {
		return NewMissingRequiredFieldError("first_step_name")
	}
	stepNames := make(map[string]bool)
	for _, step := range pipeline.Steps {
		if !models.IsValidPipelineStepType(step.StepType) {
			return NewInvalidStepTypeError(step.StepName, string(step.StepType))
		}
		if step.StepName == "" {
			return NewMissingRequiredFieldError("step_name")
		} else if stepNames[step.StepName] {
			return NewDuplicateStepNameError(step.StepName)
		} else {
			stepNames[step.StepName] = true
		}
		if step.NextStepName == "" && !step.IsTerminalStep {
			return NewNoNextStepError(step.StepName)
		}
		if stepNames[step.NextStepName] {
			return NewCircularReferenceError(step.NextStepName, step.StepName)
		}
		if step.StepName == pipeline.FirstStepName && step.PrevStepName != "" {
			return NewFirstStepContainsPrevStepError(step.StepName)
		}
	}

	if !stepNames[pipeline.FirstStepName] {
		return NewInvalidFirstStepReference(pipeline.FirstStepName)
	}

	for _, step := range pipeline.Steps {
		if step.PrevStepName != "" && !stepNames[step.PrevStepName] {
			return NewNoStepNameFoundError("prev_step_name", step.PrevStepName)
		}
		if step.PrevStepName != "" && pipeline.GetPipelineStep(step.PrevStepName).NextStepName != step.StepName {
			prevStep := pipeline.GetPipelineStep(step.PrevStepName)
			return NewInvalidStepReferenceError(prevStep.StepName, prevStep.NextStepName, step.StepName, step.PrevStepName)
		}
		if step.NextStepName != "" && !stepNames[step.NextStepName] {
			return NewNoStepNameFoundError("next_step_name", step.NextStepName)
		}
	}

	return nil
}
