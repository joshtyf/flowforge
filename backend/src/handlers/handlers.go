package handlers

import (
	"net/http"

	_ "github.com/lib/pq"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if serverHealthy() {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server working!"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server not working!"))
	}
}

func serverHealthy() bool {
	// TODO: Include database health check

	return true
}
