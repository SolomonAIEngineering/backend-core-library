package batch_job // import "github.com/SolomonAIEngineering/backend-core-library/batch-job"

import (
	"context"
	"errors"

	taskprocessor "github.com/SolomonAIEngineering/backend-core-library/task-processor"
	"github.com/hibiken/asynq"
)

// Runnable defines the interface for registering recurring batch jobs.
// Implementations of this interface are responsible for setting up and
// configuring batch jobs that need to run on a recurring schedule.
type Runnable interface {
	// RegisterRecurringBatchJobs registers all batch jobs with the task processor.
	// It validates each job's configuration before registration and returns an error
	// if any job fails validation or registration.
	//
	// Example:
	//   batcher := NewBatcher(WithBatchJobs(jobs), WithTaskProcessor(processor))
	//   if err := batcher.RegisterRecurringBatchJobs(ctx); err != nil {
	//       log.Fatal(err)
	//   }
	RegisterRecurringBatchJobs(ctx context.Context) error
}

// Batcher implements the Runnable interface and provides functionality to
// register and manage recurring batch jobs.
type Batcher struct {
	batchJobs *BatchJobs
	processor *taskprocessor.TaskProcessor
}

var _ Runnable = (*Batcher)(nil)

// BatcherOption defines a function type for configuring a Batcher instance.
// This follows the functional options pattern for flexible configuration.
type BatcherOption func(*Batcher)

// WithBatchJobsRef sets the batch jobs configuration for the batcher.
// The BatchJobs parameter contains the collection of jobs to be registered.
//
// Example:
//
//	jobs := &BatchJobs{...}
//	batcher := NewBatcher(WithBatchJobsRef(jobs))
func WithBatchJobsRef(batchJobs *BatchJobs) BatcherOption {
	return func(b *Batcher) {
		b.batchJobs = batchJobs
	}
}

// WithTaskProcessor sets the task processor that will handle the execution
// of batch jobs. The task processor is responsible for managing the job queue
// and executing jobs according to their schedules.
//
// Example:
//
//	processor := taskprocessor.New(...)
//	batcher := NewBatcher(WithTaskProcessor(processor))
func WithTaskProcessor(processor *taskprocessor.TaskProcessor) BatcherOption {
	return func(b *Batcher) {
		b.processor = processor
	}
}

// NewBatcher creates a new Batcher instance with the provided options.
// It uses the functional options pattern to allow flexible configuration.
//
// Example:
//
//	batcher := NewBatcher(
//	    WithBatchJobsRef(jobs),
//	    WithTaskProcessor(processor),
//	)
func NewBatcher(options ...BatcherOption) *Batcher {
	batcher := &Batcher{}

	for _, option := range options {
		option(batcher)
	}

	return batcher
}

// RegisterRecurringBatchJobs implements the Runnable interface.
// It validates and registers each batch job with the task processor.
// If a job with the same ID already exists, it will be skipped without error.
// Any other registration errors will be returned immediately.
func (b *Batcher) RegisterRecurringBatchJobs(ctx context.Context) error {
	for _, job := range b.batchJobs.jobs {
		// first validate the job
		if err := job.Validate(); err != nil {
			return err
		}

		// then register the job
		_, err := b.processor.EnqueueRecurringTask(
			ctx,
			job.task,
			job.ProcessingInterval(),
			asynq.TaskID(*job.taskId),
		)

		if err != nil {
			switch {
			case errors.Is(err, asynq.ErrTaskIDConflict):
				return nil
			case err != nil:
				return err
			}

			return nil
		}
	}

	return nil
}
