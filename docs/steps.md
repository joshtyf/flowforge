# Steps

There are currently only 2 step types available in the pipeline: `WAIT_FOR_APPROVAL` and `API`. More step types can be easily added according to your requirements.

The behavior of each step is defined in `backend/src/execute/steps.go` within the `execute` function of each executor.

### WAIT_FOR_APPROVAL

This step will trigger a pause in the service request. The service request will not proceed until an admin has approved the request.

There are no parameters for this step.

### API

This step will make an API request. The URL, headers, query parameters, request body, and method can be configured when defining the pipeline.

The response and request body will be logged. A non-200 status code will indicate a step failure.

#### Parameters

- `url`: The URL to make the request to.
  - required: `true`
- `headers`: The headers to include in the request.
  - type: `object`
  - required: `true`
- `data`: The data to include in the request body.
  - type: `object`
  - required: `true`
  - notes: Currently, only JSON payload is supported.
- `method`: The HTTP method to use.
  - type: `string`
  - required: `true`

**Example**

```json
{
  "step_name": "Make API Call",
  "step_type": "API",
  "next_step_name": "",
  "prev_step_name": "",
  "parameters": {
    "method": "POST",
    "url": "https://httpbin.org/post?q=${query_param}",
    "data": {
      "key": "${value}"
    },
    "headers": {
      "content-type": "application/json"
    }
  },
  "is_terminal_step": true
}
```

Here, `${query_param}` and `${value}` are placeholders that will be replaced with actual values provided by the user when creating a service request.

## Step execution flow

Service requests are executed sequentially based on an events approach. When a service request is started, a new `NewServiceRequestEvent` will be emitted in the main server and handled by the `StepExecutionManager`. This manager will prepare and trigger the execution of the first step in the pipeline.

After the execution of every step, a new `StepCompletedEvent` will be emitted and handled. The emission of this event can done either at the server side (for the `WAIT_FOR_APPROVAL` step) or by the step executor itself (for the `API` step). The handler of this event will clean up the current step and trigger the execution of the next step in the pipeline if any.

### Step event types

The step event types are not to be confused with the macro events that are handled by the `StepExecutionManager`. Step event types are lifecycle events that happen within the execution of a step. The following are the step event types (found in `backend/src/database/models/service_request_event.go`):

```go

type EventType string

const (
	STEP_NOT_STARTED EventType = "Not Started"
	STEP_RUNNING     EventType = "Running"
	STEP_FAILED      EventType = "Failed"
	STEP_CANCELLED   EventType = "Cancelled"
	STEP_COMPLETED   EventType = "Completed"
)
```

When a service request is first created, every step will have the initial event of `STEP_NOT_STARTED`. When a step is being executed, the event will be `STEP_RUNNING`. If the step fails, the event will be `STEP_FAILED`. If the step is cancelled, the event will be `STEP_CANCELLED`. If the step is successfully executed, the event will be `STEP_COMPLETED`.

These events will only be **appended** in the database table, and never replaced. This immutable log design will allow for tracking of the lifecycle of the step execution. The latest event of any step is used to determine the current state of the step.

## Creating new step types

To define a new step type for the pipeline, you must first define a struct that implements the `stepExecutor` interface. The behavior of the step should be defined in the `execute` function of the struct. There should also be a corresponding `PipelineStepType` constant defined in the pipeline model.

Lastly, register the new executor with the `StepExecutionManager`.

### Getting execution contextual information about the step

The `execute` function of the `stepExecutor` interface takes in a `context` parameter. This context will contain information about the current service request and the current step. The information can be accessed by using the correct context keys.

**Examples**

```go
func (e *stepExecutor) execute(ctx context.Context, l *logger.ExecutorLogger) (*stepExecResult, error) {
    step, ok := ctx.Value(util.StepKey).(*models.PipelineStepModel)
	serviceRequest, ok := ctx.Value(util.ServiceRequestKey).(*models.ServiceRequestModel)
    step_parameters := step.Parameters
    // Rest of code here
}
```

### Logging steps

The `execute` function of the `stepExecutor` interface takes in an `ExecutorLogger` parameter. This logger is prepared and initialised by the `StepExecutionManager`. When used, the logger will log messages to the step's log file within the service requests logs directory. All service requests logs can be found in the base directory at `./executor_logs`.

For example, a service request with `id=1` and step `step_name=Make API Call` will have a log file located at `./executor_logs/1/Make API Call.log`.

**Note**: steps which are not started will not have a log file created for them.
