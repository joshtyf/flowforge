package main

import (
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/api"
	"github.com/joshtyf/flowforge/src/execute"
)

func main() {
	srm := execute.NewStepExecutionManager(
		execute.WithStepExecutor(execute.NewApiStepExecutor()),
		execute.WithStepExecutor(execute.NewWaitForApprovalStepExecutor()),
	)
	srm.Start()

	r := mux.NewRouter()
	r.HandleFunc("/api/healthcheck", api.NewHandler(api.HealthCheck)).Methods("GET")
	r.HandleFunc("/api/service_request", api.NewHandler(api.CreateServiceRequest)).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/service_request/{requestId}", api.NewHandler(api.GetServiceRequest)).Methods("GET")
	r.HandleFunc("/api/service_request", api.NewHandler(api.GetAllServiceRequest)).Methods("GET")
	r.HandleFunc("/api/service_request/{requestId}/approve", api.NewHandler(api.ApproveServiceRequest)).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/service_request/{requestId}/cancel", api.NewHandler(api.CancelStartedServiceRequest)).Methods("GET")
	r.HandleFunc("/api/service_request/{requestId}/start", api.NewHandler(api.StartServiceRequest)).Methods("GET")
	r.HandleFunc("/api/service_request/{requestId}", api.NewHandler(api.UpdateServiceRequest)).Methods("PATCH").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/pipeline", api.NewHandler(api.CreatePipeline)).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/pipeline", api.NewHandler(api.GetAllPipelines)).Methods("GET")
	r.HandleFunc("/api/pipeline/{pipelineId}", api.NewHandler(api.GetPipeline)).Methods("GET")
	http.ListenAndServe(":8080", gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000"}),
		gorillaHandlers.AllowedHeaders([]string{
			"Content-Type",
			"Authorization",
			"Accept",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Credentials",
		}),
	)(r))
}
