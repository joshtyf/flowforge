package server

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrWrongStepType       = errors.New("wrong step type")
	ErrJsonParseError      = errors.New("unable to parse json request body")

	ErrPipelineCreateFail = errors.New("failed to create pipeline")
	ErrInvalidPipelineId  = errors.New("invalid pipeline id")

	ErrInvalidServiceRequestId        = errors.New("invalid service request id")
	ErrInvalidServiceRequestStatus    = errors.New("invalid service request status")
	ErrServiceRequestNotStarted       = errors.New("service request not started")
	ErrServiceRequestAlreadyStarted   = errors.New("service request already started")
	ErrServiceRequestAlreadyCompleted = errors.New("service request already completed")

	ErrUnableToValidateJWT = errors.New("unable to validate JWT")
	ErrUnauthorised        = errors.New("user does not have required permissions")

	ErrInvalidUserId          = errors.New("invalid user id")
	ErrUserCreateFail         = errors.New("failed to create user")
	ErrUserRetrieve           = errors.New("failed to retrieve user")
	ErrOrganizationRetrieve   = errors.New("failed to retrieve user organizations")
	ErrOrganizationCreateFail = errors.New("failed to create organization")
	ErrOrganizationDeleteFail = errors.New("failed to delete organization")
	ErrInvalidOrganizationId  = errors.New("invalid organization id")

	ErrMembershipCreateFail = errors.New("failed to create membership")
	ErrMembershipUpdateFail = errors.New("failed to update membership")
	ErrMembershipDeleteFail = errors.New("failed to delete membership")
	ErrMembershipInvalid    = errors.New("invalid membership")

	ErrInvalidOffset = errors.New("invalid offset")
)
