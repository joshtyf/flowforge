package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gookit/event"
	"github.com/gorilla/mux"
	handlermodels "github.com/joshtyf/flowforge/src/api/handlers/models"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	dbmodels "github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/events"
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
	res, err := database.NewServiceRequest(client).Create(srm)
	if err != nil {
		JSONError(w, handlermodels.NewHttpError(err), http.StatusInternalServerError)
		return
	}
	insertedId, _ := res.InsertedID.(primitive.ObjectID)
	srm.Id = insertedId
	event.FireEvent(events.NewNewServiceRequestEvent(srm))
	w.WriteHeader(http.StatusCreated)
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
