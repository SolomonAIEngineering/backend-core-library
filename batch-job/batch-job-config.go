package batch_job // import "github.com/SolomonAIEngineering/backend-core-library/batch-job"

import (
	"fmt"

	"github.com/hibiken/asynq"
)

// BatchJob represents a configurable batch job with its associated task and identifier.
// It extends BaseConfig with task-specific configuration for asynq processing.
type BatchJob struct {
	BaseConfig
	task   *asynq.Task
	taskId *string
}

// BatchJobConfigOption defines a function type for configuring BatchJob instances.
// This follows the functional options pattern for flexible configuration.
type BatchJobConfigOption func(*BatchJob)

// Validate ensures that the BatchJob is properly configured with all required fields.
// Returns an error if any required field is missing or if the base configuration is invalid.
func (c *BatchJob) Validate() error {
	if c.task == nil {
		return fmt.Errorf("task is nil")
	}

	if c.taskId == nil {
		return fmt.Errorf("taskId is nil")
	}

	return c.BaseConfig.Validate()
}

// String provides a string representation of the BatchJob for debugging and logging purposes.
//
// Example output: "BatchJob: cfg - {BaseConfig}, taskId - abc123, task - {Task}"
func (c *BatchJob) String() string {
	return fmt.Sprintf("BatchJob: cfg - %v, taskId - %s, task - %v", c.BaseConfig, *c.taskId, c.task)
}

// WithBaseConfig sets the base configuration for the batch job.
// This includes common settings like processing interval and retry policy.
//
// Example:
//
//	baseCfg := BaseConfig{...}
//	job := NewBatchJob(WithBaseConfig(baseCfg))
func WithBaseConfig(cfg BaseConfig) BatchJobConfigOption {
	return func(b *BatchJob) {
		b.BaseConfig = cfg
	}
}

// WithTask sets the asynq task for the batch job.
// The task contains the actual work to be performed when the job executes.
//
// Example:
//
//	task := asynq.NewTask(...)
//	job := NewBatchJob(WithTask(task))
func WithTask(task *asynq.Task) BatchJobConfigOption {
	return func(b *BatchJob) {
		b.task = task
	}
}

// WithTaskId sets the unique identifier for the batch job.
// This ID is used to prevent duplicate job registrations.
//
// Example:
//
//	taskId := "unique-job-id"
//	job := NewBatchJob(WithTaskId(&taskId))
func WithTaskId(taskId *string) BatchJobConfigOption {
	return func(b *BatchJob) {
		b.taskId = taskId
	}
}

// NewBatchJob creates a new BatchJob instance with the provided options.
// It uses the functional options pattern to allow flexible configuration.
//
// Example:
//
//	taskId := "job-123"
//	job := NewBatchJob(
//	    WithBaseConfig(baseCfg),
//	    WithTask(task),
//	    WithTaskId(&taskId),
//	)
func NewBatchJob(options ...BatchJobConfigOption) *BatchJob {
	config := &BatchJob{}

	for _, option := range options {
		option(config)
	}

	return config
}
