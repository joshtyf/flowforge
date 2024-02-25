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

	ErrBearerTokenNotFound       = errors.New("bearer token not found")
	ErrUnableToVerifyBearerToken = errors.New("unable to verify token")
	ErrUnableToRetrieveProfile   = errors.New("unable to retrieve profile")
)
