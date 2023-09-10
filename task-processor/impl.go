package taskprocessor

import (
	"context"

	"github.com/SimifiniiCTO/asynq"
	"go.uber.org/zap"
)

// EnqueueTask implements IProcessor.
func (tp *TaskProcessor) EnqueueTask(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	if task == nil {
		return nil, ErrTaskNotSet
	}

	tp.logger.Info("enqueueing task", zap.Any("task", task))

	return tp.client.EnqueueContext(ctx, task, opts...)
}

// The `EnqueueRecurringTask` function is used to enqueue a recurring task with a specified interval.
// It takes in the context, the task to be enqueued, the interval at which the task should be repeated,
// and optional options for the task. It returns a pointer to a string representing the entry ID of the
// recurring task and an error if any.
func (tp *TaskProcessor) EnqueueRecurringTask(ctx context.Context, task *asynq.Task, interval ProcessingInterval, opts ...asynq.Option) (*string, error) {
	entryID, err := tp.scheduler.Register(interval.String(), task, opts...)
	if err != nil {
		return nil, err
	}

	return &entryID, nil
}

// Start starts the task processor worker as well as the scheduler
// ```go
//
//			tp, err := NewTaskProcessor(...opts)
//			if err != nil {
//				return err
//			}
//
//	     	// start the worker asynchronously in another go routine
//		 	go tp.Start()
//
// ```
func (tp *TaskProcessor) Start() error {
	// start the worker
	if err := tp.worker.Start(); err != nil {
		return err
	}

	// start the scheduler
	if err := tp.scheduler.Start(); err != nil {
		return err
	}

	return nil
}

// Close closes the task processor
//
// ```go
//
//			tp, err := NewTaskProcessor(...opts)
//			if err != nil {
//				return err
//			}
//
//			defer tp.Close()
//	     	// start the worker asynchronously in another go routine
//		 	go func(fn TaskProcessorHandler) {
//				tp.Start(fn)
//			}(fn)
//
// ```
func (tp *TaskProcessor) Close() error {
	tp.worker.Stop()

	// close the redis connection
	if err := tp.client.Close(); err != nil {
		return err
	}

	return nil
}

// Validate validates the task processor
func (tp *TaskProcessor) Validate() error {
	if tp.client == nil {
		return ErrClientNotSet
	}

	if tp.worker == nil {
		return ErrWorkerNotSet
	}

	if tp.logger == nil {
		return ErrLoggerNotSet
	}

	if tp.instrumentationClient == nil {
		return ErrInstrumentationClientNotSet
	}

	if tp.concurrencyFactor == nil {
		return ErrConcurrencyFactorNotSet
	}

	if tp.taskHandler == nil {
		return ErrTaskHandlerNotSet
	}

	return nil
}
