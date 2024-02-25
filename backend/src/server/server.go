package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/database/client"
)

func addRoutes(r *mux.Router) {
	mongoClient, err := client.GetMongoClient()
	if err != nil {
		panic(err)
	}

	// Health Check
	r.Handle("/api/healthcheck", handleHealthCheck()).Methods("GET")

	// Service Request
	r.Handle("/api/service_request", handleGetAllServiceRequest(mongoClient)).Methods("GET")
	r.Handle("/api/service_request/{requestId}", handleGetServiceRequest(mongoClient)).Methods("GET")
	r.Handle("/api/service_request", handleCreateServiceRequest(mongoClient)).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}", handleUpdateServiceRequest(mongoClient)).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}/cancel", handleCancelStartedServiceRequest(mongoClient)).Methods("GET")
	r.Handle("/api/service_request/{requestId}/start", handleStartServiceRequest(mongoClient)).Methods("GET")
	r.Handle("/api/service_request/{requestId}/approve", handleApproveServiceRequest(mongoClient)).Methods("POST").Headers("Content-Type", "application/json")

	// Pipeline
	r.Handle("/api/pipeline", handleGetAllPipelines(mongoClient)).Methods("GET")
	r.Handle("/api/pipeline/{pipelineId}", handleGetPipeline(mongoClient)).Methods("GET")
	r.Handle("/api/pipeline", validateCreatePipelineRequest(handleCreatePipeline(mongoClient))).Methods("POST").Headers("Content-Type", "application/json")
}

func New() http.Handler {
	router := mux.NewRouter()
	addRoutes(router)
	handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedHeaders([]string{
			"Content-Type",
			"Authorization",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Credentials",
		}),
	)(router)
	return router
}
