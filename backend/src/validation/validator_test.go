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

func TestValidateFormField(t *testing.T) {
	testcases := []struct {
		testDescription string
		formField       models.FormField
		expected        error
	}{
		{
			"Valid text field",
			models.FormField{
				Name: "text", Type: models.TextField,
			},
			nil,
		},
		{
			"Valid dropdown field",
			models.FormField{
				Name: "dropdown", Type: models.DropdownField, Values: []string{"test1", "test2"},
			},
			nil,
		},
		{
			"Valid checkbox field",
			models.FormField{
				Name: "checkbox", Type: models.CheckboxField, Values: []string{"test1", "test2"},
			},
			nil,
		},
		{
			"Valid option field",
			models.FormField{
				Name: "option", Type: models.OptionField, Values: []string{"test1", "test2"},
			},
			nil,
		},
		{
			"Form field missing name",
			models.FormField{
				Type: models.TextField,
			},
			NewMissingRequiredFieldError("name"),
		},
		{
			"Form field missing type",
			models.FormField{
				Name: "test",
			},
			NewMissingRequiredFieldError("type"),
		},
		{
			"Dropdown field nil values",
			models.FormField{
				Name: "dropdown", Type: models.DropdownField,
			},
			NewMissingRequiredFieldError("values"),
		},
		{
			"Checkbox field nil values",
			models.FormField{
				Name: "checkbox", Type: models.CheckboxField,
			},
			NewMissingRequiredFieldError("values"),
		},
		{
			"Option field nil values",
			models.FormField{
				Name: "option", Type: models.OptionField,
			},
			NewMissingRequiredFieldError("values"),
		},
		{
			"Dropdown field 0 values",
			models.FormField{
				Name: "dropdown", Type: models.DropdownField, Values: []string{},
			},
			NewMissingRequiredFieldError("values"),
		},
		{
			"Checkbox field 0 values",
			models.FormField{
				Name: "checkbox", Type: models.CheckboxField, Values: []string{},
			},
			NewMissingRequiredFieldError("values"),
		},
		{
			"Option field 0 values",
			models.FormField{
				Name: "Option", Type: models.OptionField, Values: []string{},
			},
			NewMissingRequiredFieldError("values"),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.testDescription, func(t *testing.T) {
			err := ValidateFormField(tc.formField)
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
func TestTextFieldDataValidator(t *testing.T) {
	formField := models.FormField{
		Name: "test",
		Type: models.TextField,
	}
	testCases := []struct {
		testDescription string
		field           models.FormField
		value           any
		expected        error
	}{
		{
			"Valid text field",
			formField,
			"test",
			nil,
		},
		{
			"Empty text field",
			formField,
			"",
			nil,
		},
		{
			"Data of type int for text field",
			formField,
			1,
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type float for text field",
			formField,
			1.0,
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type bool for text field",
			formField,
			true,
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type []string for text field",
			formField,
			[]string{"test"},
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type map[string]string for text field",
			formField,
			map[string]string{"test": "test"},
			NewInvalidFormDataTypeError("test", "string"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testDescription, func(t *testing.T) {
			validator := FormFieldDataValidator(defaultTextFieldDataValidator)
			err := validator.validate(tc.field, tc.value)
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

func TestDropdownFieldDataValidator(t *testing.T) {
	formField := models.FormField{
		Name:   "test",
		Type:   models.DropdownField,
		Values: []string{"test1", "test2", "test3"},
	}
	testcases := []struct {
		testDescription string
		field           models.FormField
		value           any
		expected        error
	}{
		{
			"Valid dropdown value (1)",
			formField,
			"test1",
			nil,
		},
		{
			"Valid dropdown value (2)",
			formField,
			"test2",
			nil,
		},
		{
			"Valid dropdown value (3)",
			formField,
			"test3",
			nil,
		},
		{
			"Invalid dropdown value",
			formField,
			"test4",
			NewInvalidSelectedFormDataError(formField.Values, "test4"),
		},
		{
			"Data of type int for dropdown field",
			formField,
			1,
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type float for dropdown field",
			formField,
			1.0,
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type bool for dropdown field",
			formField,
			true,
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type []string for dropdown field",
			formField,
			[]string{"test"},
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type map[string]string for dropdown field",
			formField,
			map[string]string{"test": "test"},
			NewInvalidFormDataTypeError("test", "string"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testDescription, func(t *testing.T) {
			validator := FormFieldDataValidator(defaultDropdownFieldDataValidator)
			err := validator.validate(tc.field, tc.value)
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

func TestOptionFieldDataValidator(t *testing.T) {
	formField := models.FormField{
		Name:   "test",
		Type:   models.OptionField,
		Values: []string{"test1", "test2", "test3"},
	}
	testcases := []struct {
		testDescription string
		field           models.FormField
		value           any
		expected        error
	}{
		{
			"Valid option value (1)",
			formField,
			"test1",
			nil,
		},
		{
			"Valid option value (2)",
			formField,
			"test2",
			nil,
		},
		{
			"Valid option value (3)",
			formField,
			"test3",
			nil,
		},
		{
			"Invalid option value",
			formField,
			"test4",
			NewInvalidSelectedFormDataError(formField.Values, "test4"),
		},
		{
			"Data of type int for option field",
			formField,
			1,
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type float for option field",
			formField,
			1.0,
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type bool for option field",
			formField,
			true,
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type []string for option field",
			formField,
			[]string{"test"},
			NewInvalidFormDataTypeError("test", "string"),
		},
		{
			"Data of type map[string]string for option field",
			formField,
			map[string]string{"test": "test"},
			NewInvalidFormDataTypeError("test", "string"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testDescription, func(t *testing.T) {
			validator := FormFieldDataValidator(defaultOptionFieldDataValidator)
			err := validator.validate(tc.field, tc.value)
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

func TestCheckboxFieldDataValidator(t *testing.T) {
	formField := models.FormField{
		Name:   "test",
		Type:   models.CheckboxField,
		Values: []string{"test1", "test2", "test3"},
	}
	testcases := []struct {
		testDescription string
		field           models.FormField
		value           any
		expected        error
	}{
		{
			"Valid checkbox value",
			formField,
			[]string{"test1"},
			nil,
		},
		{
			"Valid multiple checkbox values",
			formField,
			[]string{"test1", "test2"},
			nil,
		},
		{
			"Invalid checkbox value",
			formField,
			[]string{"test4"},
			NewInvalidSelectedFormDataError(formField.Values, "test4"),
		},
		{
			"Invalid multiple checkbox values",
			formField,
			[]string{"test1", "test4"},
			NewInvalidSelectedFormDataError(formField.Values, "test4"),
		},
		{
			"Data of type string for checkbox field",
			formField,
			"test1",
			NewInvalidFormDataTypeError("test", "[]string"),
		},
		{
			"Data of type int for checkbox field",
			formField,
			1,
			NewInvalidFormDataTypeError("test", "[]string"),
		},
		{
			"Data of type float for checkbox field",
			formField,
			1.0,
			NewInvalidFormDataTypeError("test", "[]string"),
		},
		{
			"Data of type bool for checkbox field",
			formField,
			true,
			NewInvalidFormDataTypeError("test", "[]string"),
		},
		{
			"Data of type map[string]string for checkbox field",
			formField,
			map[string]string{"test": "test"},
			NewInvalidFormDataTypeError("test", "[]string"),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.testDescription, func(t *testing.T) {
			validator := FormFieldDataValidator(defaultCheckboxFieldDataValidator)
			err := validator.validate(tc.field, tc.value)
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

func TestValidateFormData_RequiredFieldsValidation(t *testing.T) {
	defaultFormDataValidator := NewFormDataValidator(nil)
	testCases := []struct {
		testDescription   string
		form              models.Form
		formData          models.FormData
		formDataValidator *FormDataValidator
		expected          error
	}{
		{
			"Form with zero required fields. Form data has all fields",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: false},
					{Name: "test2", Type: models.TextField, IsRequired: false},
				},
			},
			models.FormData{
				"test":  "test",
				"test2": "test2",
			},
			defaultFormDataValidator,
			nil,
		},
		{
			"Form with zero required fields. Form data has no fields",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: false},
					{Name: "test2", Type: models.TextField, IsRequired: false},
				},
			},
			models.FormData{},
			defaultFormDataValidator,
			nil,
		},
		{
			"Form with zero required fields. Form data has some fields",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: false},
					{Name: "test2", Type: models.TextField, IsRequired: false},
				},
			},
			models.FormData{
				"test": "test",
			},
			defaultFormDataValidator,
			nil,
		},
		{
			"Form with one required field. Form data has all fields",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: true},
					{Name: "test2", Type: models.TextField, IsRequired: false},
				},
			},
			models.FormData{
				"test":  "test",
				"test2": "test2",
			},
			defaultFormDataValidator,
			nil,
		},
		{
			"Form with one required field. Form data has no fields",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: true},
					{Name: "test2", Type: models.TextField, IsRequired: false},
				},
			},
			models.FormData{},
			defaultFormDataValidator,
			NewMissingRequiredFieldError("test"),
		},
		{
			"Form with one required field. Form data has required field",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: true},
					{Name: "test2", Type: models.TextField, IsRequired: false},
				},
			},
			models.FormData{
				"test": "test",
			},
			defaultFormDataValidator,
			nil,
		},
		{
			"Form with one required field. Form data does not have field",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: true},
					{Name: "test2", Type: models.TextField, IsRequired: false},
				},
			},
			models.FormData{
				"test2": "test2",
			},
			defaultFormDataValidator,
			NewMissingRequiredFieldError("test"),
		},
		{
			"Form with multiple required fields. Form data has all fields",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: true},
					{Name: "test2", Type: models.TextField, IsRequired: true},
				},
			},
			models.FormData{
				"test":  "test",
				"test2": "test2",
			},
			defaultFormDataValidator,
			nil,
		},
		{
			"Form with multiple required fields. Form data has no fields",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: true},
					{Name: "test2", Type: models.TextField, IsRequired: true},
				},
			},
			models.FormData{},
			defaultFormDataValidator,
			NewMissingRequiredFieldError("test"),
		},
		{
			"Form with multiple required fields. Form data has some fields",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: true},
					{Name: "test2", Type: models.TextField, IsRequired: true},
				},
			},
			models.FormData{
				"test": "test",
			},
			defaultFormDataValidator,
			NewMissingRequiredFieldError("test2"),
		},
		{
			"Form only has 3 fields. Form data has more",
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField, IsRequired: true},
					{Name: "test2", Type: models.TextField, IsRequired: true},
					{Name: "test3", Type: models.TextField, IsRequired: true},
				},
			},
			models.FormData{
				"test":  "test",
				"test2": "test2",
				"test3": "test3",
				"test4": "test4",
			},
			defaultFormDataValidator,
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testDescription, func(t *testing.T) {
			err := tc.formDataValidator.Validate(&tc.formData, &tc.form)
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

func TestValidateFormData_ValidatorsCalledOnce(t *testing.T) {
	type counter struct {
		count int
	}
	// This dummy validator does not perform any validation, only increments the counter if it was called on the correct form field type
	dummyValidator := func(c *counter, f models.FormFieldType) FormFieldDataValidator {
		return FormFieldDataValidator(func(field models.FormField, data any) error {
			if field.Type == f {
				c.count += 1
			}
			return nil
		})
	}
	testcases := []struct {
		testDescription            string
		expectedValidatorFieldType models.FormFieldType
		form                       models.Form
		formData                   models.FormData
	}{
		{
			"Test text field validator called",
			models.TextField,
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.TextField},
				},
			},
			models.FormData{
				"test": "test",
			},
		},
		{
			"Test checkbox field validator called",
			models.CheckboxField,
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.CheckboxField, Values: []string{"test1"}},
				},
			},
			models.FormData{
				"test": "test1",
			},
		},
		{
			"Test dropdown field validator called",
			models.DropdownField,
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.DropdownField, Values: []string{"test1"}},
				},
			},
			models.FormData{
				"test": "test1",
			},
		},
		{
			"Test option field validator called",
			models.OptionField,
			models.Form{
				Fields: []models.FormField{
					{Name: "test", Type: models.OptionField, Values: []string{"test1"}},
				},
			},
			models.FormData{
				"test": "test1",
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.testDescription, func(t *testing.T) {
			counter := &counter{count: 0}
			formValidator := NewFormDataValidator(&map[models.FormFieldType]FormFieldDataValidator{
				tc.expectedValidatorFieldType: dummyValidator(counter, tc.expectedValidatorFieldType),
			})
			formValidator.Validate(&tc.formData, &tc.form)
			if counter.count != 1 {
				t.Errorf("Expected validator of %s to have been called once, actual calls: %d", tc.expectedValidatorFieldType, counter.count)
			}
		})
	}
}
