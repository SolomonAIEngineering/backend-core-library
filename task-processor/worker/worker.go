package worker // import "github.com/SolomonAIEngineering/backend-core-library/task-processor/worker"

import (
	"context"
	"errors"
	"fmt"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/SolomonAIEngineering/backend-core-library/task-processor/taskhandler"
)

// `type Worker struct` is defining a new struct type called `Worker`. This struct type has three
// fields: `redisAddress` of type string, `concurrencyFactor` of type int, and `srv` of type
// `*asynq.Server`. These fields represent the configuration and state of a worker that will process
// tasks from a Redis server using the `asynq` package.
type Worker struct {
	// `redisAddress` is a field of the `Worker` struct that holds the address of the Redis server that the
	// worker will connect to. It is a string type variable.
	redisAddress string
	// `concurrencyFactor` is a field of the `Worker` struct that represents the maximum number of tasks
	// that the worker can process concurrently. It is used to configure the `asynq` server instance that
	// the worker creates. The `WithConcurrencyFactor` option function sets the value of this field during
	// the creation of a new worker instance.
	concurrencyFactor int
	// `srv` is a field of the `Worker` struct that holds an instance of the `asynq.Server` type. This
	// server instance is created during the creation of a new worker instance using the `NewServer`
	// function from the `asynq` package. The `asynq.Server` type is responsible for managing the worker's
	// connection to the Redis server and processing tasks from the task queues. The `srv` field is used in
	// the `Start` and `Stop` methods of the `Worker` struct to start and stop the server instance.
	srv *asynq.Server

	// The `taskHandler` field of the `Worker` struct is of type `task.TaskHandler`. This field represents
	// an interface that defines the methods for handling different types of tasks. It allows the worker to
	// delegate the processing of tasks to the appropriate handler based on the type of the task. The
	// implementation of the `TaskHandler` interface is provided by the `task` package, which defines the
	// specific handlers for each type of task.
	taskHandler taskhandler.ITaskHandler

	// The `instrumentationClient` field is a pointer to an instance of the `instrumentation.Client` type.
	// This field is used to send metrics and traces to a monitoring system for observability purposes. The
	// `instrumentation` package provides a client implementation for sending metrics and traces to various
	// monitoring systems such as Datadog, Prometheus, and OpenTelemetry. By including this field in the
	// `Worker` struct, the worker can send metrics and traces to a monitoring system to help diagnose and
	// troubleshoot issues with the task processing system.
	instrumentationClient *instrumentation.Client
}

var (
	ErrRedisAddressNotSet          = errors.New("redis address not set")
	ErrRedisPasswordNotSet         = errors.New("redis password not set")
	ErrConcurrencyFactorNotSet     = errors.New("concurrency factor not set")
	ErrTaskHandlerNotSet           = errors.New("task handler not set")
	ErrInstrumentationClientNotSet = errors.New("instrumentation client not set")
)

// Option is a function that configures a worker
type Option func(*Worker)

// NewWorker creates a new worker
//
// ```go
//
//	worker, err := NewWorker(
//		WithRedisAddress("localhost:6379"),
//		WithConcurrencyFactor(10),
//	)
//	if err != nil {
//		return err
//	}
//
//	defer worker.Stop()
//
//	// start the worker asynchronously in another go routine
//	go worker.Start()
//
// ```
func NewWorker(opts ...Option) (*Worker, error) {
	w := &Worker{}
	for _, opt := range opts {
		opt(w)
	}

	// validate the worker
	if err := w.Validate(); err != nil {
		return nil, err
	}

	asyncClientOpt, err := asynq.ParseRedisURI(w.redisAddress)
	if err != nil {
		return nil, err
	}

	// create a new worker srv instance
	w.srv = asynq.NewServer(
		asyncClientOpt,
		asynq.Config{Concurrency: w.concurrencyFactor},
	)

	return w, nil
}

// Start starts the worker
// ```go
//
//	worker, err := NewWorker(
//		WithRedisAddress("localhost:6379"),
//		WithConcurrencyFactor(10),
//	)
//	if err != nil {
//		return err
//	}
//
//	defer worker.Stop()
//
//	// start the worker asynchronously in another go routine
//	go worker.Start()
//
// ```
func (w *Worker) Start() error {
	mux := w.taskHandler.RegisterTaskHandler()
	mux.Use(w.instrumentationMiddleware)

	return w.srv.Start(mux)
}

// `func (w *Worker) instrumentationMiddleware(next asynq.Handler) asynq.Handler` is a middleware
// function that wraps the `asynq.Handler` passed as an argument to it. It returns a new
// `asynq.Handler` that adds instrumentation to the processing of tasks by the worker. Specifically, it
// creates a new trace span for each task processed by the worker and sends metrics to a monitoring
// system to track the success or failure of each task. The middleware function takes the
// `asynq.Handler` passed as an argument and returns a new `asynq.Handler` that wraps the original
// handler. The new handler executes the original handler and then sends metrics and traces to the
// monitoring system based on the success or failure of the task.
func (w *Worker) instrumentationMiddleware(next asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		tnxName := fmt.Sprintf("processing_task/%s", t.Type())
		txn := w.instrumentationClient.GetTraceFromContext(ctx)
		span := w.instrumentationClient.StartSegment(txn, tnxName)
		defer span.End()

		err := next.ProcessTask(ctx, t)
		// TODO: decrement task in progress counter metric
		if err != nil {
			txnName := fmt.Sprintf("task_failed/%s", t.Type())
			w.instrumentationClient.RecordMetric(txnName, 1)
		}

		txnName := fmt.Sprintf("task_suceeded/%s", t.Type())
		w.instrumentationClient.RecordMetric(txnName, 1)
		return err
	})
}

// Stop stops the worker
// ```go
//
//	worker, err := NewWorker(
//		WithRedisAddress("localhost:6379"),
//		WithConcurrencyFactor(10),
//	)
//	if err != nil {
//		return err
//	}
//
//	defer worker.Stop()
func (w *Worker) Stop() {
	// signals the server to stop pulling new tasks off queues.
	// Stop is used before shutting down the server to ensure that all currently active tasks are
	// processed before server shutdown.
	w.srv.Stop()

	// shutdown the server
	w.srv.Shutdown()
}

// `func (w *Worker) Validate() error` is a method of the `Worker` struct that validates whether the
// required fields of the worker have been set or not. It returns an error if any of the required
// fields are not set. This method is called during the creation of a new worker instance to ensure
// that the worker is properly configured before it is started.
func (w *Worker) Validate() error {
	if w.redisAddress == "" {
		return ErrRedisAddressNotSet
	}

	if w.concurrencyFactor == 0 {
		return ErrConcurrencyFactorNotSet
	}

	if w.taskHandler == nil {
		return ErrTaskHandlerNotSet
	}

	if w.instrumentationClient == nil {
		return ErrInstrumentationClientNotSet
	}

	// validate the worker
	return nil
}

// WithRedisAddress sets the redis address
func WithRedisAddress(address string) Option {
	return func(w *Worker) {
		w.redisAddress = address
	}
}

// WithConcurrencyFactor sets the concurrency factor
func WithConcurrencyFactor(concurrencyFactor int) Option {
	return func(w *Worker) {
		w.concurrencyFactor = concurrencyFactor
	}
}

// WithTaskHandler sets the task handler
func WithTaskHandler(taskHandler taskhandler.ITaskHandler) Option {
	return func(w *Worker) {
		w.taskHandler = taskHandler
	}
}

// WithInstrumentationClient sets the instrumentation client
func WithInstrumentationClient(instrumentationClient *instrumentation.Client) Option {
	return func(w *Worker) {
		w.instrumentationClient = instrumentationClient
	}
}
