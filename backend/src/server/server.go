package server

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServerConfig struct {
	PsqlClient   *sql.DB
	MongoClient  *mongo.Client
	ServerLogger *logger.Logger
}

func New(c *ServerConfig) http.Handler {
	router := mux.NewRouter()
	serverHandler := NewServerHandler(c.PsqlClient, c.MongoClient, c.ServerLogger)
	serverHandler.registerRoutes(router)
	return handlers.CORS(
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
}
