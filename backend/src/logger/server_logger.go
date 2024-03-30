package logger

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"runtime"
	"strings"

	"github.com/joshtyf/flowforge/src/database/models"
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

/* PUBLIC LOGGING METHODS */

func (l *ServerLog) ErrClaimsMissingPermission(permission string) {
	l.Error(fmt.Sprintf("unauthorized: missing permission %s", permission))
}

func (l *ServerLog) ErrEventLogStepMissingInPipeline(stepName string) {
	l.Error(fmt.Sprintf("%s exists in event log but not in pipeline template", stepName))
}

func (l *ServerLog) ErrExecutingStep(stepName string, err error) {
	l.Error(fmt.Sprintf("error executing %s: %s", stepName, err))
}

// Generic error message for handling events
func (l *ServerLog) ErrHandlingEvent(err error) {
	l.Error(fmt.Sprintf("error encountered while handling event: %s", err))
}

// Generic error message for handling API requests
func (l *ServerLog) ErrHandlingRequest(err error) {
	l.Error(fmt.Sprintf("error encountered while handling API request: %s", err))
}

func (l *ServerLog) ErrInvalidStepType(expectedStepType, actualStepType models.PipelineStepType) {
	l.Error(fmt.Sprintf("invalid step type: expected %s got %s", expectedStepType, actualStepType))
}

func (l *ServerLog) ErrMissingExecutorForStep(stepName string) {
	l.Error(fmt.Sprintf("missing executor for step: %s", stepName))
}

func (l *ServerLog) ErrMissingPipelineStep(stepName string) {
	l.Error(fmt.Sprintf("missing pipeline step: %s", stepName))
}

func (l *ServerLog) ErrParsingIssuerURL(err error) {
	l.Error(fmt.Sprintf("failed to parse the issuer url: %s", err))
}

func (l *ServerLog) ErrParsingJsonRequestBody(err error) {
	l.Error(fmt.Sprintf("failed to parse json request body: %s", err))
}

func (l *ServerLog) ErrResourceNotFound(resourceName, resourceId string) {
	l.Error(fmt.Sprintf("%s %s not found", resourceName, resourceId))
}

func (l *ServerLog) ErrServiceRequestStatusUpdateFailed(action, serviceRequestId, reason string) {
	l.Error(fmt.Sprintf("failed to %s service request %s: %s", action, serviceRequestId, reason))
}

func (l *ServerLog) ErrSettingUpJWTValidator(err error) {
	l.Error(fmt.Sprintf("failed to set up jwt validator: %s", err))
}

func (l *ServerLog) ErrValidatingJWT(err error) {
	l.Error(fmt.Sprintf("failed to validate jwt: %s", err))
}

func (l *ServerLog) ErrValidatingPipeline(err error) {
	l.Error(fmt.Sprintf("failed to validate pipeline: %s", err))
}

func (l *ServerLog) ErrEventMissingData(eventName, data string) {
	l.Error(fmt.Sprintf("event %s missing data: %s", eventName, data))
}

func (l *ServerLog) HandleEvent(eventName string) {
	l.Info(fmt.Sprintf("handling event: %s", eventName))
}

func (l *ServerLog) ResourceCreated(resourceName, resourceId string) {
	l.Info(fmt.Sprintf("%s %s created", resourceName, resourceId))
}

func (l *ServerLog) ResourceDeleted(resourceName, resourceId string) {
	l.Info(fmt.Sprintf("%s %s deleted", resourceName, resourceId))
}

func (l *ServerLog) ServerHealthy() {
	l.Info("server is healthy")
}

func (l *ServerLog) ShutdownServer(reason string) {
	l.Info(fmt.Sprintf("shutting down server: %s", reason))
}

func (l *ServerLog) SkippingAdminCheck() {
	l.Warn("skipping admin check in dev environment")
}

func (l *ServerLog) SkippingAuth() {
	l.Warn("skipping authentication in dev environment")
}

func (l *ServerLog) UserLoggedIn(userId string) {
	l.Info(fmt.Sprintf("user %s logged in", userId))
}
