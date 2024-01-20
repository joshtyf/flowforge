package logger

import (
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
		argsSlice = append(argsSlice, k+"="+v.(string))
	}
	log.Printf(logFormat, "INFO", msg, strings.Join(argsSlice, ", "))
}

func Error(msg string, args map[string]interface{}) {
	argsSlice := make([]string, 0, len(args))
	for k, v := range args {
		argsSlice = append(argsSlice, k+"="+v.(string))
	}
	log.Printf(logFormat, "ERR", msg, strings.Join(argsSlice, ", "))
}

func Warn(msg string, args map[string]interface{}) {
	argsSlice := make([]string, 0, len(args))
	for k, v := range args {
		argsSlice = append(argsSlice, k+"="+v.(string))
	}
	log.Printf(logFormat, "WARN", msg, strings.Join(argsSlice, ", "))
}
