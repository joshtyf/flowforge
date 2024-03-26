package logger

import (
	"fmt"
	"io"
	"log"
	"strings"
)

const (
	logFormat = "[%s] %s {%s}"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

func Info(msg string, args map[string]interface{}) {
	argsSlice := make([]string, 0, len(args))
	for k, v := range args {
		argsSlice = append(argsSlice, fmt.Sprintf("%s=%v", k, v))
	}
	log.Printf(logFormat, "INFO", msg, strings.Join(argsSlice, ", "))
}

func Error(msg string, args map[string]interface{}) {
	argsSlice := make([]string, 0, len(args))
	for k, v := range args {
		argsSlice = append(argsSlice, fmt.Sprintf("%s=%v", k, v))
	}
	log.Printf(logFormat, "ERR", msg, strings.Join(argsSlice, ", "))
}

func Warn(msg string, args map[string]interface{}) {
	argsSlice := make([]string, 0, len(args))
	for k, v := range args {
		argsSlice = append(argsSlice, fmt.Sprintf("%s=%v", k, v))
	}
	log.Printf(logFormat, "WARN", msg, strings.Join(argsSlice, ", "))
}

type Logger struct {
	logger *log.Logger
}

func NewLogger(f io.Writer) *Logger {
	return &Logger{
		logger: log.New(f, "", log.LstdFlags),
	}
}
