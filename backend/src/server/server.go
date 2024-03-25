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

	psqlClient, err := client.GetPsqlClient()
	if err != nil {
		panic(err)
	}

	// Health Check
	r.Handle("/api/healthcheck", handleHealthCheck()).Methods("GET")

	// Service Request
	r.Handle("/api/service_request", isAuthenticated(handleGetAllServiceRequest(mongoClient))).Methods("GET")
	// TODO: @Zheng-Zhi-Qiang this route conflicts with `/api/service_request/{requestId}`.
	// r.Handle("/api/service_request/{organisationId}", isAuthenticated(handleGetServiceRequestsByOrganisation(mongoClient))).Methods("GET")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(handleGetServiceRequest(mongoClient, psqlClient))).Methods("GET")
	r.Handle("/api/service_request", isAuthenticated(handleCreateServiceRequest(mongoClient, psqlClient))).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(handleUpdateServiceRequest(mongoClient))).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}/cancel", isAuthenticated(handleCancelStartedServiceRequest(mongoClient))).Methods("GET")
	r.Handle("/api/service_request/{requestId}/start", isAuthenticated(handleStartServiceRequest(mongoClient))).Methods("GET")
	r.Handle("/api/service_request/{requestId}/approve", isAuthenticated(isAuthorisedAdmin(handleApproveServiceRequest(mongoClient)))).Methods("POST").Headers("Content-Type", "application/json")

	// Pipeline
	r.Handle("/api/pipeline", isAuthenticated(isAuthorisedAdmin(handleGetAllPipelines(mongoClient)))).Methods("GET")
	r.Handle("/api/pipeline/{pipelineId}", isAuthenticated(isAuthorisedAdmin(handleGetPipeline(mongoClient)))).Methods("GET")
	r.Handle("/api/pipeline", isAuthenticated(isAuthorisedAdmin(validateCreatePipelineRequest(handleCreatePipeline(mongoClient))))).Methods("POST").Headers("Content-Type", "application/json")

	// User
	r.Handle("/api/user/{userId}", isAuthenticated(handleGetUserById(psqlClient))).Methods("Get")
	r.Handle("/api/login", isAuthenticated(handleUserLogin(psqlClient))).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/organisation", isAuthenticated(handleCreateOrganisation(psqlClient))).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/membership", isAuthenticated(handleCreateMembership(psqlClient))).Methods("POST").Headers("Content-Type", "application/json")
}

func New() http.Handler {
	router := mux.NewRouter()
	addRoutes(router)
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{
			"Content-Type",
			"Authorization",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Credentials",
		}),
	)(router)
}
