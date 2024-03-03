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
	r.Handle("/api/healthcheck", isAuthenticated(isAuthorisedAdmin(handleHealthCheck()))).Methods("GET")

	// Service Request
	r.Handle("/api/service_request", isAuthenticated(isAuthorisedUser(handleGetAllServiceRequest(mongoClient)))).Methods("GET")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(isAuthorisedUser(handleGetServiceRequest(mongoClient)))).Methods("GET")
	r.Handle("/api/service_request", isAuthenticated(isAuthorisedUser(handleCreateServiceRequest(mongoClient)))).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(isAuthorisedUser(handleUpdateServiceRequest(mongoClient)))).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}/cancel", isAuthenticated(isAuthorisedUser(handleCancelStartedServiceRequest(mongoClient)))).Methods("GET")
	r.Handle("/api/service_request/{requestId}/start", isAuthenticated(isAuthorisedUser(handleStartServiceRequest(mongoClient)))).Methods("GET")
	r.Handle("/api/service_request/{requestId}/approve", isAuthenticated(isAuthorisedAdmin(handleApproveServiceRequest(mongoClient)))).Methods("POST").Headers("Content-Type", "application/json")

	// Pipeline
	r.Handle("/api/pipeline", isAuthenticated(isAuthorisedUser(handleGetAllPipelines(mongoClient)))).Methods("GET")
	r.Handle("/api/pipeline/{pipelineId}", isAuthenticated(isAuthorisedUser(handleGetPipeline(mongoClient)))).Methods("GET")
	r.Handle("/api/pipeline", isAuthenticated(isAuthorisedAdmin(validateCreatePipelineRequest(handleCreatePipeline(mongoClient))))).Methods("POST").Headers("Content-Type", "application/json")
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
