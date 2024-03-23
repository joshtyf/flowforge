package main

import (
	"context"
	"net/http"
	"os"

	"github.com/joshtyf/flowforge/src/database/client"
	"github.com/joshtyf/flowforge/src/execute"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/server"
)

func run() {
	psqlClient, err := client.GetPsqlClient()
	if err != nil {
		panic(err)
	}
	defer psqlClient.Close()
	mongoClient, err := client.GetMongoClient()
	if err != nil {
		panic(err)
	}
	defer mongoClient.Disconnect(context.Background())

	srm, err := execute.NewStepExecutionManager(
		mongoClient,
		psqlClient,
		execute.WithStepExecutor(execute.NewApiStepExecutor()),
		execute.WithStepExecutor(execute.NewWaitForApprovalStepExecutor()),
	)
	if err != nil {
		panic(err)
	}
	srm.Start()
	config := &server.ServerConfig{
		PsqlClient:   psqlClient,
		MongoClient:  mongoClient,
		ServerLogger: logger.NewLogger(os.Stdout),
	}
	http.ListenAndServe(":8080", server.New(config))
}

func main() {
	run()
}
