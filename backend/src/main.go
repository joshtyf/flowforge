package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/healthcheck", handlers.HealthCheck).Methods("GET")
	http.ListenAndServe(":8080", r)
}
