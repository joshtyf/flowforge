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

type ServerLogger struct {
	logger *log.Logger
}

func NewServerLogger(f io.Writer) *ServerLogger {
	return &ServerLogger{
		logger: log.New(f, "", log.LstdFlags),
	}
}

// Returns information about the original caller of the function that called the public logging method.
// This method should not be called outside of the main info/error/warn methods as it is dependent on the callstack depth.
//
// Format: FILE_PATH:FUNC_NAME:LINE_NUMBER
//
// If the original caller is an anonymous function, the FUNC_NAME will be replaced with _anon_fn_.
func (l *ServerLogger) getOriginalCaller() (string, error) {
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
func (l *ServerLogger) info(msg string) {
	if originalCaller, err := l.getOriginalCaller(); err == nil {
		l.logger.Printf(logWithFuncFormat, "INFO", originalCaller, msg)
		return
	}
	l.logger.Printf(logWithoutFunc, "INFO", msg)
}

// Helper method to tag log message with "[ERROR]"
//
// This function should be called immediately after a public logging method so as to maintain the correct callstack depth.
func (l *ServerLogger) error(msg string) {
	if originalCaller, err := l.getOriginalCaller(); err == nil {
		l.logger.Printf(logWithFuncFormat, "ERROR", originalCaller, msg)
		return
	}
	l.logger.Printf(logWithoutFunc, "ERROR", msg)
}

// Helper method to tag log message with "[WARN]"
//
// This function should be called immediately after a public logging method so as to maintain the correct callstack depth.
func (l *ServerLogger) warn(msg string) {
	if originalCaller, err := l.getOriginalCaller(); err == nil {
		l.logger.Printf(logWithFuncFormat, "WARN", originalCaller, msg)
		return
	}
	l.logger.Printf(logWithoutFunc, "WARN", msg)
}

/* PUBLIC LOGGING METHODS */

func (l *ServerLogger) ErrClaimsMissingPermission(permission string) {
	l.error(fmt.Sprintf("unauthorized: missing permission %s", permission))
}

func (l *ServerLogger) ErrEventLogStepMissingInPipeline(stepName string) {
	l.error(fmt.Sprintf("%s exists in event log but not in pipeline template", stepName))
}

func (l *ServerLogger) ErrExecutingStep(stepName string, err error) {
	l.error(fmt.Sprintf("error executing %s: %s", stepName, err))
}

func (l *ServerLogger) ErrHandlingEvent(err error) {
	l.error(fmt.Sprintf("error encountered while handling event: %s", err))
}

func (l *ServerLogger) ErrHandlingRequest(err error) {
	l.error(fmt.Sprintf("error encountered while handling API request: %s", err))
}

func (l *ServerLogger) ErrInvalidStepType(expectedStepType, actualStepType models.PipelineStepType) {
	l.error(fmt.Sprintf("invalid step type: expected %s got %s", expectedStepType, actualStepType))
}

func (l *ServerLogger) ErrMissingExecutorForStep(stepName string) {
	l.error(fmt.Sprintf("missing executor for step: %s", stepName))
}

func (l *ServerLogger) ErrMissingPipelineStep(stepName string) {
	l.error(fmt.Sprintf("missing pipeline step: %s", stepName))
}

func (l *ServerLogger) ErrParsingIssuerURL(err error) {
	l.error(fmt.Sprintf("failed to parse the issuer url: %s", err))
}

func (l *ServerLogger) ErrParsingJsonRequestBody(err error) {
	l.error(fmt.Sprintf("failed to parse json request body: %s", err))
}

func (l *ServerLogger) ErrResourceNotFound(resourceName, resourceId string) {
	l.error(fmt.Sprintf("%s %s not found", resourceName, resourceId))
}

func (l *ServerLogger) ErrServiceRequestStatusUpdateFailed(action, serviceRequestId, reason string) {
	l.error(fmt.Sprintf("failed to %s service request %s: %s", action, serviceRequestId, reason))
}

func (l *ServerLogger) ErrSettingUpJWTValidator(err error) {
	l.error(fmt.Sprintf("failed to set up jwt validator: %s", err))
}

func (l *ServerLogger) ErrValidatingJWT(err error) {
	l.error(fmt.Sprintf("failed to validate jwt: %s", err))
}

func (l *ServerLogger) ErrValidatingPipeline(err error) {
	l.error(fmt.Sprintf("failed to validate pipeline: %s", err))
}

func (l *ServerLogger) ErrEventMissingData(eventName, data string) {
	l.error(fmt.Sprintf("event %s missing data: %s", eventName, data))
}

func (l *ServerLogger) HandleEvent(eventName string) {
	l.info(fmt.Sprintf("handling event: %s", eventName))
}

func (l *ServerLogger) ResourceCreated(resourceName, resourceId string) {
	l.info(fmt.Sprintf("%s %s created", resourceName, resourceId))
}

func (l *ServerLogger) ResourceDeleted(resourceName, resourceId string) {
	l.info(fmt.Sprintf("%s %s deleted", resourceName, resourceId))
}

func (l *ServerLogger) ServerHealthy() {
	l.info("server is healthy")
}

func (l *ServerLogger) ShutdownServer(reason string) {
	l.info(fmt.Sprintf("shutting down server: %s", reason))
}

func (l *ServerLogger) SkippingAdminCheck() {
	l.warn("skipping admin check in dev environment")
}

func (l *ServerLogger) SkippingAuth() {
	l.warn("skipping authentication in dev environment")
}

func (l *ServerLogger) UserLoggedIn(userId string) {
	l.info(fmt.Sprintf("user %s logged in", userId))
}
