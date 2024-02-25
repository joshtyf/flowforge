package main

import (
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/execute"
	"github.com/joshtyf/flowforge/src/server"
)

func main() {
	srm := execute.NewStepExecutionManager(
		execute.WithStepExecutor(execute.NewApiStepExecutor()),
		execute.WithStepExecutor(execute.NewWaitForApprovalStepExecutor()),
	)
	srm.Start()

	r := mux.NewRouter()
	r.HandleFunc("/api/healthcheck", server.NewHandler(server.HealthCheck)).Methods("GET")
	r.HandleFunc("/api/service_request", server.NewHandler(server.CreateServiceRequest)).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/service_request/{requestId}", server.NewHandler(server.GetServiceRequest)).Methods("GET")
	r.HandleFunc("/api/service_request", server.NewHandler(server.GetAllServiceRequest)).Methods("GET")
	r.HandleFunc("/api/service_request/{requestId}/approve", server.NewHandler(server.ApproveServiceRequest)).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/service_request/{requestId}/cancel", server.NewHandler(server.CancelStartedServiceRequest)).Methods("GET")
	r.HandleFunc("/api/service_request/{requestId}/start", server.NewHandler(server.StartServiceRequest)).Methods("GET")
	r.HandleFunc("/api/service_request/{requestId}", server.NewHandler(server.UpdateServiceRequest)).Methods("PATCH").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/pipeline", server.NewHandler(
		server.CreatePipeline,
		server.ValidateCreatePipelineRequest,
	)).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/pipeline", server.NewHandler(server.GetAllPipelines)).Methods("GET")
	r.HandleFunc("/api/pipeline/{pipelineId}", server.NewHandler(server.GetPipeline)).Methods("GET")
	http.ListenAndServe(":8080", gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000"}),
		gorillaHandlers.AllowedHeaders([]string{
			"Content-Type",
			"Authorization",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Credentials",
		}),
	)(r))
}
