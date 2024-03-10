package validation

import (
	"fmt"

	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/helper"
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

// Validates a form field of a newly created pipeline
func ValidateFormField(f models.FormField) error {
	if f.Name == "" {
		return NewMissingRequiredFieldError("name")
	}
	if f.Type == "" {
		return NewMissingRequiredFieldError("type")
	}
	if f.Type == models.DropdownField || f.Type == models.OptionField || f.Type == models.CheckboxField {
		if f.Values == nil || len(f.Values) == 0 {
			return NewMissingRequiredFieldError("values")
		}
	}
	return nil
}

type FormFieldDataValidator func(models.FormField, any) error

func (f FormFieldDataValidator) validate(field models.FormField, data any) error {
	return f(field, data)
}

// Default text field validator
func defaultTextFieldDataValidator(field models.FormField, data any) error {
	if _, ok := data.(string); !ok {
		return NewInvalidFormDataTypeError(field.Name, "string")
	}

	return nil
}

// Default dropdown field validator
func defaultDropdownFieldDataValidator(field models.FormField, data any) error {
	dataStr, ok := data.(string)
	if !ok {
		return NewInvalidFormDataTypeError(field.Name, "string")
	}
	if !helper.StringInSlice(dataStr, field.Values) {
		return NewInvalidSelectedFormDataError(field.Values, dataStr)
	}

	return nil
}

// Default option field validator
func defaultOptionFieldDataValidator(field models.FormField, data any) error {
	dataStr, ok := data.(string)
	if !ok {
		return NewInvalidFormDataTypeError(field.Name, "string")
	}
	if !helper.StringInSlice(dataStr, field.Values) {
		return NewInvalidSelectedFormDataError(field.Values, dataStr)
	}

	return nil
}

// Default checkbox field validator
func defaultCheckboxFieldDataValidator(field models.FormField, data any) error {
	dataStrings, ok := data.([]string)
	if !ok {
		return NewInvalidFormDataTypeError(field.Name, "[]string")
	}
	for _, s := range dataStrings {
		if !helper.StringInSlice(s, field.Values) {
			return NewInvalidSelectedFormDataError(field.Values, s)
		}
	}
	return nil
}

type FormDataValidator struct {
	fieldDataValidators map[models.FormFieldType]FormFieldDataValidator
}

func NewFormDataValidator(customValidators *map[models.FormFieldType]FormFieldDataValidator) *FormDataValidator {
	validator := &FormDataValidator{
		// Initialize with default field validators
		fieldDataValidators: map[models.FormFieldType]FormFieldDataValidator{
			models.TextField:     FormFieldDataValidator(defaultTextFieldDataValidator),
			models.DropdownField: FormFieldDataValidator(defaultDropdownFieldDataValidator),
			models.OptionField:   FormFieldDataValidator(defaultOptionFieldDataValidator),
			models.CheckboxField: FormFieldDataValidator(defaultCheckboxFieldDataValidator),
		},
	}
	if customValidators != nil {
		for fieldType, customFieldValidator := range *customValidators {
			validator.fieldDataValidators[fieldType] = customFieldValidator
		}
	}
	return validator
}

func (v *FormDataValidator) Validate(formData *models.FormData, form *models.Form) error {
	for _, field := range form.Fields {
		fieldData, ok := (*formData)[field.Name]
		if !ok {
			if field.IsRequired {
				return NewMissingRequiredFieldError(field.Name)
			} else {
				continue
			}
		}

		fieldDataValidator, ok := v.fieldDataValidators[field.Type]
		if !ok {
			return fmt.Errorf("no data validator defined for field '%s' of type '%s'", field.Name, field.Type)
		}
		err := fieldDataValidator.validate(field, fieldData)
		if err != nil {
			return err
		}
	}
	return nil
}
