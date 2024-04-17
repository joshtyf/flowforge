package validation

import (
	"fmt"
	"strings"
)

type InvalidStepTypeError struct {
	stepName         string
	receivedStepType string
}

func NewInvalidStepTypeError(stepName, receivedStepType string) *InvalidStepTypeError {
	return &InvalidStepTypeError{
		stepName:         stepName,
		receivedStepType: receivedStepType,
	}
}

func (e *InvalidStepTypeError) Error() string {
	return fmt.Sprintf("invalid step type for step '%s': '%s'", e.stepName, e.receivedStepType)
}

type MissingRequiredFieldError struct {
	fieldName string
}

func NewMissingRequiredFieldError(fieldName string) *MissingRequiredFieldError {
	return &MissingRequiredFieldError{
		fieldName: fieldName,
	}
}

type InvalidPropertyError struct {
	fieldName string
}

func (e *InvalidPropertyError) Error() string {
	return fmt.Sprintf("invalid value found for field: '%s'", e.fieldName)
}

func NewInvalidPropertyValue(fieldName string) *InvalidPropertyError {
	return &InvalidPropertyError{
		fieldName: fieldName,
	}
}

func (e *MissingRequiredFieldError) Error() string {
	return fmt.Sprintf("missing required field: '%s'", e.fieldName)
}

type ZeroStepsError struct{}

func NewZeroStepsError() *ZeroStepsError {
	return &ZeroStepsError{}
}

func (e *ZeroStepsError) Error() string {
	return "pipeline must have at least one step"
}

type NoStepNameFoundError struct {
	fieldName        string
	receivedStepName string
}

func NewNoStepNameFoundError(fieldName, receivedStepName string) *NoStepNameFoundError {
	return &NoStepNameFoundError{
		fieldName:        fieldName,
		receivedStepName: receivedStepName,
	}
}

func (e *NoStepNameFoundError) Error() string {
	return fmt.Sprintf("no step name '%s' found for field '%s'", e.receivedStepName, e.fieldName)
}

type NoNextStepError struct {
	stepName string
}

func NewNoNextStepError(stepName string) *NoNextStepError {
	return &NoNextStepError{
		stepName: stepName,
	}
}

func (e *NoNextStepError) Error() string {
	return fmt.Sprintf("non-terminal step '%s' must define next step", e.stepName)
}

type DuplicateStepNameError struct {
	stepName string
}

func NewDuplicateStepNameError(stepName string) *DuplicateStepNameError {
	return &DuplicateStepNameError{
		stepName: stepName,
	}
}

func (e *DuplicateStepNameError) Error() string {
	return fmt.Sprintf("duplicate step name found: '%s'", e.stepName)
}

type InvalidStepReferenceError struct {
	firstStepName         string
	firstStepNextStepRef  string
	secondStepName        string
	secondStepPrevStepRef string
}

func NewInvalidStepReferenceError(firstStepName, firstStepNextStepRef, secondStepName, secondStepPrevStepRef string) *InvalidStepReferenceError {
	return &InvalidStepReferenceError{
		firstStepName:         firstStepName,
		firstStepNextStepRef:  firstStepNextStepRef,
		secondStepName:        secondStepName,
		secondStepPrevStepRef: secondStepPrevStepRef,
	}
}

func (e *InvalidStepReferenceError) Error() string {
	return fmt.Sprintf("step '%s' references step '%s' as next step, but step '%s' references step '%s' as previous step",
		e.firstStepName, e.firstStepNextStepRef, e.secondStepName, e.secondStepPrevStepRef)
}

type FirstStepContainsPrevStepError struct {
	stepName string
}

func NewFirstStepContainsPrevStepError(stepName string) *FirstStepContainsPrevStepError {
	return &FirstStepContainsPrevStepError{
		stepName: stepName,
	}
}

func (e *FirstStepContainsPrevStepError) Error() string {
	return fmt.Sprintf("first step '%s' cannot contain prev step", e.stepName)
}

type InvalidFirstStepReference struct {
	expectedFirstStep string
}

func NewInvalidFirstStepReference(expectedFirstStep string) *InvalidFirstStepReference {
	return &InvalidFirstStepReference{
		expectedFirstStep: expectedFirstStep,
	}
}

func (e *InvalidFirstStepReference) Error() string {
	return fmt.Sprintf("first step '%s' not found", e.expectedFirstStep)
}

type CircularReferenceError struct {
	startStepName string
	endStepName   string
}

func NewCircularReferenceError(startStepName, endStepName string) *CircularReferenceError {
	return &CircularReferenceError{
		startStepName: startStepName,
		endStepName:   endStepName,
	}
}

func (e *CircularReferenceError) Error() string {
	return fmt.Sprintf("circular reference detected between steps '%s' and '%s'", e.startStepName, e.endStepName)
}

type InvalidFormDataTypeError struct {
	fieldName    string
	expectedType string
}

func NewInvalidFormDataTypeError(fieldName, expectedType string) *InvalidFormDataTypeError {
	return &InvalidFormDataTypeError{
		fieldName:    fieldName,
		expectedType: expectedType,
	}
}

func (e *InvalidFormDataTypeError) Error() string {
	return fmt.Sprintf("data provided for '%s' is not of type '%s'", e.fieldName, e.expectedType)
}

type InvalidSelectedFormDataError struct {
	expectedValues []string
	receivedValue  string
}

func NewInvalidSelectedFormDataError(expectedValues []string, receivedValue string) *InvalidSelectedFormDataError {
	return &InvalidSelectedFormDataError{
		expectedValues: expectedValues,
		receivedValue:  receivedValue,
	}
}

func (e *InvalidSelectedFormDataError) Error() string {
	return fmt.Sprintf("expected selected value to be one of '%s', got '%s' instead", strings.Join(e.expectedValues, ","), e.receivedValue)
}
