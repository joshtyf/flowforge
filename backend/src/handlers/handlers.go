package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	dbmodels "github.com/joshtyf/flowforge/src/database/models"
	handlermodels "github.com/joshtyf/flowforge/src/handlers/models"
	_ "github.com/lib/pq"
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
	var srm *dbmodels.ServiceRequestModel
	err = json.NewDecoder(r.Body).Decode(&srm)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		JSONError(w, handlermodels.NewHttpError(err), http.StatusBadRequest)
		return
	}
	err = database.NewServiceRequest(client).Create(srm)
	if err != nil {
		JSONError(w, handlermodels.NewHttpError(err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
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
