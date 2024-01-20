package main

import (
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/joshtyf/flowforge/src/handlers/api"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/healthcheck", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/api/servicerequest/new", handlers.CreateServiceRequest).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/api/servicerequest/{requestId}", handlers.GetServiceRequest).Methods("GET")
	r.HandleFunc("/api/servicerequest", handlers.GetAllServiceRequest).Methods("GET")
	http.ListenAndServe(":8080", r)
}
