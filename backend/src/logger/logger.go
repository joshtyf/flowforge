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

func (l *Logger) ErrClaimsMissingPermission(permission string) {
	l.error(fmt.Sprintf("unauthorized: missing permission %s", permission))
}

func (l *Logger) ErrEventLogStepMissingInPipeline(stepName string) {
	l.error(fmt.Sprintf("%s exists in event log but not in pipeline template", stepName))
}

func (l *Logger) ErrExecutingStep(stepName string, err error) {
	l.error(fmt.Sprintf("error executing %s: %s", stepName, err))
}

func (l *Logger) ErrHandlingEvent(err error) {
	l.error(fmt.Sprintf("error encountered while handling event: %s", err))
}

func (l *Logger) ErrHandlingRequest(err error) {
	l.error(fmt.Sprintf("error encountered while handling API request: %s", err))
}

func (l *Logger) ErrInvalidStepType(expectedStepType, actualStepType models.PipelineStepType) {
	l.error(fmt.Sprintf("invalid step type: expected %s got %s", expectedStepType, actualStepType))
}

func (l *Logger) ErrMissingExecutorForStep(stepName string) {
	l.error(fmt.Sprintf("missing executor for step: %s", stepName))
}

func (l *Logger) ErrMissingPipelineStep(stepName string) {
	l.error(fmt.Sprintf("missing pipeline step: %s", stepName))
}

func (l *Logger) ErrParsingIssuerURL(err error) {
	l.error(fmt.Sprintf("failed to parse the issuer url: %s", err))
}

func (l *Logger) ErrParsingJsonRequestBody(err error) {
	l.error(fmt.Sprintf("failed to parse json request body: %s", err))
}

func (l *Logger) ErrResourceNotFound(resourceName, resourceId string) {
	l.error(fmt.Sprintf("%s %s not found", resourceName, resourceId))
}

func (l *Logger) ErrServiceRequestStatusUpdateFailed(action, serviceRequestId, reason string) {
	l.error(fmt.Sprintf("failed to %s service request %s: %s", action, serviceRequestId, reason))
}

func (l *Logger) ErrSettingUpJWTValidator(err error) {
	l.error(fmt.Sprintf("failed to set up jwt validator: %s", err))
}

func (l *Logger) ErrValidatingJWT(err error) {
	l.error(fmt.Sprintf("failed to validate jwt: %s", err))
}

func (l *Logger) ErrValidatingPipeline(err error) {
	l.error(fmt.Sprintf("failed to validate pipeline: %s", err))
}

func (l *Logger) ErrEventMissingData(eventName, data string) {
	l.error(fmt.Sprintf("event %s missing data: %s", eventName, data))
}

func (l *Logger) HandleEvent(eventName string) {
	l.info(fmt.Sprintf("handling event: %s", eventName))
}

func (l *Logger) ResourceCreated(resourceName, resourceId string) {
	l.info(fmt.Sprintf("%s %s created", resourceName, resourceId))
}

func (l *Logger) ResourceDeleted(resourceName, resourceId string) {
	l.info(fmt.Sprintf("%s %s deleted", resourceName, resourceId))
}

func (l *Logger) ServerHealthy() {
	l.info("server is healthy")
}

func (l *Logger) ShutdownServer(reason string) {
	l.info(fmt.Sprintf("shutting down server: %s", reason))
}

func (l *Logger) SkippingAdminCheck() {
	l.warn("skipping admin check in dev environment")
}

func (l *Logger) SkippingAuth() {
	l.warn("skipping authentication in dev environment")
}

func (l *Logger) UserLoggedIn(userId string) {
	l.info(fmt.Sprintf("user %s logged in", userId))
}
