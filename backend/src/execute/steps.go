package execute

import (
	"context"
	"fmt"
	"net/http"

	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/util"
)

type stepExecResult struct {
}

type stepExecutor interface {
	execute(ctx context.Context) (*stepExecResult, error)
	getStepType() models.PipelineStepType
}

type apiStepExecutor struct {
}

func NewApiStepExecutor() *apiStepExecutor {
	return &apiStepExecutor{}
}

func (e *apiStepExecutor) execute(ctx context.Context) (*stepExecResult, error) {
	step, ok := ctx.Value(util.StepKey).(*models.PipelineStepModel)
	if !ok {
		logger.Error("[APIStepExecutor] Error getting step from context", nil)
		return nil, fmt.Errorf("error getting step from context")
	}
	url := step.Parameters["url"]
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error("[APIStepExecutor] Error creating request", map[string]interface{}{"err": err})
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("[APIStepExecutor] Error sending request", map[string]interface{}{"err": err})
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Error("[APIStepExecutor] Non-200 response", map[string]interface{}{"status": resp.StatusCode})
		return nil, fmt.Errorf("Non-200 response")
	}
	logger.Info("[APIStepExecutor] Success", map[string]interface{}{"status": resp.StatusCode})
	return &stepExecResult{}, nil
}

func (e *apiStepExecutor) getStepType() models.PipelineStepType {
	return models.APIStep
}
