package server

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrWrongStepType       = errors.New("wrong step type")
	ErrJsonParseError      = errors.New("unable to parse json request body")

	ErrPipelineCreateFail = errors.New("failed to create pipeline")
	ErrInvalidPipelineId  = errors.New("invalid pipeline id")

	ErrInvalidServiceRequestId        = errors.New("invalid service request id")
	ErrServiceRequestNotStarted       = errors.New("service request not started")
	ErrServiceRequestAlreadyStarted   = errors.New("service request already started")
	ErrServiceRequestAlreadyCompleted = errors.New("service request already completed")

	ErrUnableToValidateJWT = errors.New("unable to validate JWT")
	ErrUnauthorised        = errors.New("user does not have required permissions")

	ErrInvalidUserId          = errors.New("invalid user id")
	ErrUserCreateFail         = errors.New("failed to create user")
	ErrUserRetrieve           = errors.New("failed to retrieve user")
	ErrOrganisationRetrieve   = errors.New("failed to retrieve user organisations")
	ErrOrganisationCreateFail = errors.New("failed to create organisation")
	ErrInvalidOrganisationId  = errors.New("invalid organisation id")
	ErrMembershipCreateFail   = errors.New("failed to create membership")
)
