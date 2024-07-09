package batch_job // import "github.com/SolomonAIEngineering/backend-core-library/batch-job"

import (
	"context"
	"errors"

	taskprocessor "github.com/SolomonAIEngineering/backend-core-library/task-processor"
	"github.com/hibiken/asynq"
)

// Runnable is an interface that defines a method to register recurring batch jobs.
type Runnable interface {
	RegisterRecurringBatchJobs(ctx context.Context) error
}

type Batcher struct {
	batchJobs *BatchJobs
	processor *taskprocessor.TaskProcessor
}

var _ Runnable = (*Batcher)(nil)

type BatcherOption func(*Batcher)

// WithBatchJobs sets the batch jobs for the batcher.
func WithBatchJobsRef(batchJobs *BatchJobs) BatcherOption {
	return func(b *Batcher) {
		b.batchJobs = batchJobs
	}
}

// WithTaskProcessor sets the task processor for the batcher.
func WithTaskProcessor(processor *taskprocessor.TaskProcessor) BatcherOption {
	return func(b *Batcher) {
		b.processor = processor
	}
}

// NewBatcher creates a new batcher.
func NewBatcher(options ...BatcherOption) *Batcher {
	batcher := &Batcher{}

	for _, option := range options {
		option(batcher)
	}

	return batcher
}

// RegisterRecurringBatchJobs registers all recurring batch jobs.
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
