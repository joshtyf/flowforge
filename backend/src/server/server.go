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
	Address      string
	Router       *mux.Router
	PsqlClient   *sql.DB
	MongoClient  *mongo.Client
	ServerLogger *logger.ServerLogger
}

func New(c *ServerConfig) http.Server {
	serverHandler := NewServerHandler(c.PsqlClient, c.MongoClient, c.ServerLogger)
	serverHandler.registerRoutes(c.Router)
	return http.Server{
		Addr: c.Address,
		Handler: handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedHeaders([]string{
				"Content-Type",
				"Authorization",
				"Access-Control-Allow-Origin",
				"Access-Control-Allow-Headers",
				"Access-Control-Allow-Methods",
				"Access-Control-Allow-Credentials",
			}),
		)(c.Router),
	}
}
