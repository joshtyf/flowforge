package execute

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gookit/event"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/events"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type stepExecResult struct {
}

type stepExecutor interface {
	execute(context.Context, *logger.ExecutorLogger) (*stepExecResult, error)
	getStepType() models.PipelineStepType
}

type apiStepExecutor struct {
}

func NewApiStepExecutor() *apiStepExecutor {
	return &apiStepExecutor{}
}

func (e *apiStepExecutor) execute(ctx context.Context, l *logger.ExecutorLogger) (*stepExecResult, error) {
	step, ok := ctx.Value(util.StepKey).(*models.PipelineStepModel)
	if !ok {
		l.ErrGettingStepFromCtx()
		return nil, errors.New("error getting step from context")
	}
	serviceRequest, ok := ctx.Value(util.ServiceRequestKey).(*models.ServiceRequestModel)
	if !ok {
		l.ErrGettingServiceReqFromCtx()
		return nil, errors.New("error getting service request from context")
	}
	requestMethod := step.Parameters["method"]
	url := step.Parameters["url"]
	req, err := http.NewRequest(requestMethod, url, nil)
	l.HttpRequest("GET", url)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	l.HttpResponseStatus(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response")
	}
	event.FireAsync(events.NewStepCompletedEvent(step.StepName, serviceRequest, &stepExecResult{}, nil))
	return &stepExecResult{}, nil
}

func (e *apiStepExecutor) getStepType() models.PipelineStepType {
	return models.APIStep
}

type waitForApprovalStepExecutor struct {
	mongoClient *mongo.Client
}

func NewWaitForApprovalStepExecutor(mongoClient *mongo.Client) *waitForApprovalStepExecutor {
	return &waitForApprovalStepExecutor{
		mongoClient: mongoClient,
	}
}

func (e *waitForApprovalStepExecutor) execute(ctx context.Context, l *logger.ExecutorLogger) (*stepExecResult, error) {
	l.WaitingForApproval()
	serviceRequest, ok := ctx.Value(util.ServiceRequestKey).(*models.ServiceRequestModel)
	if !ok {
		l.ErrGettingServiceReqFromCtx()
		return nil, errors.New("error getting service request from context")
	}
	err := database.NewServiceRequest(e.mongoClient).UpdateStatus(serviceRequest.Id.Hex(), models.PENDING)
	if err != nil {
		l.ErrUpdatingServiceRequestStatus(err)
		return nil, err
	}
	return &stepExecResult{}, nil
}

func (e *waitForApprovalStepExecutor) getStepType() models.PipelineStepType {
	return models.WaitForApprovalStep
}
