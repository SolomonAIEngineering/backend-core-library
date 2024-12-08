package workerpool // import "github.com/SolomonAIEngineering/backend-core-library/worker-pool"

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

// EnqueueTask enqueues a task to be processed by the worker pool with a specified delay.
// It returns the task information and any error encountered during enqueuing.
//
// Example usage:
//
//	delay := 5 * time.Minute
//	taskInfo, err := pool.EnqueueTask(ctx, task, &delay)
//	if err != nil {
//	    log.Printf("failed to enqueue task: %v", err)
//	    return err
//	}
func (t *TaskPoolProcessor) EnqueueTask(ctx context.Context, task *asynq.Task, delay *time.Duration) (*asynq.TaskInfo, error) {
	if task == nil {
		return nil, fmt.Errorf("invalid task object. task cannot be nil")
	}

	return t.client.EnqueueContext(ctx, task, asynq.ProcessIn(*delay))
}

// NewTask creates a new task with the specified parameters and options.
// It marshals the taskPayload into JSON and sets up the task with retry and timeout configurations.
//
// Parameters:
//   - taskId: unique identifier for the task
//   - taskType: type/category of the task (e.g., "email:send", "image:resize")
//   - scheduledTime: optional time when the task should be executed
//   - taskPayload: data required for task execution (will be marshaled to JSON)
//
// Example usage:
//
//	type EmailPayload struct {
//	    UserID    string
//	    Template  string
//	    Variables map[string]string
//	}
//
//	payload := EmailPayload{
//	    UserID:    "123",
//	    Template:  "welcome_email",
//	    Variables: map[string]string{"name": "John"},
//	}
//	task, err := pool.NewTask("task-123", "email:send", nil, payload)
func (t *TaskPoolProcessor) NewTask(taskId, taskType string, scheduledTime *time.Time, taskPayload any) (*asynq.Task, error) {
	payload, err := json.Marshal(taskPayload)
	if err != nil {
		return nil, err
	}

	opts := []asynq.Option{
		asynq.MaxRetry(*t.maxRetry),
		asynq.Timeout(*t.taskTimeout),
		asynq.TaskID(taskId),
	}

	return asynq.NewTask(taskType, payload, opts...), nil
}

// ProcessTask unmarshals the task payload and executes the task using the provided executor.
// It returns an error if either the unmarshal operation fails or the task execution fails.
//
// Parameters:
//   - ctx: context for task execution
//   - task: the task to be processed
//   - f: executor implementing the Execute method
//   - payload: pointer to a struct where the task payload will be unmarshaled
//
// Example usage:
//
//	type EmailExecutor struct{}
//
//	func (e *EmailExecutor) Execute(ctx context.Context, payload any) error {
//	    emailPayload, ok := payload.(*EmailPayload)
//	    if !ok {
//	        return fmt.Errorf("invalid payload type")
//	    }
//	    // Process the email task
//	    return nil
//	}
//
//	var payload EmailPayload
//	err := pool.ProcessTask(ctx, task, &EmailExecutor{}, &payload)
func (t *TaskPoolProcessor) ProcessTask(ctx context.Context, task *asynq.Task, f Executor, payload any) error {
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v, %w", err, asynq.SkipRetry)
	}

	// execute the task
	return f.Execute(ctx, payload)
}
