package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gookit/event"
	"github.com/gorilla/mux"
	handlermodels "github.com/joshtyf/flowforge/src/api/handlers/models"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	dbmodels "github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/events"
	"github.com/joshtyf/flowforge/src/logger"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/////////////////// Handlers ///////////////////

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if serverHealthy() {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server working!"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server not working!"))
	}
}

func GetAllServiceRequest(w http.ResponseWriter, r *http.Request) {
	client, err := client.GetMongoClient()
	if err != nil {
		JSONError(w, handlermodels.NewHttpError(err), http.StatusInternalServerError)
		return
	}
	allsr, err := database.NewServiceRequest(client).GetAll()
	if err != nil {
		JSONError(w, handlermodels.NewHttpError(err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allsr)
	w.WriteHeader(http.StatusOK)
}

func GetServiceRequest(w http.ResponseWriter, r *http.Request) {
	client, err := client.GetMongoClient()
	if err != nil {
		JSONError(w, handlermodels.NewHttpError(err), http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	requestId := vars["requestId"]
	sr, err := database.NewServiceRequest(client).GetById(requestId)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		JSONError(w, handlermodels.NewHttpError(err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(sr)
	w.WriteHeader(http.StatusOK)
	return
}

func CreateServiceRequest(w http.ResponseWriter, r *http.Request) {
	client, err := client.GetMongoClient()
	if err != nil {
		JSONError(w, handlermodels.NewHttpError(err), http.StatusInternalServerError)
		return
	}
	srm := &dbmodels.ServiceRequestModel{
		CreatedOn:   time.Now(),
		LastUpdated: time.Now(),
		Status:      dbmodels.Pending,
	}
	err = json.NewDecoder(r.Body).Decode(srm)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		JSONError(w, handlermodels.NewHttpError(err), http.StatusBadRequest)
		return
	}
	_, err = database.NewServiceRequest(client).Create(srm)
	if err != nil {
		JSONError(w, handlermodels.NewHttpError(err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func StartServiceRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["requestId"]
	client, err := client.GetMongoClient()
	if err != nil {
		logger.Error("[StartServiceRequest] Unable to get mongo client", map[string]interface{}{"err": err})
		JSONError(w, handlermodels.NewHttpError(errors.New("internal server error")), http.StatusInternalServerError)
		return
	}
	srm, err := database.NewServiceRequest(client).GetById(requestId)
	if err != nil {
		logger.Error("[StartServiceRequest] Unable retrieve service request", map[string]interface{}{"err": err})
		JSONError(w, handlermodels.NewHttpError(errors.New("internal server error")), http.StatusInternalServerError)
		return
	}
	event.FireAsync(events.NewNewServiceRequestEvent(srm))
	w.WriteHeader(http.StatusOK)
}

func CreatePipeline(w http.ResponseWriter, r *http.Request) {
	pipeline := &dbmodels.PipelineModel{
		CreatedOn: time.Now(),
		Version:   1,
	}
	err := json.NewDecoder(r.Body).Decode(pipeline)
	if err != nil {
		logger.Error("[CreatePipeline] Unable to parse json request body", map[string]interface{}{"err": err})
		JSONError(w, handlermodels.NewHttpError(errors.New("unable to parse json request body")), http.StatusBadRequest)
	}
	client, err := client.GetMongoClient()
	if err != nil {
		logger.Error("[CreatePipeline] Unable to get mongo client", map[string]interface{}{"err": err})
		JSONError(w, handlermodels.NewHttpError(errors.New("internal server error")), http.StatusInternalServerError)
		return
	}
	res, err := database.NewPipeline(client).Create(pipeline)
	if err != nil {
		logger.Error("[CreatePipeline] Unable to create pipeline", map[string]interface{}{"err": err})
		JSONError(w, handlermodels.NewHttpError(errors.New("failed to create pipeline")), http.StatusInternalServerError)
		return
	}
	insertedId, _ := res.InsertedID.(primitive.ObjectID)
	pipeline.Id = insertedId
	logger.Info("[CreatePipeline] Successfully created pipeline", map[string]interface{}{"pipeline": pipeline})
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pipeline)
}

func GetPipeline(w http.ResponseWriter, r *http.Request) {
	client, err := client.GetMongoClient()
	if err != nil {
		logger.Error("[GetPipeline] Unable to get mongo client", map[string]interface{}{"err": err})
		JSONError(w, handlermodels.NewHttpError(errors.New("internal server error")), http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	pipelineId := vars["pipelineId"]
	pipeline, err := database.NewPipeline(client).GetById(pipelineId)
	if err != nil {
		logger.Error("[GetPipeline] Unable to get pipeline", map[string]interface{}{"err": err})
		JSONError(w, handlermodels.NewHttpError(errors.New("internal server error")), http.StatusInternalServerError)
		return
	}
	logger.Info("[GetPipeline] Successfully retrieved pipeline", map[string]interface{}{"pipeline": pipeline})
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pipeline)
}

func GetAllPipelines(w http.ResponseWriter, r *http.Request) {
	client, err := client.GetMongoClient()
	if err != nil {
		logger.Error("[GetAllPipelines] Unable to get mongo client", map[string]interface{}{"err": err})
		JSONError(w, handlermodels.NewHttpError(errors.New("internal server error")), http.StatusInternalServerError)
		return
	}
	pipelines, err := database.NewPipeline(client).GetAll()
	if err != nil {
		logger.Error("[GetAllPipelines] Unable to get pipelines", map[string]interface{}{"err": err})
		JSONError(w, handlermodels.NewHttpError(errors.New("internal server error")), http.StatusInternalServerError)
		return
	}
	logger.Info("[GetAllPipelines] Successfully retrieved pipelines", map[string]interface{}{"count": len(pipelines)})
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pipelines)
}

/////////////////// Helper Functions ///////////////////

func serverHealthy() bool {
	// TODO: Include database health check

	return true
}

func JSONError(w http.ResponseWriter, err *handlermodels.HttpError, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
