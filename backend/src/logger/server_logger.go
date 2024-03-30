package logger

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"runtime"
	"strings"
)

const (
	logWithFuncFormat = "[%s] %s: %s" // [MSG_TYPE] CALLER_DETAILS: MSG
	logWithoutFunc    = "[%s] %s"     // [MSG_TYPE] MSG
)

type ServerLogger interface {
	Info(msg string)
	Error(msg string)
	Warn(msg string)
}

type ServerLog struct {
	logger *log.Logger
}

func NewServerLog(f io.Writer) ServerLogger {
	return &ServerLog{
		logger: log.New(f, "", log.LstdFlags),
	}
}

// Returns information about the original caller of the function that called the public logging method.
// This method should not be called outside of the main info/error/warn methods as it is dependent on the callstack depth.
//
// Format: FILE_PATH:FUNC_NAME:LINE_NUMBER
//
// If the original caller is an anonymous function, the FUNC_NAME will be replaced with _anon_fn_.
func (l *ServerLog) getOriginalCaller() (string, error) {
	pc := make([]uintptr, 1)
	n := runtime.Callers(4, pc) // Callstack: getOriginalCaller -> info/error/warn -> public logging method -> caller
	if n == 0 {
		return "", fmt.Errorf("could not get caller. callstack depth is too shallow")
	}
	frames := runtime.CallersFrames(pc)
	frame, _ := frames.Next()
	frameFunc := frame.Func
	if frameFunc == nil {
		return "", fmt.Errorf("could not get caller frame")
	}
	callerFullPath := frameFunc.Name()
	callerFuncs := strings.Split(callerFullPath, ".")
	originalCaller := callerFuncs[len(callerFuncs)-1]
	// Check if original caller is anonymous function
	re := regexp.MustCompile(`\bfunc\d+\b`)
	if re.MatchString(originalCaller) {
		return fmt.Sprintf("%s:_anon_fn_:%d", frame.File, frame.Line), nil
	}
	return fmt.Sprintf("%s:%s:%d", frame.File, originalCaller, frame.Line), nil
}

// Helper method to tag log message with "[INFO]"
//
// This function should be called immediately after a public logging method so as to maintain the correct callstack depth.
func (l *ServerLog) Info(msg string) {
	if originalCaller, err := l.getOriginalCaller(); err == nil {
		l.logger.Printf(logWithFuncFormat, "INFO", originalCaller, msg)
		return
	}
	l.logger.Printf(logWithoutFunc, "INFO", msg)
}

// Helper method to tag log message with "[ERROR]"
//
// This function should be called immediately after a public logging method so as to maintain the correct callstack depth.
func (l *ServerLog) Error(msg string) {
	if originalCaller, err := l.getOriginalCaller(); err == nil {
		l.logger.Printf(logWithFuncFormat, "ERROR", originalCaller, msg)
		return
	}
	l.logger.Printf(logWithoutFunc, "ERROR", msg)
}

// Helper method to tag log message with "[WARN]"
//
// This function should be called immediately after a public logging method so as to maintain the correct callstack depth.
func (l *ServerLog) Warn(msg string) {
	if originalCaller, err := l.getOriginalCaller(); err == nil {
		l.logger.Printf(logWithFuncFormat, "WARN", originalCaller, msg)
		return
	}
	l.logger.Printf(logWithoutFunc, "WARN", msg)
}
