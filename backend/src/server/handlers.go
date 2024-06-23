package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gookit/event"
	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/events"
	"github.com/joshtyf/flowforge/src/helper"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/util"
	"github.com/joshtyf/flowforge/src/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandlerError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func newHandlerError(err error, code int) *HandlerError {
	return &HandlerError{Message: err.Error(), Code: code}
}

type ServerHandler struct {
	logger      logger.ServerLogger
	psqlClient  *sql.DB
	mongoClient *mongo.Client
}

func NewServerHandler(psqlClient *sql.DB, mongoCLient *mongo.Client, logger logger.ServerLogger) *ServerHandler {
	return &ServerHandler{
		psqlClient:  psqlClient,
		mongoClient: mongoCLient,
		logger:      logger,
	}
}

func (s *ServerHandler) registerRoutes(r *mux.Router) {
	// Health Check
	r.Handle("/api/healthcheck", handleHealthCheck(s.logger)).Methods("GET")

	// Service Request
	r.Handle("/api/service_request", isAuthenticated(getOrgIdFromQuery(isOrgMember(s.psqlClient, handleGetServiceRequestsByUserAndOrganization(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("GET")
	r.Handle("/api/service_request/admin", isAuthenticated(getOrgIdFromQuery(isOrgAdmin(s.psqlClient, handleGetServiceRequestsForAdminByOrganization(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("GET")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgMember(s.psqlClient, handleGetServiceRequest(s.logger, s.mongoClient, s.psqlClient), s.logger), s.logger), s.logger)).Methods("GET")
	r.Handle("/api/service_request", isAuthenticated(getOrgIdFromRequestBody(isOrgMember(s.psqlClient, handleCreateServiceRequest(s.logger, s.mongoClient, s.psqlClient), s.logger), s.logger), s.logger)).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(getOrgIdFromRequestBody(isOrgMember(s.psqlClient, handleUpdateServiceRequest(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}/cancel", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgMember(s.psqlClient, handleCancelServiceRequest(s.logger, s.mongoClient, s.psqlClient), s.logger), s.logger), s.logger)).Methods("PUT")
	r.Handle("/api/service_request/{requestId}/start", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgMember(s.psqlClient, handleStartServiceRequest(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("PUT")
	r.Handle("/api/service_request/{requestId}/approve", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgAdmin(s.psqlClient, handleApproveServiceRequest(s.logger, s.mongoClient, s.psqlClient), s.logger), s.logger), s.logger)).Methods("PUT")
	r.Handle("/api/service_request/{requestId}/reject", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgAdmin(s.psqlClient, handleRejectServiceRequest(s.logger, s.mongoClient, s.psqlClient), s.logger), s.logger), s.logger)).Methods("PUT")
	r.Handle("/api/service_request/{requestId}/logs/{stepName}", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgMember(s.psqlClient, handleGetStepExecutionLogs(s.logger, s.psqlClient), s.logger), s.logger), s.logger)).Methods("GET")
	r.Handle("/api/service_request/{requestId}/steps", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgMember(s.psqlClient, handleGetServiceRequestStepDetails(s.logger, s.mongoClient, s.psqlClient), s.logger), s.logger), s.logger)).Methods("GET")

	// Pipeline
	r.Handle("/api/pipeline", isAuthenticated(getOrgIdFromQuery(isOrgMember(s.psqlClient, handleGetAllPipelines(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("GET")
	r.Handle("/api/pipeline/{pipelineId}", isAuthenticated(getOrgIdFromQuery(isOrgMember(s.psqlClient, handleGetPipeline(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("GET")
	r.Handle("/api/pipeline", isAuthenticated(getOrgIdFromRequestBody(isOrgAdmin(s.psqlClient, validateCreatePipelineRequest(handleCreatePipeline(s.logger, s.mongoClient), s.logger), s.logger), s.logger), s.logger)).Methods("POST").Headers("Content-Type", "application/json")

	// User
	r.Handle("/api/user", isAuthenticated(handleGetAllUsers(s.logger, s.psqlClient), s.logger)).Methods("GET")
	r.Handle("/api/user/{userId}", isAuthenticated(handleGetUserById(s.logger, s.psqlClient), s.logger)).Methods("GET")
	r.Handle("/api/login", isAuthenticated(handleUserLogin(s.logger, s.psqlClient), s.logger)).Methods("GET")
	r.Handle("/api/user", isAuthenticated(handleCreateUser(s.logger, s.psqlClient), s.logger)).Methods("POST").Headers("Content-Type", "application/json")

	// Organization
	r.Handle("/api/organization", isAuthenticated(handleCreateOrganization(s.logger, s.psqlClient), s.logger)).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/organization", isAuthenticated(handleGetOrganizationsForUser(s.logger, s.psqlClient), s.logger)).Methods("GET")
	r.Handle("/api/organization", isAuthenticated(getOrgIdFromRequestBody(isOrgOwner(s.psqlClient, handleUpdateOrganization(s.logger, s.psqlClient), s.logger), s.logger), s.logger)).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/organization", isAuthenticated(getOrgIdFromRequestBody(isOrgOwner(s.psqlClient, handleDeleteOrganization(s.logger, s.psqlClient), s.logger), s.logger), s.logger)).Methods("DELETE").Headers("Content-Type", "application/json")
	r.Handle("/api/organization/{orgId}/members", isAuthenticated(handleGetOrganizationMembers(s.logger, s.psqlClient), s.logger)).Methods("GET")
	r.Handle("/api/organization/{orgId}/membership", isAuthenticated(handleLeaveOrganization(s.logger, s.psqlClient), s.logger)).Methods("DELETE").Headers("Content-Type", "application/json")

	// Membership
	r.Handle("/api/membership", isAuthenticated(handleGetMembershipsForUser(s.logger, s.psqlClient), s.logger)).Methods("GET")
	r.Handle("/api/membership", isAuthenticated(getOrgIdFromRequestBody(validateMembershipChange(s.psqlClient, isOrgAdmin(s.psqlClient, handleCreateMembership(s.logger, s.psqlClient), s.logger), s.logger), s.logger), s.logger)).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/membership", isAuthenticated(getOrgIdFromRequestBody(validateMembershipChange(s.psqlClient, isOrgAdmin(s.psqlClient, handleUpdateMembership(s.logger, s.psqlClient), s.logger), s.logger), s.logger), s.logger)).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/membership", isAuthenticated(getOrgIdFromRequestBody(validateMembershipChange(s.psqlClient, isOrgAdmin(s.psqlClient, handleDeleteMembership(s.logger, s.psqlClient), s.logger), s.logger), s.logger), s.logger)).Methods("DELETE").Headers("Content-Type", "application/json")
	r.Handle("/api/membership/ownership_transfer", isAuthenticated(getOrgIdFromRequestBody(isOrgOwner(s.psqlClient, handleOwnershipTransfer(s.logger, s.psqlClient), s.logger), s.logger), s.logger)).Methods("POST").Headers("Content-Type", "application/json")
}

func handleHealthCheck(l logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if serverHealthy() {
			l.Info("server is healthy")
			encode(w, r, http.StatusOK, "Server working!")
			return
		}
		encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
	})
}

func handleGetServiceRequest(logger logger.ServerLogger, mongoClient *mongo.Client, psqlClient *sql.DB) http.Handler {
	type ResponseBodyStep struct {
		Name         string           `json:"name"`
		Status       models.EventType `json:"status"`
		UpdatedAt    time.Time        `json:"updated_at"`
		UpdatedBy    string           `json:"updated_by"`
		NextStepName string           `json:"next_step_name"`
	}
	type ResponseBodyPipeline struct {
		Name string       `json:"name"`
		Form *models.Form `json:"form"`
	}
	type ResponseBody struct {
		ServiceRequest *models.ServiceRequestModel `json:"service_request"`
		Steps          map[string]ResponseBodyStep `json:"steps"`
		FirstStepName  string                      `json:"first_step_name"`
		Pipeline       *ResponseBodyPipeline       `json:"pipeline"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		sr, err := database.NewServiceRequest(mongoClient).GetById(requestId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		pipeline, err := database.NewPipeline(mongoClient).GetById(sr.PipelineId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		sre, err := database.NewServiceRequestEvent(psqlClient).GetStepsLatestEvent(requestId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		steps := make(map[string]ResponseBodyStep, len(sre))
		for _, event := range sre {
			step := pipeline.GetPipelineStep(event.StepName)
			if step == nil {
				logger.Error(fmt.Sprintf("%s exists in event log but not in pipeline template", event.StepName))
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
				return
			}
			steps[event.StepName] = ResponseBodyStep{
				Name:         event.StepName,
				Status:       event.EventType,
				UpdatedAt:    event.CreatedAt,
				UpdatedBy:    event.CreatedBy,
				NextStepName: step.NextStepName,
			}
		}
		response := ResponseBody{
			ServiceRequest: sr,
			Steps:          steps,
			FirstStepName:  pipeline.FirstStepName,
			Pipeline: &ResponseBodyPipeline{
				Name: pipeline.PipelineName,
				Form: &pipeline.Form,
			},
		}
		encode(w, r, http.StatusOK, response)
	})
}

func handleGetServiceRequestStepDetails(logger logger.ServerLogger, mongoClient *mongo.Client, psqlClient *sql.DB) http.Handler {
	type ResponseBodyStep struct {
		Name         string           `json:"name"`
		Status       models.EventType `json:"status"`
		UpdatedAt    time.Time        `json:"updated_at"`
		UpdatedBy    string           `json:"updated_by"`
		NextStepName string           `json:"next_step_name"`
	}
	type ResponseBody struct {
		Steps            map[string]ResponseBodyStep `json:"steps"`
		ServiceRequestId string                      `json:"service_request_id"`
		PipelineId       string                      `json:"pipeline_id"`
		PipelineVersion  int                         `json:"pipeline_version"`
		FirstStepName    string                      `json:"first_step_name"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		sre, err := database.NewServiceRequestEvent(psqlClient).GetStepsLatestEvent(requestId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		sr, err := database.NewServiceRequest(mongoClient).GetById(requestId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		pipeline, err := database.NewPipeline(mongoClient).GetById(sr.PipelineId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		steps := make(map[string]ResponseBodyStep, len(sre))
		for _, event := range sre {
			step := pipeline.GetPipelineStep(event.StepName)
			if step == nil {
				logger.Error(fmt.Sprintf("%s exists in event log but not in pipeline template", event.StepName))
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
				return
			}
			steps[event.StepName] = ResponseBodyStep{
				Name:         event.StepName,
				Status:       event.EventType,
				UpdatedAt:    event.CreatedAt,
				UpdatedBy:    event.CreatedBy,
				NextStepName: step.NextStepName,
			}
		}
		response := ResponseBody{
			Steps:            steps,
			ServiceRequestId: requestId,
			PipelineId:       pipeline.Id.Hex(),
			PipelineVersion:  pipeline.Version,
			FirstStepName:    pipeline.FirstStepName,
		}

		encode(w, r, http.StatusOK, response)
	})
}

func handleCreateServiceRequest(logger logger.ServerLogger, mongoClient *mongo.Client, psqlClient *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srm, err := decode[models.ServiceRequestModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}

		pipeline, err := database.NewPipeline(mongoClient).GetById(srm.PipelineId)
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("%s %s not found", "pipeline", srm.PipelineId))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidPipelineId, http.StatusBadRequest))
			return
		} else if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		srm.CreatedOn = time.Now()
		srm.LastUpdated = time.Now()
		srm.Status = models.NOT_STARTED
		srm.PipelineName = pipeline.PipelineName
		srm.PipelineVersion = pipeline.Version

		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		userId := token.RegisteredClaims.Subject
		srm.UserId = userId

		res, err := database.NewServiceRequest(mongoClient).Create(&srm)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		stepEventDAO := database.NewServiceRequestEvent(psqlClient)
		for _, step := range pipeline.Steps {
			err = stepEventDAO.Create(&models.ServiceRequestEventModel{
				EventType:        models.STEP_NOT_STARTED,
				ServiceRequestId: res.InsertedID.(primitive.ObjectID).Hex(),
				StepName:         step.StepName,
				StepType:         step.StepType,
			})
			if err != nil {
				logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
				return
			}
		}
		insertedId, _ := res.InsertedID.(primitive.ObjectID)
		srm.Id = insertedId
		encode(w, r, http.StatusCreated, srm)
	})
}

func handleCancelServiceRequest(logger logger.ServerLogger, client *mongo.Client, psqlClient *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		sr, err := database.NewServiceRequest(client).GetById(requestId)
		if err != nil {
			logger.Error(fmt.Sprintf("%s %s not found", "service request", requestId))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		status := sr.Status
		if status == models.COMPLETED || status == models.FAILED || status == models.CANCELLED {
			logger.Error(fmt.Sprintf("failed to %s service request %s: sr status %s not eligible for cancellation", "cancel", requestId, status))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrServiceRequestAlreadyCompleted, http.StatusBadRequest))
			return
		}

		sre, err := database.NewServiceRequestEvent(psqlClient).GetLatestStepEvent(requestId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		err = database.NewServiceRequest(client).UpdateStatus(requestId, models.CANCELLED)

		// TODO: discuss how to handle this error
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		if models.IsCancellablePipelineStepType(sre.StepType) {
			userId := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).RegisteredClaims.Subject
			// Log step cancelled event
			serviceRequestEvent := database.NewServiceRequestEvent(psqlClient)
			err = serviceRequestEvent.Create(&models.ServiceRequestEventModel{
				EventType:        models.STEP_COMPLETED,
				ServiceRequestId: sr.Id.Hex(),
				StepName:         sre.StepName,
				CreatedBy:        userId,
				StepType:         sre.StepType,
			})
			if err != nil {
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
				return
			}
		}

		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleUpdateServiceRequest(logger logger.ServerLogger, client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srm, err := decode[models.ServiceRequestModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		sr, err := database.NewServiceRequest(client).GetById(requestId)
		if err != nil {
			logger.Error(fmt.Sprintf("%s %s not found", "service request", requestId))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		status := sr.Status
		if status != models.NOT_STARTED {
			logger.Error(fmt.Sprintf("failed to %s service request %s: %s", "update", requestId, "execution has been started"))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrServiceRequestAlreadyStarted, http.StatusBadRequest))
			return
		}
		_, err = database.NewServiceRequest(client).UpdateById(requestId, &srm)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleStartServiceRequest(logger logger.ServerLogger, client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		srm, err := database.NewServiceRequest(client).GetById(requestId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		if srm.Status != models.NOT_STARTED {
			logger.Error(fmt.Sprintf("failed to %s service request %s: %s", "start", requestId, "execution has already been started"))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrServiceRequestAlreadyStarted, http.StatusBadRequest))
			return
		}
		event.FireAsync(events.NewNewServiceRequestEvent(srm))
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleApproveServiceRequest(logger logger.ServerLogger, client *mongo.Client, psqlClient *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		serviceRequestId := params["requestId"]
		serviceRequest, err := database.NewServiceRequest(client).GetById(serviceRequestId)
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("%s %s not found", "service request", serviceRequestId))
			encode(w, r, http.StatusNotFound, newHandlerError(ErrInvalidServiceRequestId, http.StatusNotFound))
		}
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		if serviceRequest.Status != models.PENDING {
			logger.Error(fmt.Sprintf("unable to approve service request %s: request is not pending approval", serviceRequestId))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrFailedToApproveServiceRequest, http.StatusBadRequest))
			return
		}

		latestStep, err := database.NewServiceRequestEvent(psqlClient).GetLatestStepEvent(serviceRequestId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		if latestStep.StepType != models.WaitForApprovalStep {
			// This check should be redundant but just in case
			// Because if the status of the request is pending, the latest step should be an approval step
			// Unless we have a bug in the code somewhere
			logger.Error(fmt.Sprintf("unable to approve service request %s: latest step is not approval step ", serviceRequestId))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrFailedToApproveServiceRequest, http.StatusBadRequest))
			return
		}

		userId := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).RegisteredClaims.Subject

		user, err := database.NewUser(psqlClient).GetUserById(userId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		logger.Info(fmt.Sprintf("approving service request \"%s\" at step \"%s\", performed by %s", serviceRequestId, latestStep.StepName, user.Name))
		// TODO: figure out how to pass the step result prior to the approval to the next step
		event.FireAsync(events.NewStepCompletedEvent(latestStep.StepName, serviceRequest.Id.Hex(), userId, nil, nil))
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleRejectServiceRequest(logger logger.ServerLogger, client *mongo.Client, psqlClient *sql.DB) http.Handler {
	type requestBody struct {
		Remarks string `json:"remarks"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		serviceRequestId := params["requestId"]
		serviceRequest, err := database.NewServiceRequest(client).GetById(serviceRequestId)
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("%s %s not found", "service request", serviceRequestId))
			encode(w, r, http.StatusNotFound, newHandlerError(ErrInvalidServiceRequestId, http.StatusNotFound))
		}
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		if serviceRequest.Status != models.PENDING {
			logger.Error(fmt.Sprintf("unable to reject service request %s: request is not pending approval", serviceRequestId))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrFailedToApproveServiceRequest, http.StatusBadRequest))
			return
		}

		latestStep, err := database.NewServiceRequestEvent(psqlClient).GetLatestStepEvent(serviceRequestId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		if latestStep.StepType != models.WaitForApprovalStep {
			// This check should be redundant but just in case
			logger.Error(fmt.Sprintf("unable to approve service request %s: latest step is not approval step ", serviceRequestId))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrFailedToApproveServiceRequest, http.StatusBadRequest))
			return
		}

		body, err := decode[requestBody](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}

		userId := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).RegisteredClaims.Subject

		user, err := database.NewUser(psqlClient).GetUserById(userId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		logger.Info(fmt.Sprintf("rejecting service request \"%s\" at step \"%s\", performed by %s", serviceRequestId, latestStep.StepName, user.Name))

		// Update Service Request Status
		err = database.NewServiceRequest(client).UpdateStatus(serviceRequestId, models.FAILED)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		// Add that SR is rejected at start of remarks
		failedEventRemarks := fmt.Sprintf("%s\n%s\n%s", fmt.Sprintf("Rejected by %s", user.Name), "Remarks by admin:", body.Remarks)
		event.FireAsync(events.NewStepFailedEvent(latestStep.StepName, serviceRequest, userId, failedEventRemarks, nil))
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleCreatePipeline(logger logger.ServerLogger, client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pipeline, err := decode[models.PipelineModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		if pipeline.Form.Fields == nil {
			logger.Error("missing form field")
			encode(w, r, http.StatusUnprocessableEntity, newHandlerError(ErrJsonParseError, http.StatusUnprocessableEntity))
			return
		}

		for _, element := range pipeline.Form.Fields {
			err = validation.ValidateFormField(element)
			if err != nil {
				logger.Error(fmt.Sprintf("invalid form field: %s", err))
				encode(w, r, http.StatusBadRequest, newHandlerError(err, http.StatusUnprocessableEntity))
				return
			}
		}

		userId := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).RegisteredClaims.Subject
		pipeline.CreatedOn = time.Now()
		pipeline.Version = 1
		pipeline.UserId = userId
		res, err := database.NewPipeline(client).Create(&pipeline)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrPipelineCreateFail, http.StatusInternalServerError))
			return
		}
		insertedId, _ := res.InsertedID.(primitive.ObjectID)
		pipeline.Id = insertedId
		encode(w, r, http.StatusCreated, pipeline)
	})
}

func handleGetPipeline(logger logger.ServerLogger, client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		pipelineId := vars["pipelineId"]
		pipeline, err := database.NewPipeline(client).GetById(pipelineId)
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("%s %s not found", "pipeline", pipelineId))
		}
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		encode(w, r, http.StatusOK, pipeline)
	})
}

func handleGetAllPipelines(logger logger.ServerLogger, client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orgId, err := extractQueryParam[int](r.URL.Query(), "org_id", false, -1, integerConverter)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to extract org_id from query params: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidOrganizationId, http.StatusBadRequest))
			return
		}
		pipelines, err := database.NewPipeline(client).GetAllByOrgId(orgId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		encode(w, r, http.StatusOK, pipelines)
	})
}

func handleUserLogin(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).RegisteredClaims.Subject
		user, err := database.NewUser(client).GetUserById(userId)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error(fmt.Sprintf("user %s has not been created", userId))
			encode(w, r, http.StatusNotFound, newHandlerError(ErrInvalidUserId, http.StatusNotFound))
			return
		} else if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		logger.Info(fmt.Sprintf("user %s logged in", userId))
		encode(w, r, http.StatusOK, user)
	})
}

func handleCreateUser(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		um, err := decode[models.UserModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		userId := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).RegisteredClaims.Subject
		um.UserId = userId

		err = helper.GetAuth0UserDetailsForUser(&um)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while retrieving user details: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		user, err := database.NewUser(client).Create(&um)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to create user: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		logger.Info(fmt.Sprintf("user %s created", user.UserId))
		encode(w, r, http.StatusOK, user)
	})
}

func handleCreateOrganization(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		om, err := decode[models.OrganizationModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		// Assign user id retrieved from token to organization
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		userId := token.RegisteredClaims.Subject
		om.Owner = userId
		org, err := database.NewOrganization(client).Create(&om)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrOrganizationCreateFail, http.StatusInternalServerError))
			return
		}
		logger.Info(fmt.Sprintf("%s %s created", "organization", fmt.Sprint(org.OrgId)))
		encode(w, r, http.StatusCreated, org)
	})
}

func handleCreateMembership(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mm, err := decode[models.MembershipModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		membership, err := database.NewMembership(client).Create(&mm)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrMembershipCreateFail, http.StatusInternalServerError))
			return
		}
		logger.Info(fmt.Sprintf("%s %s created", "membership", fmt.Sprintf("%s-%d", mm.UserId, mm.OrgId)))
		encode(w, r, http.StatusCreated, membership)
	})
}

func handleGetServiceRequestsByUserAndOrganization(logger logger.ServerLogger, client *mongo.Client) http.Handler {
	type ResponseBodyMetadata struct {
		TotalCount int `json:"total_count"`
	}
	type ResponseBody struct {
		Data     []*models.ServiceRequestModel `json:"data"`
		Metadata ResponseBodyMetadata          `json:"metadata"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		userId := token.RegisteredClaims.Subject

		orgId, err := extractQueryParam[int](r.URL.Query(), "org_id", false, -1, integerConverter)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to extract org_id from query params: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidOrganizationId, http.StatusBadRequest))
			return
		}

		statusFilters, err := extractQueryParam[string](r.URL.Query(), "status", false, "", stringConverter)
		queryFilters := database.GetServiceRequestFilters{
			UserId: userId,
		}
		if statusFilters != "" {
			statuses := strings.Split(statusFilters, ",")
			for _, status := range statuses {
				if !models.ValidateServiceRequestStatus(status) {
					logger.Error(fmt.Sprintf("invalid status filter: %s", status))
					encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidServiceRequestStatus, http.StatusBadRequest))
					return
				}
			}
			queryFilters.Statuses = strings.Split(statusFilters, ",")
		}
		logger.Info(fmt.Sprintf("query: %v", r.URL.Query()))
		pageParam, err := extractQueryParam[int](r.URL.Query(), "page", false, 1, integerConverter)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to extract page from query params: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrPaginationParamsError, http.StatusBadRequest))
			return
		}
		pageSizeParam, err := extractQueryParam[int](r.URL.Query(), "page_size", false, 10, integerConverter)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to extract page_size from query params: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrPaginationParamsError, http.StatusBadRequest))
			return
		}
		if pageParam < 1 || pageSizeParam < 1 {
			logger.Error(fmt.Sprintf("invalid page or page_size: page=%d, page_size=%d", pageParam, pageSizeParam))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrPaginationParamsError, http.StatusBadRequest))
			return
		}
		logger.Info(fmt.Sprintf("querying for service requests: org_id=%d, query_filters=%v, page=%d, page_size=%d", orgId, queryFilters, pageParam, pageSizeParam))
		result, err := database.NewServiceRequest(client).GetAllServiceRequestByOrg(orgId, queryFilters, database.Pagination{
			Page: pageParam, PageSize: pageSizeParam,
		})
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		response := ResponseBody{
			Data: result.Data,
			Metadata: ResponseBodyMetadata{
				TotalCount: result.TotalCount,
			},
		}
		encode(w, r, http.StatusOK, response)
	})
}

func handleGetServiceRequestsForAdminByOrganization(logger logger.ServerLogger, client *mongo.Client) http.Handler {
	type ResponseBodyMetadata struct {
		TotalCount int `json:"total_count"`
	}
	type ResponseBody struct {
		Data     []*models.ServiceRequestModel `json:"data"`
		Metadata ResponseBodyMetadata          `json:"metadata"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orgId, err := extractQueryParam[int](r.URL.Query(), "org_id", false, -1, integerConverter)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to extract org_id from query params: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidOrganizationId, http.StatusBadRequest))
			return
		}

		statusFilters, err := extractQueryParam[string](r.URL.Query(), "status", false, "", stringConverter)
		queryFilters := database.GetServiceRequestFilters{}
		if statusFilters != "" {
			statuses := strings.Split(statusFilters, ",")
			for _, status := range statuses {
				if !models.ValidateServiceRequestStatus(status) {
					logger.Error(fmt.Sprintf("invalid status filter: %s", status))
					encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidServiceRequestStatus, http.StatusBadRequest))
					return
				}
			}
			queryFilters.Statuses = strings.Split(statusFilters, ",")
		}

		pageParam, err := extractQueryParam[int](r.URL.Query(), "page", false, 1, integerConverter)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to extract page from query params: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrPaginationParamsError, http.StatusBadRequest))
			return
		}
		pageSizeParam, err := extractQueryParam[int](r.URL.Query(), "page_size", false, 10, integerConverter)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to extract page_size from query params: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrPaginationParamsError, http.StatusBadRequest))
			return
		}
		if pageParam < 1 || pageSizeParam < 1 {
			logger.Error(fmt.Sprintf("invalid page or page_size: page=%d, page_size=%d", pageParam, pageSizeParam))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrPaginationParamsError, http.StatusBadRequest))
			return
		}
		logger.Info(fmt.Sprintf("querying for service requests: org_id=%d, query_filters=%v, page=%d, page_size=%d", orgId, queryFilters, pageParam, pageSizeParam))
		result, err := database.NewServiceRequest(client).GetAllServiceRequestByOrg(orgId, queryFilters, database.Pagination{
			Page: pageParam, PageSize: pageSizeParam,
		})
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		response := ResponseBody{
			Data: result.Data,
			Metadata: ResponseBodyMetadata{
				TotalCount: result.TotalCount,
			},
		}
		encode(w, r, http.StatusOK, response)
	})
}

func handleGetUserById(logger logger.ServerLogger, client *sql.DB) http.Handler {
	type ResponseBodyMembership struct {
		OrgId    int         `json:"org_id"`
		Role     models.Role `json:"role"`
		JoinedOn time.Time   `json:"joined_on"`
	}
	type ResponseBody struct {
		UserId           string                   `json:"user_id"`
		Name             string                   `json:"name"`
		IdentityProvider string                   `json:"identity_provider"`
		CreatedOn        time.Time                `json:"created_on"`
		Memberships      []ResponseBodyMembership `json:"memberships"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["userId"]
		user, err := database.NewUser(client).GetUserById(userId)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error(fmt.Sprintf("%s %s not found", "user", userId))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInvalidUserId, http.StatusInternalServerError))
			return
		}
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUserRetrieve, http.StatusInternalServerError))
			return
		}
		memberships, err := database.NewMembership(client).GetUserMemberships(userId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrMembershipRetrieve, http.StatusInternalServerError))
			return
		}
		response := ResponseBody{
			UserId:           user.UserId,
			Name:             user.Name,
			IdentityProvider: user.IdentityProvider,
			CreatedOn:        user.CreatedOn,
		}
		for _, membership := range memberships {
			response.Memberships = append(response.Memberships, ResponseBodyMembership{
				OrgId:    membership.OrgId,
				Role:     membership.Role,
				JoinedOn: membership.JoinedOn,
			})
		}
		encode(w, r, http.StatusOK, response)
	})
}

func handleGetAllUsers(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := database.NewUser(client).GetAllUsers()
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUserRetrieve, http.StatusInternalServerError))
			return
		}
		encode(w, r, http.StatusOK, users)
	})
}

func handleOwnershipTransfer(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ownerId := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).RegisteredClaims.Subject
		targetOwner, err := decode[models.MembershipModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}

		_, err = database.NewMembership(client).GetMembershipByUserAndOrgId(targetOwner.UserId, targetOwner.OrgId)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("target user for transfer is not part of organisation")
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrNotOrgMember, http.StatusBadRequest))
			return
		} else if err != nil {
			logger.Error(fmt.Sprintf("unable to verify if target belongs to org: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		err = database.NewMembership(client).TransferOwnership(ownerId, targetOwner.UserId, targetOwner.OrgId)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to update membership: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrMembershipUpdateFail, http.StatusInternalServerError))
			return
		}

		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleUpdateMembership(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mm, err := decode[models.MembershipModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		_, err = database.NewMembership(client).UpdateUserMembership(&mm)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to update membership: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrMembershipUpdateFail, http.StatusInternalServerError))
			return
		}

		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleDeleteMembership(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mm, err := decode[models.MembershipModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		_, err = database.NewMembership(client).DeleteUserMembership(&mm)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to delete membership: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrMembershipDeleteFail, http.StatusInternalServerError))
			return
		}
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleLeaveOrganization(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		orgId, err := strconv.Atoi(vars["orgId"])
		if err != nil {
			logger.Error(fmt.Sprintf("unable to parse orgId: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidOrganizationId, http.StatusBadRequest))
			return
		}
		userId := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims).RegisteredClaims.Subject
		mm := models.MembershipModel{
			UserId: userId,
			OrgId:  orgId,
		}

		_, err = database.NewMembership(client).DeleteUserMembership(&mm)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to delete membership: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrMembershipDeleteFail, http.StatusInternalServerError))
			return
		}
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleGetStepExecutionLogs(l logger.ServerLogger, psqlClient *sql.DB) http.Handler {
	type ResponseBody struct {
		StepName   string   `json:"step_name"`
		Logs       []string `json:"logs"`
		EndOfLogs  bool     `json:"end_of_logs"`
		NextOffset int      `json:"next_offset"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Initialise variables
		vars := mux.Vars(r)
		stepName := vars["stepName"]
		serviceRequestId := vars["requestId"]
		offset, err := extractQueryParam[int](r.URL.Query(), "offset", false, 0, integerConverter)
		if err != nil {
			l.Error(fmt.Sprintf("unable to extract offset from query params: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidOffset, http.StatusBadRequest))
			return
		}
		if offset < 0 {
			l.Error(fmt.Sprintf("invalid offset: %d", offset))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidOffset, http.StatusBadRequest))
			return
		}

		// Fetch log file
		f, err := logger.FindExecutorLogFile(serviceRequestId, stepName)
		if err != nil {
			l.Error(fmt.Sprintf("unable to get log file: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		defer f.Close()

		// Scan logs from offset
		logs, err := util.ScanFileFromOffset(f, offset)
		if err != nil {
			l.Error(fmt.Sprintf("unable to scan log file: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		// Prepare response
		stepEvent, err := database.NewServiceRequestEvent(psqlClient).GetStepLatestEvent(serviceRequestId, stepName)
		if err != nil {
			l.Error(fmt.Sprintf("unable to get latest event: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		endOfLogs := stepEvent.EventType != models.STEP_RUNNING
		response := ResponseBody{
			StepName:  stepName,
			Logs:      logs,
			EndOfLogs: endOfLogs,
		}
		if endOfLogs {
			response.NextOffset = -1
		} else {
			response.NextOffset = offset + len(logs)
		}

		encode(w, r, http.StatusOK, response)
	})
}

func handleGetOrganizationMembers(logger logger.ServerLogger, client *sql.DB) http.Handler {
	type ResponseBody struct { // Response body when org_id is provided
		OrgId   int                                    `json:"org_id"`
		Members []*database.GetAllUsersByOrgIdResponse `json:"members"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		orgId, err := strconv.Atoi(vars["orgId"])
		if err != nil {
			logger.Error(fmt.Sprintf("unable to parse orgId: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidOrganizationId, http.StatusBadRequest))
			return
		}
		users, err := database.NewOrganization(client).GetAllUsersByOrgId(orgId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUserRetrieve, http.StatusInternalServerError))
			return
		}
		response := ResponseBody{
			OrgId:   orgId,
			Members: users,
		}
		encode(w, r, http.StatusOK, response)
	})

}
func handleGetOrganizationsForUser(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		userId := token.RegisteredClaims.Subject

		orgs, err := database.NewOrganization(client).GetAllOrgsByUserId(userId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrOrganizationRetrieve, http.StatusInternalServerError))
			return
		}
		encode(w, r, http.StatusOK, orgs)
	})
}

func handleUpdateOrganization(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		om, err := decode[models.OrganizationModel](r)

		if err != nil {
			logger.Error(fmt.Sprintf("unable to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		userId := token.RegisteredClaims.Subject

		_, err = database.NewOrganization(client).UpdateOrgName(om.Name, om.OrgId, userId)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to update organization: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrOrganizationUpdateFail, http.StatusInternalServerError))
			return
		}
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleDeleteOrganization(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		om, err := decode[models.OrganizationModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		userId := token.RegisteredClaims.Subject
		_, err = database.NewOrganization(client).DeleteOrg(om.OrgId, userId)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to delete organization: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrOrganizationDeleteFail, http.StatusInternalServerError))
			return
		}
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleGetMembershipsForUser(logger logger.ServerLogger, client *sql.DB) http.Handler {
	type ResponseBodyMembership struct {
		OrgId    int         `json:"org_id"`
		Role     models.Role `json:"role"`
		JoinedOn time.Time   `json:"joined_on"`
	}
	type ResponseBody struct {
		UserId      string                   `json:"user_id"`
		Memberships []ResponseBodyMembership `json:"memberships"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		userId := token.RegisteredClaims.Subject

		memberships, err := database.NewMembership(client).GetUserMemberships(userId)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrMembershipRetrieve, http.StatusInternalServerError))
			return
		}
		response := ResponseBody{
			UserId: userId,
		}
		for _, membership := range memberships {
			response.Memberships = append(response.Memberships, ResponseBodyMembership{
				OrgId:    membership.OrgId,
				Role:     membership.Role,
				JoinedOn: membership.JoinedOn,
			})
		}
		encode(w, r, http.StatusOK, response)
	})
}
