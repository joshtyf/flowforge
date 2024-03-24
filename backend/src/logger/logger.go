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
	logFormat         = "[%s] %s {%s}"
	logWithFuncFormat = "[%s] %s: %s" // [MSG_TYPE] FUNC_NAME: MSG
	logWithoutFunc    = "[%s] %s"     // [MSG_TYPE] MSG

	shutdownServerMsg = "shutting down server: %s" // shutting down server: REASON
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

func (l *Logger) getOriginalCaller() (string, error) {
	pc := make([]uintptr, 1)
	n := runtime.Callers(4, pc) // Callstack: getOriginalCaller -> info/error/warn -> message method -> caller
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

func (l *Logger) info(msg string) {
	if originalCaller, err := l.getOriginalCaller(); err == nil {
		l.logger.Printf(logWithFuncFormat, "INFO", originalCaller, msg)
		return
	}
	l.logger.Printf(logWithoutFunc, "INFO", msg)
}

func (l *Logger) error(msg string) {
	if originalCaller, err := l.getOriginalCaller(); err == nil {
		l.logger.Printf(logWithFuncFormat, "ERROR", originalCaller, msg)
		return
	}
	l.logger.Printf(logWithoutFunc, "ERROR", msg)
}

func (l *Logger) warn(msg string) {
	if originalCaller, err := l.getOriginalCaller(); err == nil {
		l.logger.Printf(logWithFuncFormat, "WARN", originalCaller, msg)
		return
	}
	l.logger.Printf(logWithoutFunc, "WARN", msg)
}

func (l *Logger) ShutdownServer(reason string) {
	l.info(fmt.Sprintf(shutdownServerMsg, reason))
}
