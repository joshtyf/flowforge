package main

import (
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	handlers "github.com/joshtyf/flowforge/src/api/handlers"
	"github.com/joshtyf/flowforge/src/execute"
)

func main() {
	srm := execute.NewStepExecutionManager(
		execute.WithStepExecutor(execute.NewApiStepExecutor()),
		execute.WithStepExecutor(execute.NewWaitForApprovalStepExecutor()),
	)
	srm.Start()

	r := mux.NewRouter()
	r.HandleFunc("/api/healthcheck", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/api/service_request", handlers.CreateServiceRequest).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/service_request/{requestId}", handlers.GetServiceRequest).Methods("GET")
	r.HandleFunc("/api/service_request", handlers.GetAllServiceRequest).Methods("GET")
	r.HandleFunc("/api/service_request/{requestId}/start", handlers.StartServiceRequest).Methods("GET")
	r.HandleFunc("/api/service_request/{requestId}/approve", handlers.ApproveServiceRequest).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/service_request/{requestId}/cancel", handlers.CancelStartedServiceRequest).Methods("GET")
	r.HandleFunc("/api/service_request/{requestId}", handlers.UpdateServiceRequest).Methods("PATCH").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/pipeline", handlers.CreatePipeline).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/pipeline", handlers.GetAllPipelines).Methods("GET")
	r.HandleFunc("/api/pipeline/{pipelineId}", handlers.GetPipeline).Methods("GET")
	http.ListenAndServe(":8080", gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"localhost:3000"}))(r))
}
