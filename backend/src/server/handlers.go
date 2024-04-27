package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gookit/event"
	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/events"
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
	r.Handle("/api/service_request", isAuthenticated(getOrgIdFromQuery(isOrgMember(s.psqlClient, handleGetServiceRequestsByOrganization(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("GET")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgMember(s.psqlClient, handleGetServiceRequest(s.logger, s.mongoClient, s.psqlClient), s.logger), s.logger), s.logger)).Methods("GET")
	r.Handle("/api/service_request", isAuthenticated(getOrgIdFromRequestBody(isOrgMember(s.psqlClient, handleCreateServiceRequest(s.logger, s.mongoClient, s.psqlClient), s.logger), s.logger), s.logger)).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(getOrgIdFromRequestBody(isOrgMember(s.psqlClient, handleUpdateServiceRequest(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}/cancel", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgMember(s.psqlClient, handleCancelStartedServiceRequest(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("PUT")
	r.Handle("/api/service_request/{requestId}/start", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgMember(s.psqlClient, handleStartServiceRequest(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("PUT")
	r.Handle("/api/service_request/{requestId}/approve", isAuthenticated(getOrgIdFromRequestBody(isOrgAdmin(s.psqlClient, handleApproveServiceRequest(s.logger, s.mongoClient), s.logger), s.logger), s.logger)).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}/logs/{stepName}", isAuthenticated(getOrgIdUsingSrId(s.mongoClient, isOrgMember(s.psqlClient, handleGetStepExecutionLogs(s.logger, s.psqlClient), s.logger), s.logger), s.logger)).Methods("GET")

	// Pipeline
	// TODO: @joshtyf need to integrate orgId in some way into these routes or the pipeline model, esp for the post method.
	r.Handle("/api/pipeline", isAuthenticated(handleGetAllPipelines(s.logger, s.mongoClient), s.logger)).Methods("GET")
	r.Handle("/api/pipeline/{pipelineId}", isAuthenticated(handleGetPipeline(s.logger, s.mongoClient), s.logger)).Methods("GET")
	r.Handle("/api/pipeline", isAuthenticated(validateCreatePipelineRequest(handleCreatePipeline(s.logger, s.mongoClient), s.logger), s.logger)).Methods("POST").Headers("Content-Type", "application/json")

	// User
	r.Handle("/api/user", isAuthenticated(handleGetUserById(s.logger, s.psqlClient), s.logger)).Methods("GET")
	// TODO: review the need for this route
	// r.Handle("/api/user/{userId}", isAuthenticated(handleGetUserById(s.psqlClient))).Methods("GET")
	r.Handle("/api/login", isAuthenticated(handleUserLogin(s.logger, s.psqlClient), s.logger)).Methods("POST").Headers("Content-Type", "application/json")

	// Organization
	r.Handle("/api/organization", isAuthenticated(handleCreateOrganization(s.logger, s.psqlClient), s.logger)).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/organization", isAuthenticated(handleGetOrganizationsForUser(s.logger, s.psqlClient), s.logger)).Methods("GET")

	// Membership
	r.Handle("/api/membership", isAuthenticated(getOrgIdFromRequestBody(validateMembershipChange(s.psqlClient, isOrgAdmin(s.psqlClient, handleCreateMembership(s.logger, s.psqlClient), s.logger), s.logger), s.logger), s.logger)).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/membership", isAuthenticated(getOrgIdFromRequestBody(validateMembershipChange(s.psqlClient, isOrgAdmin(s.psqlClient, handleUpdateMembership(s.logger, s.psqlClient), s.logger), s.logger), s.logger), s.logger)).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/membership", isAuthenticated(getOrgIdFromRequestBody(validateMembershipChange(s.psqlClient, isOrgOwner(s.psqlClient, handleDeleteMembership(s.logger, s.psqlClient), s.logger), s.logger), s.logger), s.logger)).Methods("DELETE").Headers("Content-Type", "application/json")
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
		ApprovedBy   string           `json:"approved_by"`
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
				ApprovedBy:   event.ApprovedBy,
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

func handleCancelStartedServiceRequest(logger logger.ServerLogger, client *mongo.Client) http.Handler {
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
		if status != models.RUNNING {
			logger.Error(fmt.Sprintf("failed to %s service request %s: %s", "cancel", requestId, "execution has been completed"))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrServiceRequestAlreadyCompleted, http.StatusBadRequest))
			return
		}

		if status == models.NOT_STARTED {
			logger.Error(fmt.Sprintf("failed to %s service request %s: %s", "cancel", requestId, "execution has not been started"))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrServiceRequestNotStarted, http.StatusBadRequest))
			return
		}

		// TODO: implement cancellation of sr

		err = database.NewServiceRequest(client).UpdateStatus(requestId, models.CANCELLED)

		// TODO: discuss how to handle this error
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
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

func handleApproveServiceRequest(logger logger.ServerLogger, client *mongo.Client) http.Handler {
	type requestBody struct {
		StepName string `json:"step_name"`
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
		body, err := decode[requestBody](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		pipeline, err := database.NewPipeline(client).GetById(serviceRequest.PipelineId)
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("%s %s not found", "pipeline", serviceRequest.PipelineId))
			encode(w, r, http.StatusNotFound, newHandlerError(ErrInvalidPipelineId, http.StatusNotFound))
		}
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		waitForApprovalStep := pipeline.GetPipelineStep(body.StepName)

		if waitForApprovalStep == nil {
			logger.Error(fmt.Sprintf("missing pipeline step: %s", body.StepName))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		if waitForApprovalStep.StepType != models.WaitForApprovalStep {
			logger.Error(fmt.Sprintf("invalid step type: expected %s got %s", models.WaitForApprovalStep, waitForApprovalStep.StepType))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrWrongStepType, http.StatusBadRequest))
			return
		}
		// TODO: figure out how to pass the step result prior to the approval to the next step
		event.FireAsync(events.NewStepCompletedEvent(waitForApprovalStep, serviceRequest, nil, nil))
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

		pipeline.CreatedOn = time.Now()
		pipeline.Version = 1
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
		pipelines, err := database.NewPipeline(client).GetAll()
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		encode(w, r, http.StatusOK, pipelines)
	})
}

// NOTE: handler and data functions used in here are subjected to change depending on if frontend is able to validate that user has been previously registered in auth0
func handleUserLogin(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		um, err := decode[models.UserModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		_, err = database.NewUser(client).GetUserById(um.UserId)
		if errors.Is(err, sql.ErrNoRows) {
			user, err := database.NewUser(client).Create(&um)
			if err != nil {
				logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUserCreateFail, http.StatusInternalServerError))
				return
			}

			logger.Info(fmt.Sprintf("%s %s created", "user", user.UserId))
			encode[any](w, r, http.StatusCreated, nil)
			return
		} else if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		logger.Info(fmt.Sprintf("user %s logged in", um.UserId))
		encode[any](w, r, http.StatusOK, nil)
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

func handleGetServiceRequestsByOrganization(logger logger.ServerLogger, client *mongo.Client) http.Handler {
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

		logger.Info(fmt.Sprintf("querying for service requests: user_id=%s org_id=%d, query_filters=%v", userId, orgId, queryFilters))
		allsr, err := database.NewServiceRequest(client).GetAllServiceRequestsForOrgId(userId, orgId, queryFilters)
		if err != nil {
			logger.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		encode(w, r, http.StatusOK, allsr)
	})
}

func handleGetUserById(logger logger.ServerLogger, client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user_id := vars["userId"]
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		userId := token.RegisteredClaims.Subject

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

		encode(w, r, http.StatusCreated, user)
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
		stepEvent, err := database.NewServiceRequestEvent(psqlClient).GetLatestEvent(serviceRequestId, stepName)
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
