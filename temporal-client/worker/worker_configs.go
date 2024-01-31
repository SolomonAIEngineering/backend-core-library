package worker

import (
	"log"

	temporalerrors "github.com/SolomonAIEngineering/backend-core-library/temporal-client/temporal-errors"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
)

// WorkerConfig holds the configuration for the Temporal worker.
type WorkerConfig struct {
	TaskQueue               string
	MaxConcurrentActivities int
	MaxConcurrentWorkflows  int
	WorkerOptions           worker.Options
	Client                  client.Client
	Logger                  *zap.Logger
	Activities              []any
	Workflows               []any
}

// WorkerOption defines a type for functional options.
type WorkerOption func(*WorkerConfig)

// WithTaskQueue sets the task queue for the worker.
func WithTaskQueue(taskQueue string) WorkerOption {
	return func(wc *WorkerConfig) {
		wc.TaskQueue = taskQueue
	}
}

// WithMaxConcurrentActivities sets the maximum number of concurrent activities.
func WithMaxConcurrentActivities(max int) WorkerOption {
	return func(wc *WorkerConfig) {
		wc.MaxConcurrentActivities = max
	}
}

// WithMaxConcurrentWorkflows sets the maximum number of concurrent workflows.
func WithMaxConcurrentWorkflows(max int) WorkerOption {
	return func(wc *WorkerConfig) {
		wc.MaxConcurrentWorkflows = max
	}
}

// WithWorkerOptions sets additional options for the worker.
func WithWorkerOptions(options worker.Options) WorkerOption {
	return func(wc *WorkerConfig) {
		wc.WorkerOptions = options
	}
}

// WithClient sets the Temporal client for the worker.
func WithClient(client client.Client) WorkerOption {
	return func(wc *WorkerConfig) {
		wc.Client = client
	}
}

// WithLogger sets the logger for the worker.
func WithLogger(logger *zap.Logger) WorkerOption {
	return func(wc *WorkerConfig) {
		wc.Logger = logger
	}
}

func WithActivities(activities ...any) WorkerOption {
	return func(wc *WorkerConfig) {
		wc.Activities = activities
	}
}

// WithMaxConcurrentActivities sets the maximum number of concurrent activities.
func WithWorkflows(workflows ...any) WorkerOption {
	return func(wc *WorkerConfig) {
		wc.Workflows = workflows
	}
}

// NewWorkerConfig initializes a WorkerConfig with the given options.
func NewWorkerConfig(opts ...WorkerOption) *WorkerConfig {
	config := &WorkerConfig{}
	for _, opt := range opts {
		opt(config)
	}

	if err := config.validate(); err != nil {
		log.Fatal(err)
	}

	return config
}

// validate checks the validity of the WorkerConfig settings.
// It ensures that the configuration is not nil, the client is set, the task queue is specified,
// and the values for MaxConcurrentActivities and MaxConcurrentWorkflows are positive.
//
// This method returns an error if any of the following conditions are met:
//   - The WorkerConfig itself is nil, indicating it was not properly initialized.
//   - The Client field is nil, which is required to establish a connection to the Temporal server.
//   - The TaskQueue string is empty. A valid task queue name is essential for the worker to receive tasks.
//   - MaxConcurrentActivities is less than or equal to 0. A positive value is required to determine
//     the maximum number of activities that can be executed concurrently.
//   - MaxConcurrentWorkflows is less than or equal to 0. A positive value is needed to set the maximum
//     number of workflows that can be handled concurrently.
//
// If any of these checks fail, the method returns an error from the temporalerrors package,
// corresponding to the specific validation that failed.
func (wc *WorkerConfig) validate() error {
	if wc == nil {
		return temporalerrors.ErrNilWorkerConfig
	}

	if wc.Client == nil {
		return temporalerrors.ErrNilClient
	}

	if wc.TaskQueue == "" {
		return temporalerrors.ErrEmptyTaskQueue
	}

	if wc.MaxConcurrentActivities <= 0 {
		return temporalerrors.ErrInvalidMaxConcurrentActivities
	}

	if wc.MaxConcurrentWorkflows <= 0 {
		return temporalerrors.ErrInvalidMaxConcurrentWorkflows
	}

	if len(wc.Activities) == 0 {
		return ErrNoActivities
	}

	if len(wc.Workflows) == 0 {
		return ErrNoWorkflows
	}

	return nil
}
