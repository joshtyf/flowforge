package main

import (
	"net/http"

	"github.com/joshtyf/flowforge/src/execute"
	"github.com/joshtyf/flowforge/src/server"
)

func main() {
	srm := execute.NewStepExecutionManager(
		execute.WithStepExecutor(execute.NewApiStepExecutor()),
		execute.WithStepExecutor(execute.NewWaitForApprovalStepExecutor()),
	)
	srm.Start()

	http.ListenAndServe(":8080", server.New())
}
