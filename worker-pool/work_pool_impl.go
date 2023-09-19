package workerpool // import "github.com/SimifiniiCTO/simfiny-core-lib/worker-pool"

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

// EnqueueTask implements IJobPool.
func (t *TaskPoolProcessor) EnqueueTask(ctx context.Context, task *asynq.Task, delay *time.Duration) (*asynq.TaskInfo, error) {
	if task == nil {
		return nil, fmt.Errorf("invalid task object. task cannot be nil")
	}

	return t.client.EnqueueContext(ctx, task, asynq.ProcessIn(*delay))
}

// NewTask implements IJobPool.
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

func (t *TaskPoolProcessor) ProcessTask(ctx context.Context, task *asynq.Task, f Executor, payload any) error {
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v, %w", err, asynq.SkipRetry)
	}

	// execute the task
	return f.Execute(ctx, payload)
}
