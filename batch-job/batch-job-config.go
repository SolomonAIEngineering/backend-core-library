package batch_job

import (
	"fmt"

	"github.com/SimifiniiCTO/asynq"
)

type BatchJob struct {
	BaseConfig
	task   *asynq.Task
	taskId *string
}

type BatchJobConfigOption func(*BatchJob)

func (c *BatchJob) Validate() error {
	if c.task == nil {
		return fmt.Errorf("task is nil")
	}

	if c.taskId == nil {
		return fmt.Errorf("taskId is nil")
	}

	return c.BaseConfig.Validate()
}

func (c *BatchJob) String() string {
	return fmt.Sprintf("BatchJob: cfg - %v, taskId - %s, task - %v", c.BaseConfig, *c.taskId, c.task)
}

func WithBaseConfig(cfg BaseConfig) BatchJobConfigOption {
	return func(b *BatchJob) {
		b.BaseConfig = cfg
	}
}

func WithTask(task *asynq.Task) BatchJobConfigOption {
	return func(b *BatchJob) {
		b.task = task
	}
}

func WithTaskId(taskId *string) BatchJobConfigOption {
	return func(b *BatchJob) {
		b.taskId = taskId
	}
}

func NewBatchJob(options ...BatchJobConfigOption) *BatchJob {
	config := &BatchJob{}

	for _, option := range options {
		option(config)
	}

	return config
}
