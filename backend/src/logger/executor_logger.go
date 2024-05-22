package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	baseLogDir = "./executor_logs"
)

func CreateExecutorLogDir(serviceRequestId string) error {
	return os.MkdirAll(fmt.Sprintf("%s/%s", baseLogDir, serviceRequestId), 0755)
}

func CreateExecutorLogFilePath(serviceRequestId, stepName string) string {
	return fmt.Sprintf("%s/%s/%s.log", baseLogDir, serviceRequestId, stepName)
}

func FindExecutorLogFile(serviceRequestId, stepName string) (*os.File, error) {
	return os.Open(fmt.Sprintf("%s/%s/%s.log", baseLogDir, serviceRequestId, stepName))
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

func (l *ExecutorLogger) Info(msg string) {
	l.logger.Println(fmt.Printf("[INFO] %s", msg))
}

func (l *ExecutorLogger) Error(msg string) {
	l.logger.Println(fmt.Printf("[ERROR] %s", msg))
}

func (l *ExecutorLogger) Warn(msg string) {
	l.logger.Println(fmt.Printf("[WARN] %s", msg))
}
