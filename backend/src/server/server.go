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
	r.Handle("/api/service_request", isAuthenticated(isOrgMember(mongoClient, psqlClient, handleGetAllServiceRequestsForOrganisation(mongoClient)))).Methods("GET")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(isOrgMember(mongoClient, psqlClient, handleGetServiceRequest(mongoClient, psqlClient)))).Methods("GET")
	r.Handle("/api/service_request", isAuthenticated(isOrgMember(mongoClient, psqlClient, handleCreateServiceRequest(mongoClient, psqlClient)))).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}", isAuthenticated(isOrgMember(mongoClient, psqlClient, handleUpdateServiceRequest(mongoClient)))).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/service_request/{requestId}/cancel", isAuthenticated(isOrgMember(mongoClient, psqlClient, handleCancelStartedServiceRequest(mongoClient)))).Methods("PUT")
	r.Handle("/api/service_request/{requestId}/start", isAuthenticated(isOrgMember(mongoClient, psqlClient, handleStartServiceRequest(mongoClient)))).Methods("PUT")
	r.Handle("/api/service_request/{requestId}/approve", isAuthenticated(isOrgAdmin(mongoClient, psqlClient, handleApproveServiceRequest(mongoClient)))).Methods("POST").Headers("Content-Type", "application/json")

	// Pipeline
	r.Handle("/api/pipeline", isAuthenticated(isOrgAdmin(mongoClient, psqlClient, handleGetAllPipelines(mongoClient)))).Methods("GET")
	r.Handle("/api/pipeline/{pipelineId}", isAuthenticated(isOrgAdmin(mongoClient, psqlClient, handleGetPipeline(mongoClient)))).Methods("GET")
	r.Handle("/api/pipeline", isAuthenticated(isOrgAdmin(mongoClient, psqlClient, validateCreatePipelineRequest(handleCreatePipeline(mongoClient))))).Methods("POST").Headers("Content-Type", "application/json")

	// User
	r.Handle("/api/user/{userId}", isAuthenticated(handleGetUserById(psqlClient))).Methods("Get")
	r.Handle("/api/login", isAuthenticated(handleUserLogin(psqlClient))).Methods("POST").Headers("Content-Type", "application/json")

	// Organisation
	r.Handle("/api/organisation", isAuthenticated(handleCreateOrganisation(psqlClient))).Methods("POST").Headers("Content-Type", "application/json")

	// Membership
	r.Handle("/api/membership", isAuthenticated(isOrgAdmin(mongoClient, psqlClient, handleCreateMembership(psqlClient)))).Methods("POST").Headers("Content-Type", "application/json")
	r.Handle("/api/membership", isAuthenticated(isOrgAdmin(mongoClient, psqlClient, handleUpdateMembership(psqlClient)))).Methods("PATCH").Headers("Content-Type", "application/json")
	r.Handle("/api/membership", isAuthenticated(isOrgOwner(mongoClient, psqlClient, handleDeleteMembership(psqlClient)))).Methods("DELETE").Headers("Content-Type", "application/json")
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
