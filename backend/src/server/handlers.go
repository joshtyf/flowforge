package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gookit/event"
	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
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

func NewHandlerError(err error, code int) *HandlerError {
	return &HandlerError{Message: err.Error(), Code: code}
}

type customHandlerFunc func(http.ResponseWriter, *http.Request) *HandlerError

func NewHandler(handlerFunc customHandlerFunc, mws ...customMiddleWareFunc) http.HandlerFunc {
	handlerChain := handlerFunc
	// Build the handler chain
	for _, mw := range mws {
		handlerChain = mw(handlerChain)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		err := handlerChain(w, r)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(err.Code)
			json.NewEncoder(w).Encode(err)
		}
	}
}

/////////////////// Handlers ///////////////////

func handleHealthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if serverHealthy() {
			encode(w, r, http.StatusOK, []byte("Server working!"))
			return
		}
		encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
	})
}

func handleGetAllServiceRequest(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allsr, err := database.NewServiceRequest(client).GetAll()
		if err != nil {
			logger.Error("[GetAllServiceRequest] Error retrieving all service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
		}
		logger.Info("[GetAllServiceRequest] Successfully retrieved service requests", map[string]interface{}{"count": len(allsr)})
		encode(w, r, http.StatusOK, allsr)
	})
}

func handleGetServiceRequest(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		sr, err := database.NewServiceRequest(client).GetById(requestId)
		if err != nil {
			logger.Error("[GetServiceRequest] Unable to retrieve service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
		}
		logger.Info("[GetServiceRequest] Successfully retrieved service request", map[string]interface{}{"sr": sr})
		encode(w, r, http.StatusOK, sr)
	})
}

func handleCreateServiceRequest(client *mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srm, err := decode[models.ServiceRequestModel](r)
		if err != nil {
			logger.Error("[CreateServiceRequest] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, NewHandlerError(ErrJsonParseError, http.StatusBadRequest))
		}
		srm.CreatedOn = time.Now()
		srm.LastUpdated = time.Now()
		srm.Status = models.Pending

		res, err := database.NewServiceRequest(client).Create(&srm)
		if err != nil {
			logger.Error("[CreateServiceRequest] Error creating service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
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
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
		}
		status := sr.Status
		if status != models.Pending && status != models.Running {
			logger.Error("[CancelStartedServiceRequest] Unable to cancel service request as execution has been completed", nil)
			encode(w, r, http.StatusBadRequest, NewHandlerError(ErrServiceRequestAlreadyCompleted, http.StatusBadRequest))
		}

		if status == models.NotStarted {
			logger.Error("[CancelStartedServiceRequest] Unable to cancel service request as execution has not been started", nil)
			encode(w, r, http.StatusBadRequest, NewHandlerError(ErrServiceRequestNotStarted, http.StatusBadRequest))
		}

		// TODO: implement cancellation of sr

		err = database.NewServiceRequest(client).UpdateStatus(requestId, models.Canceled)

		// TODO: discuss how to handle this error
		if err != nil {
			logger.Error("[CancelStartedServiceRequest] Unable to update service request status", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
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
			encode(w, r, http.StatusBadRequest, NewHandlerError(ErrJsonParseError, http.StatusBadRequest))
		}
		vars := mux.Vars(r)
		requestId := vars["requestId"]
		sr, err := database.NewServiceRequest(client).GetById(requestId)
		if err != nil {
			logger.Error("[UpdateServiceRequest] Error getting service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
		}
		status := sr.Status
		if status != models.NotStarted {
			logger.Error("[UpdateServiceRequest] Unable to update service request as service request has been started", nil)
			encode(w, r, http.StatusBadRequest, NewHandlerError(ErrServiceRequestAlreadyStarted, http.StatusBadRequest))
		}
		res, err := database.NewServiceRequest(client).UpdateById(requestId, &srm)
		if err != nil {
			logger.Error("[UpdateServiceRequest] Error updating service request", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
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
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
		}
		if srm.Status != models.NotStarted {
			logger.Error("[StartServiceRequest] Service request has already been started", nil)
			encode(w, r, http.StatusBadRequest, NewHandlerError(ErrServiceRequestAlreadyStarted, http.StatusBadRequest))
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
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
		}
		body, err := decode[requestBody](r)
		if err != nil {
			logger.Error("[ApproveServiceRequest] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, NewHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}
		pipeline, err := database.NewPipeline(client).GetById(serviceRequest.PipelineId)
		if err != nil {
			logger.Error("[ApproveServiceRequest] Unable to get pipeline", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		waitForApprovalStep := pipeline.GetPipelineStep(body.StepName)

		if waitForApprovalStep == nil {
			logger.Error("[ApproveServiceRequest] Unable to get wait for approval step", map[string]interface{}{"step": body.StepName})
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		if waitForApprovalStep.StepType != models.WaitForApprovalStep {
			logger.Error("[ApproveServiceRequest] Step is not a wait for approval step", map[string]interface{}{"step": body.StepName})
			encode(w, r, http.StatusBadRequest, NewHandlerError(ErrWrongStepType, http.StatusBadRequest))
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
			encode(w, r, http.StatusBadRequest, NewHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}

		pipeline.CreatedOn = time.Now()
		pipeline.Version = 1
		res, err := database.NewPipeline(client).Create(&pipeline)
		if err != nil {
			logger.Error("[CreatePipeline] Unable to create pipeline", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrPipelineCreateFail, http.StatusInternalServerError))
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
			encode(w, r, http.StatusInternalServerError, NewHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}
		logger.Info("[GetPipeline] Successfully retrieved pipeline", map[string]interface{}{"pipeline": pipeline})
		encode(w, r, http.StatusOK, pipeline)
	})
}

func GetAllPipelines(w http.ResponseWriter, r *http.Request) *HandlerError {
	client, err := client.GetMongoClient()
	if err != nil {
		logger.Error("[GetAllPipelines] Unable to get mongo client", map[string]interface{}{"err": err})
		return NewHandlerError(ErrInternalServerError, http.StatusInternalServerError)
	}
	pipelines, err := database.NewPipeline(client).GetAll()
	if err != nil {
		logger.Error("[GetAllPipelines] Unable to get pipelines", map[string]interface{}{"err": err})
		return NewHandlerError(ErrInternalServerError, http.StatusInternalServerError)
	}
	logger.Info("[GetAllPipelines] Successfully retrieved pipelines", map[string]interface{}{"count": len(pipelines)})
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pipelines)
	return nil
}

/////////////////// Helper Functions ///////////////////

func serverHealthy() bool {
	// TODO: Include database health check

	return true
}
