package server

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrWrongStepType       = errors.New("wrong step type")
	ErrJsonParseError      = errors.New("unable to parse json request body")

	ErrPipelineCreateFail = errors.New("failed to create pipeline")

	ErrServiceRequestNotStarted       = errors.New("service request not started")
	ErrServiceRequestAlreadyStarted   = errors.New("service request already started")
	ErrServiceRequestAlreadyCompleted = errors.New("service request already completed")

	ErrUserCreateFail = errors.New("failed to create user")
)
