package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/database/client"
	"github.com/joshtyf/flowforge/src/execute"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/server"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	SERVER_SHUTDOWN_GRACE_PERIOD = 10 * time.Second
)

func gracefulShutdown(logger logger.ServerLogger, svr *http.Server, psqlClient *sql.DB, mongoClient *mongo.Client) func(string) {
	shutdownHandler := func(reason string) {
		logger.Info(fmt.Sprintf("shutting down server: %s", reason))
		ctx, cancel := context.WithTimeout(context.Background(), SERVER_SHUTDOWN_GRACE_PERIOD)
		defer cancel()
		if err := svr.Shutdown(ctx); err != nil {
			log.Println("Error Gracefully Shutting Down API:", err)
		}

		if err := psqlClient.Close(); err != nil {
			log.Println("Error Gracefully Shutting Down PSQL Client:", err)
		}

		ctx, cancel = context.WithTimeout(context.Background(), SERVER_SHUTDOWN_GRACE_PERIOD)
		defer cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Println("Error Gracefully Shutting Down Mongo:", err)
		}
	}
	return shutdownHandler
}

func main() {
	logger := logger.NewServerLog(os.Stdout)
	psqlClient, err := client.GetPsqlClient()
	if err != nil {
		panic(err)
	}

	mongoClient, err := client.GetMongoClient()
	if err != nil {
		panic(err)
	}
	// Start the Step Execution Manager
	srm, err := execute.NewStepExecutionManager(
		mongoClient,
		psqlClient,
		logger,
		execute.WithStepExecutor(execute.NewApiStepExecutor()),
		execute.WithStepExecutor(execute.NewWaitForApprovalStepExecutor(mongoClient)),
	)
	if err != nil {
		panic(err)
	}
	srm.Start()

	// Create the server
	config := &server.ServerConfig{
		Address:      ":8080",
		Router:       mux.NewRouter(),
		PsqlClient:   psqlClient,
		MongoClient:  mongoClient,
		ServerLogger: logger,
	}
	svr := server.New(config)

	// Start the server
	srvErrs := make(chan error, 1)
	go func() {
		srvErrs <- svr.ListenAndServe()
	}()

	// Create a channel to listen for shutdown signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received or the server stops
	select {
	case err := <-srvErrs:
		gracefulShutdown(logger, &svr, psqlClient, mongoClient)(err.Error())
	case <-done:
		gracefulShutdown(logger, &svr, psqlClient, mongoClient)("received shutdown signal")
	}
}
