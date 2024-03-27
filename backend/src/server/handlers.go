package server

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gookit/event"
	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/events"
	"github.com/joshtyf/flowforge/src/logger"
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
	logger      *logger.Logger
	psqlClient  *sql.DB
	mongoClient *mongo.Client
}

func NewServerHandler(psqlClient *sql.DB, mongoCLient *mongo.Client, logger *logger.Logger) *ServerHandler {
	return &ServerHandler{
		psqlClient:  psqlClient,
		mongoClient: mongoCLient,
		logger:      logger,
	}
}

func (s *ServerHandler) registerRoutes(r *mux.Router) {
	// Health Check
	r.Handle("/api/healthcheck", handleHealthCheck()).Methods("GET")

	// Service Request
	r.Handle("/api/service_request", isAuthenticated(isOrgMember(s.mongoClient, s.psqlClient, handleGetAllServiceRequestsForOrganisation(s.mongoClient)))).Methods("GET")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(isOrgMember(s.mongoClient, s.psqlClient, handleGetServiceRequest(s.mongoClient, s.psqlClient)))).Methods("GET")
	r.Handle("/api/service_request", isAuthenticated(isOrgMember(s.mongoClient, s.psqlClient, handleCreateServiceRequest(s.mongoClient, s.psqlClient)))).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(isOrgMember(s.mongoClient, s.psqlClient, handleUpdateServiceRequest(s.mongoClient)))).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}/cancel", isAuthenticated(isOrgMember(s.mongoClient, s.psqlClient, handleCancelStartedServiceRequest(s.mongoClient)))).Methods("PUT")
	r.Handle("/api/service_request/{requestId}/start", isAuthenticated(isOrgMember(s.mongoClient, s.psqlClient, handleStartServiceRequest(s.mongoClient)))).Methods("PUT")
	r.Handle("/api/service_request/{requestId}/approve", isAuthenticated(isOrgAdmin(s.mongoClient, s.psqlClient, handleApproveServiceRequest(s.mongoClient)))).Methods("POST").Headers("Content-Type", "application/json")

	// Pipeline
	r.Handle("/api/pipeline", isAuthenticated(isOrgAdmin(s.mongoClient, s.psqlClient, handleGetAllPipelines(s.mongoClient)))).Methods("GET")
	r.Handle("/api/pipeline/{pipelineId}", isAuthenticated(isOrgAdmin(s.mongoClient, s.psqlClient, handleGetPipeline(s.mongoClient)))).Methods("GET")
	r.Handle("/api/pipeline", isAuthenticated(isOrgAdmin(s.mongoClient, s.psqlClient, validateCreatePipelineRequest(handleCreatePipeline(s.mongoClient))))).Methods("POST").Headers("Content-Type", "application/json")

	// User
	r.Handle("/api/user/{userId}", isAuthenticated(handleGetUserById(s.psqlClient))).Methods("Get")
	r.Handle("/api/login", isAuthenticated(handleUserLogin(s.psqlClient))).Methods("POST").Headers("Content-Type", "application/json")

	// Organisation
	r.Handle("/api/organisation", isAuthenticated(handleCreateOrganisation(s.psqlClient))).Methods("POST").Headers("Content-Type", "application/json")

	// Membership
	r.Handle("/api/membership", isAuthenticated(isOrgAdmin(s.mongoClient, s.psqlClient, handleCreateMembership(s.psqlClient)))).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/membership", isAuthenticated(isOrgAdmin(s.mongoClient, s.psqlClient, handleUpdateMembership(s.psqlClient)))).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/membership", isAuthenticated(isOrgOwner(s.mongoClient, s.psqlClient, handleDeleteMembership(s.psqlClient)))).Methods("DELETE").Headers("Content-Type", "application/json")
}

func handleHealthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if serverHealthy() {
			logger.Info("[HealthCheck] Server working!", nil)
			encode(w, r, http.StatusOK, "Server working!")
			return
		}
		encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
	})
}

func handleGetAllServiceRequest(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allsr, err := database.NewServiceRequest(client).GetAll()
		if err != nil {
			logger.Error("[GetAllServiceRequest] Error retrieving all service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		logger.Info("[GetAllServiceRequest] Successfully retrieved service requests", map[string]interface{}{"count": len(allsr)})
		encode(w, r, http.StatusOK, allsr)
	})
}

func handleGetServiceRequest(mongoClient *mongo.Client, psqlClient *sql.DB) http.Handler {
	type ResponseBodyStep struct {
		Name         string           `json:"name"`
		Status       models.EventType `json:"status"`
		UpdatedAt    time.Time        `json:"updated_at"`
		ApprovedBy   string           `json:"approved_by"`
		NextStepName string           `json:"next_step_name"`
	}
	type ResponseBody struct {
		ServiceRequest *models.ServiceRequestModel `json:"service_request"`
		Steps          map[string]ResponseBodyStep `json:"steps"`
		FirstStepName  string                      `json:"first_step_name"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		sr, err := database.NewServiceRequest(mongoClient).GetById(requestId)
		if err != nil {
			logger.Error("[GetServiceRequest] Unable to retrieve service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		pipeline, err := database.NewPipeline(mongoClient).GetById(sr.PipelineId)
		if err != nil {
			logger.Error("[GetServiceRequest] Unable to retrieve pipeline", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		sre, err := database.NewServiceRequestEvent(psqlClient).GetStepsLatestEvent(requestId)
		if err != nil {
			logger.Error("[GetServiceRequest] Unable to retrieve latest service request events", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		steps := make(map[string]ResponseBodyStep, len(sre))
		for _, event := range sre {
			step := pipeline.GetPipelineStep(event.StepName)
			if step == nil {
				logger.Error("[GetServiceRequest] Found a step that exists in events log but not in pipeline", map[string]interface{}{"step": event.StepName})
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
		}
		logger.Info("[GetServiceRequest] Successfully retrieved service request", map[string]interface{}{"response": response})
		encode(w, r, http.StatusOK, response)
	})
}

func handleCreateServiceRequest(mongoClient *mongo.Client, psqlClient *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srm, err := decode[models.ServiceRequestModel](r)
		if err != nil {
			logger.Error("[CreateServiceRequest] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}

		pipeline, err := database.NewPipeline(mongoClient).GetById(srm.PipelineId)
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("[CreateServiceRequest] Invalid pipeline id, no matching pipeline found", map[string]interface{}{"pipelineId": srm.PipelineId})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidPipelineId, http.StatusBadRequest))
			return
		} else if err != nil {
			logger.Error("[CreateServiceRequest] Error getting pipeline", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		srm.CreatedOn = time.Now()
		srm.LastUpdated = time.Now()
		srm.Status = models.NotStarted

		res, err := database.NewServiceRequest(mongoClient).Create(&srm)
		if err != nil {
			logger.Error("[CreateServiceRequest] Error creating service request", map[string]interface{}{"err": err})
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
				logger.Error("[CreateServiceRequest] Error creating step not started event", map[string]interface{}{"err": err})
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
				return
			}
		}
		logger.Info("[CreateServiceRequest] Successfully created service request", map[string]interface{}{"sr": srm})
		insertedId, _ := res.InsertedID.(primitive.ObjectID)
		srm.Id = insertedId
		encode(w, r, http.StatusCreated, srm)
	})
}

func handleCancelStartedServiceRequest(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		sr, err := database.NewServiceRequest(client).GetById(requestId)
		if err != nil {
			logger.Error("[CancelStartedServiceRequest] Error getting service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		status := sr.Status
		if status != models.Pending && status != models.Running {
			logger.Error("[CancelStartedServiceRequest] Unable to cancel service request as execution has been completed", nil)
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrServiceRequestAlreadyCompleted, http.StatusBadRequest))
			return
		}

		if status == models.NotStarted {
			logger.Error("[CancelStartedServiceRequest] Unable to cancel service request as execution has not been started", nil)
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrServiceRequestNotStarted, http.StatusBadRequest))
			return
		}

		// TODO: implement cancellation of sr

		err = database.NewServiceRequest(client).UpdateStatus(requestId, models.Canceled)

		// TODO: discuss how to handle this error
		if err != nil {
			logger.Error("[CancelStartedServiceRequest] Unable to update service request status", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		logger.Info("[CancelStartedServiceRequest] Successfully canceled started service request", nil)
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleUpdateServiceRequest(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srm, err := decode[models.ServiceRequestModel](r)
		if err != nil {
			logger.Error("[UpdateServiceRequest] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		sr, err := database.NewServiceRequest(client).GetById(requestId)
		if err != nil {
			logger.Error("[UpdateServiceRequest] Error getting service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		status := sr.Status
		if status != models.NotStarted {
			logger.Error("[UpdateServiceRequest] Unable to update service request as service request has been started", nil)
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrServiceRequestAlreadyStarted, http.StatusBadRequest))
			return
		}
		res, err := database.NewServiceRequest(client).UpdateById(requestId, &srm)
		if err != nil {
			logger.Error("[UpdateServiceRequest] Error updating service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		logger.Info("[UpdateServiceRequest] Successfully updated service request", map[string]interface{}{"count": res.ModifiedCount})
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleStartServiceRequest(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		srm, err := database.NewServiceRequest(client).GetById(requestId)
		if err != nil {
			logger.Error("[StartServiceRequest] Unable retrieve service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		if srm.Status != models.NotStarted {
			logger.Error("[StartServiceRequest] Service request has already been started", nil)
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrServiceRequestAlreadyStarted, http.StatusBadRequest))
			return
		}
		event.FireAsync(events.NewNewServiceRequestEvent(srm))
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleApproveServiceRequest(client *mongo.Client) http.Handler {
	type requestBody struct {
		StepName string `json:"step_name"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		serviceRequestId := params["requestId"]
		serviceRequest, err := database.NewServiceRequest(client).GetById(serviceRequestId)
		if err != nil {
			logger.Error("[ApproveServiceRequest] Unable to get service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		body, err := decode[requestBody](r)
		if err != nil {
			logger.Error("[ApproveServiceRequest] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		pipeline, err := database.NewPipeline(client).GetById(serviceRequest.PipelineId)
		if err != nil {
			logger.Error("[ApproveServiceRequest] Unable to get pipeline", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		waitForApprovalStep := pipeline.GetPipelineStep(body.StepName)

		if waitForApprovalStep == nil {
			logger.Error("[ApproveServiceRequest] Unable to get wait for approval step", map[string]interface{}{"step": body.StepName})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		if waitForApprovalStep.StepType != models.WaitForApprovalStep {
			logger.Error("[ApproveServiceRequest] Step is not a wait for approval step", map[string]interface{}{"step": body.StepName})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrWrongStepType, http.StatusBadRequest))
			return
		}
		// TODO: figure out how to pass the step result prior to the approval to the next step
		event.FireAsync(events.NewStepCompletedEvent(waitForApprovalStep, serviceRequest, nil, nil))
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleCreatePipeline(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pipeline, err := decode[models.PipelineModel](r)
		if err != nil {
			logger.Error("[CreatePipeline] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}

		pipeline.CreatedOn = time.Now()
		pipeline.Version = 1
		res, err := database.NewPipeline(client).Create(&pipeline)
		if err != nil {
			logger.Error("[CreatePipeline] Unable to create pipeline", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrPipelineCreateFail, http.StatusInternalServerError))
			return
		}
		insertedId, _ := res.InsertedID.(primitive.ObjectID)
		pipeline.Id = insertedId
		logger.Info("[CreatePipeline] Successfully created pipeline", map[string]interface{}{"pipeline": pipeline})
		encode(w, r, http.StatusCreated, pipeline)
	})
}

func handleGetPipeline(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		pipelineId := vars["pipelineId"]
		pipeline, err := database.NewPipeline(client).GetById(pipelineId)
		if err != nil {
			logger.Error("[GetPipeline] Unable to get pipeline", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		logger.Info("[GetPipeline] Successfully retrieved pipeline", map[string]interface{}{"pipeline": pipeline})
		encode(w, r, http.StatusOK, pipeline)
	})
}

func handleGetAllPipelines(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pipelines, err := database.NewPipeline(client).GetAll()
		if err != nil {
			logger.Error("[GetAllPipelines] Unable to get pipelines", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		logger.Info("[GetAllPipelines] Successfully retrieved pipelines", map[string]interface{}{"count": len(pipelines)})
		encode(w, r, http.StatusOK, pipelines)
	})
}

// NOTE: handler and data functions used in here are subjected to change depending on if frontend is able to validate that user has been previously registered in auth0
func handleUserLogin(client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		um, err := decode[models.UserModel](r)
		if err != nil {
			logger.Error("[UserLogin] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		_, err = database.NewUser(client).GetUserById(um.UserId)
		if errors.Is(err, sql.ErrNoRows) {
			user, err := database.NewUser(client).Create(&um)
			if err != nil {
				logger.Error("[UserLogin] Unable to create user", map[string]interface{}{"err": err})
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUserCreateFail, http.StatusInternalServerError))
				return
			}
			logger.Info("[UserLogin] Successfully created user", map[string]interface{}{"user": user})
			encode[any](w, r, http.StatusCreated, nil)
			return
		} else if err != nil {
			logger.Error("[UserLogin] Error querying user table using user_id", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		logger.Info("[UserLogin] Successfully logged in user", nil)
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleCreateOrganisation(client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		om, err := decode[models.OrganisationModel](r)
		if err != nil {
			logger.Error("[CreateOrganisation] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		org, err := database.NewOrganization(client).Create(&om)
		if err != nil {
			logger.Error("[CreateOrganisation] Unable to create organisation", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrOrganisationCreateFail, http.StatusInternalServerError))
			return
		}
		logger.Info("[CreateOrganisation] Successfully created organisation", map[string]interface{}{"org": org})
		encode(w, r, http.StatusCreated, org)
	})
}

func handleCreateMembership(client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mm, err := decode[models.MembershipModel](r)
		if err != nil {
			logger.Error("[CreateMembership] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		membership, err := database.NewMembership(client).Create(&mm)
		if err != nil {
			logger.Error("[CreateMembership] Unable to create membership", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrMembershipCreateFail, http.StatusInternalServerError))
			return
		}
		logger.Info("[CreateMembership] Successfully created membership", map[string]interface{}{"membership": membership})
		encode(w, r, http.StatusCreated, membership)
	})
}

func handleGetAllServiceRequestsForOrganisation(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orgId, err := strconv.Atoi(r.URL.Query().Get("org_id"))
		if err != nil {
			logger.Error("[GetAllServiceRequestsForOrganisation] Error converting organisation id to int", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrInvalidOrganisationId, http.StatusBadRequest))
			return
		}
		allsr, err := database.NewServiceRequest(client).GetAllServiceRequestsForOrgId(orgId)
		if err != nil {
			logger.Error("[GetAllServiceRequestsForOrganisation] Error retrieving all service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		logger.Info("[GetAllServiceRequestsForOrganisation] Successfully retrieved service requests", map[string]interface{}{"count": len(allsr)})
		encode(w, r, http.StatusOK, allsr)
	})
}

func handleGetUserById(client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user_id := vars["userId"]
		user, err := database.NewUser(client).GetUserById(user_id)
		if err != nil {
			logger.Error("[GetUserById] Unable to retrieve user", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUserRetrieve, http.StatusInternalServerError))
			return
		}
		logger.Info("[GetUserById] Successfully retrieved user exists", map[string]interface{}{"user": user})

		orgs, err := database.NewOrganization(client).GetAllOrgsByUserId(user.UserId)
		if err != nil {
			logger.Error("[GetUserById] Unable to retrieve user orgs", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrOrganisationRetrieve, http.StatusInternalServerError))
			return
		}
		logger.Info("[GetUserById] Successfully retrieved user organisations", map[string]interface{}{"orgs": orgs})
		user.Organisations = orgs
		encode(w, r, http.StatusCreated, user)
	})
}

func handleUpdateMembership(client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mm, err := decode[models.MembershipModel](r)
		if err != nil {
			logger.Error("[UpdateMembership] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		result, err := database.NewMembership(client).UpdateUserMembership(&mm)
		if err != nil {
			logger.Error("[UpdateMembership] Unable to update membership", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrMembershipUpdateFail, http.StatusInternalServerError))
			return
		}

		// NOTE: may not work for all db / db drivers
		rows, err := result.RowsAffected()
		if err != nil {
			logger.Error("[UpdateMembership] Unable to retrieve rows affected", nil)
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		} else if rows < 1 {
			logger.Error("[UpdateMembership] Membership does not exist", nil)
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrMembershipInvalid, http.StatusBadRequest))
			return
		}

		logger.Info("[UpdateMembership] Successfully updated membership", nil)
		encode[any](w, r, http.StatusOK, nil)
	})
}

func handleDeleteMembership(client *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mm, err := decode[models.MembershipModel](r)
		if err != nil {
			logger.Error("[DeleteMembership] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		_, err = database.NewMembership(client).UpdateUserMembership(&mm)
		if err != nil {
			logger.Error("[DeleteMembership] Unable to delete membership", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrMembershipDeleteFail, http.StatusInternalServerError))
			return
		}
		logger.Info("[DeleteMembership] Successfully delete membership", nil)
		encode[any](w, r, http.StatusOK, nil)
	})
}
