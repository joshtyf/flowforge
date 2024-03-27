package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	BaseLogDir = "./executor_logs"

	errGettingServiceReqFromCtxMsg = "error getting service request from context"
	errGettingStepFromCtxMsg       = "error getting step from context"
	httpRequestMsg                 = "%s %s" // method, url TODO: add request body, params
	httpResponseStatusMsg          = "status code: %d"
	waitingForApprovalMsg          = "waiting for approval"
)

func FindExecutorLogFile(serviceRequestId, stepName string) (*os.File, error) {
	return os.Open(fmt.Sprintf("%s/%s/%s.log", BaseLogDir, serviceRequestId, stepName))
}

type ExecutorLogger struct {
	logger *log.Logger
}

func NewExecutorLogger(w io.Writer, stepName string) *ExecutorLogger {
	logger := log.New(w, fmt.Sprintf("[%s] ", stepName), log.LstdFlags)
	return &ExecutorLogger{
		logger: logger,
	}
}

func (l *ExecutorLogger) ErrGettingStepFromCtx() {
	l.logger.Println(errGettingStepFromCtxMsg)
}

func (l *ExecutorLogger) ErrGettingServiceReqFromCtx() {
	l.logger.Println(errGettingServiceReqFromCtxMsg)
}

// Logs the method and url of the API call
func (l *ExecutorLogger) HttpRequest(method, url string) {
	l.logger.Printf(httpRequestMsg, method, url)
}

// Logs the status code of the HTTP response
func (l *ExecutorLogger) HttpResponseStatus(statusCode int) {
	l.logger.Printf(httpResponseStatusMsg, statusCode)
}

// Logs that the step is waiting for approval
func (l *ExecutorLogger) WaitingForApproval() {
	l.logger.Println(waitingForApprovalMsg)
}
