package logger

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"
)

const (
	logFormat         = "[%s] %s {%s}"
	logWithFuncFormat = "[%s] %s: %s" // [MSG_TYPE] FUNC_NAME: MSG
	logWithoutFunc    = "[%s] %s"     // [MSG_TYPE] MSG
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

func (l *Logger) info(msg string) {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		l.logger.Printf(logWithoutFunc, "INFO", msg)
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		l.logger.Printf(logWithoutFunc, "INFO", msg)
	}
	funcName := f.Name()
	l.logger.Printf(logWithFuncFormat, "INFO", funcName, msg)
}

func (l *Logger) error(msg string) {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		l.logger.Printf(logWithoutFunc, "ERROR", msg)
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		l.logger.Printf(logWithoutFunc, "ERROR", msg)
	}
	funcName := f.Name()
	l.logger.Printf(logWithFuncFormat, "ERROR", funcName, msg)
}

func (l *Logger) warn(msg string) {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		l.logger.Printf(logWithoutFunc, "WARN", msg)
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		l.logger.Printf(logWithoutFunc, "WARN", msg)
	}
	funcName := f.Name()
	l.logger.Printf(logWithFuncFormat, "WARN", funcName, msg)
}
