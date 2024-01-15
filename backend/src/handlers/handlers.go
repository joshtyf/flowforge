package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB_USER     = os.Getenv("POSTGRES_USER")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	DB_NAME     = os.Getenv("POSTGRES_DB")
	DB_HOST     = os.Getenv("POSTGRES_HOST")
	DB_PORT     = os.Getenv("POSTGRES_PORT")
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
	postgresDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", postgresDB)
	if err != nil {
		return false
	}

	err = db.Ping()

	if err != nil {
		return false
	}

	db.Close()
	return true
}
