package taskprocessor

import (
	"context"
	"errors"
	"time"

	"github.com/SimifiniiCTO/asynq"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-core-lib/task-processor/taskhandler"
	"github.com/SimifiniiCTO/simfiny-core-lib/task-processor/worker"
	"go.uber.org/zap"
)

var (
	// `ErrClientNotSet` is an error that is returned when the `client` property of the `TaskProcessor`
	// struct is not set.
	ErrClientNotSet = errors.New("client not set")
	// `ErrWorkerNotSet` is an error that is returned when the `worker` property of the `TaskProcessor`
	// struct is not set.
	ErrWorkerNotSet = errors.New("worker not set")
	// `ErrLoggerNotSet` is an error that is returned when the `logger` property of the `TaskProcessor`
	// struct is not set.
	ErrLoggerNotSet = errors.New("logger not set")
	// `ErrInstrumentationClientNotSet` is an error that is returned when the `instrumentationClient`
	// property of the `TaskProcessor` struct is not set.
	ErrInstrumentationClientNotSet = errors.New("instrumentation client not set")
	// `ErrConcurrencyFactorNotSet` is an error that is returned when the `concurrencyFactor`
	// property of the `TaskProcessor` struct is not set.
	ErrConcurrencyFactorNotSet = errors.New("concurrently factor not set")
	// `ErrTaskNotSet` is an error that is returned when the `task` argument of the `EnqueueTask` method
	// is not set.
	ErrTaskNotSet = errors.New("task not set")
	// `ErrTaskHandlerNotSet` is an error that is returned when the `taskHandler` property of the
	// `TaskProcessor` struct is not set.
	ErrTaskHandlerNotSet = errors.New("task handler not set")
)

// `TaskProcessorHandler` is a type alias for a function that processes a task. This function is used by
// the `TaskProcessor` to process tasks from the Asynq task queue.
type TaskProcessorHandler func(context.Context, *asynq.Task) error

type TaskProcessor struct {
	// `client *asynq.Client` is a pointer to an instance of the `asynq.Client` struct. This struct is used
	// to interact with the Asynq task queue system, allowing the TaskProcessor to enqueue and dequeue
	// tasks.
	client *asynq.Client
	// `worker *worker.Worker` is a pointer to an instance of the `Worker` struct from the `worker`
	// package. This allows the `TaskProcessor` to use the methods and properties of the `Worker` struct,
	// which is responsible for processing tasks from the Asynq task queue.
	worker *worker.Worker
	// `logger *zap.Logger` is a pointer to an instance of the `zap.Logger` struct. This struct is used for
	// logging messages and events related to the task processing operations performed by the
	// `TaskProcessor`. The `zap` package is a popular logging library in the Go programming language,
	// known for its high performance and flexibility. By using a logger, the `TaskProcessor` can record
	// important information about the task processing operations, such as errors, warnings, and debug
	// messages, which can be useful for troubleshooting and monitoring the system.
	logger *zap.Logger
	// `instrumentationClient *instrumentation.Client` is a pointer to an instance of the
	// `instrumentation.Client` struct. This struct is used for collecting and sending metrics and other
	// performance-related data to a monitoring system. By using an instrumentation client, the
	// `TaskProcessor` can track important metrics related to the task processing operations, such as the
	// number of tasks processed, the processing time for each task, and any errors or failures that occur
	// during processing. This information can be used to monitor the health and performance of the system,
	// identify bottlenecks and other issues, and make improvements to the task processing workflow.
	instrumentationClient *instrumentation.Client
	// `redisConnectionAddress string` is a string containing the address of the Redis server that the
	redisConnectionAddress string
	// concurrencyFactor int is an integer representing the number of concurrent workers that will be
	// processing tasks from the Asynq task queue. This value is used to set the `concurrency` property of
	// the `worker` struct, which determines the number of concurrent workers that will be processing tasks
	// from the Asynq task queue.
	concurrencyFactor *int

	// The `taskHandler *task.TaskHandler` property is a pointer to an instance of the `task.TaskHandler`
	// struct, which is responsible for handling the processing of individual tasks. It is likely used by
	// the `TaskProcessor` to delegate the actual processing of tasks to the `task.TaskHandler` struct,
	// which may contain business logic and other processing steps specific to the type of task being
	// processed. However, without more context on how the `taskHandler` property is used within the
	// `TaskProcessor` struct, it is difficult to say for certain.
	taskHandler taskhandler.ITaskHandler

	// The `scheduler *asynq.Scheduler` property is a pointer to an instance of the `asynq.Scheduler`
	// struct. The `asynq.Scheduler` is responsible for scheduling tasks to be enqueued in the Asynq task
	// queue at a specific time in the future. It allows you to delay the execution of tasks by specifying
	// a delay duration or a specific time at which the task should be enqueued.
	scheduler *asynq.Scheduler
}

// IProcessor is an interface that defines the methods that must be implemented by a task processor
type IProcessor interface {
	// `Start() error` is a method of the `TaskProcessor` struct that starts the worker
	// to process tasks from the Asynq task queue as well as the scheduler process. It takes a `TaskProcessorHandler` function as an
	// argument, which is a function that will be called for each task that is processed by the worker. The
	// `TaskProcessorHandler` function takes a `context.Context` and an `*asynq.Task` as arguments, and
	// returns an error. The `Start` method returns an error if there is a problem starting the worker.
	Start() error
	// `Validate()` is a method of the `TaskProcessor` struct that checks if all the required properties of
	// the struct have been set. It returns an error if any of the required properties are not set,
	// indicating that the `TaskProcessor` is not ready to start processing tasks. This method is typically
	// called before starting the worker to ensure that all the necessary dependencies are in place.
	Validate() error
	// `EnqueueTask` is a method of the `TaskProcessor` struct that is used to enqueue a task in the Asynq
	// task queue. It takes a `context.Context`, an `*asynq.Task`, and an optional variadic argument of
	// `asynq.Option` as arguments. The `ctx` argument is a context that can be used to cancel the task or
	// set a timeout. The `task` argument is a pointer to an instance of the `asynq.Task` struct, which
	// contains information about the task to be executed, such as the task type, payload, and priority.
	// The `opts` argument is a variadic argument of `asynq.Option`, which can be used to set additional
	// options for the task, such as a delay or a custom queue name. The method returns an error if there
	// is a problem enqueueing the task in the Asynq task queue.
	EnqueueTask(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
}

// Option is a function that is used to set properties of the TaskProcessor struct
type Option func(*TaskProcessor)

var _ IProcessor = (*TaskProcessor)(nil)

// NewTaskProcessor creates a new TaskProcessor instance
// ```go
// tp, err := NewTaskProcessor(
//
//	WithInstrumentationClient(ic),
//	WithLogger(logger),
//	WithRedisConnectionAddress("localhost:6379"),
//	WithConcurrencyFactor(10),
//
// )
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer tp.Close()
//	// start the worker asynchronously
//	go func() {
//		if err := tp.Start(handler); err != nil {
//			log.Fatal(err)
//		}
//
// ```
func NewTaskProcessor(opts ...Option) (*TaskProcessor, error) {
	tp := &TaskProcessor{}
	for _, opt := range opts {
		opt(tp)
	}

	// error out if redis address was not provided
	if tp.redisConnectionAddress == "" {
		return nil, errors.New("redis connection address not provided")
	}

	tp.logger.Info("creating task processor for processing", zap.String("redis_connection_address", tp.redisConnectionAddress))

	asyncClientOpt, err := asynq.ParseRedisURI(tp.redisConnectionAddress)
	if err != nil {
		return nil, err
	}

	// define a client
	tp.client = asynq.NewClient(
		asyncClientOpt,
	)

	// initialize new scheudler
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	schedulerOpts := &asynq.SchedulerOpts{
		Location: loc,
		PreEnqueueFunc: func(task *asynq.Task, opts []asynq.Option) {
			// TODO: we need to emit a metric to new relic that we are enqueuing a recurring task
		},
		PostEnqueueFunc: func(info *asynq.TaskInfo, err error) {
			// TODO: we  need to emit a metric to new relic that we have processed a recurring task
		},
		EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
		},
	}

	tp.scheduler = asynq.NewScheduler(asyncClientOpt, schedulerOpts)

	// define the worker
	worker, err := worker.NewWorker([]worker.Option{
		worker.WithConcurrencyFactor(*tp.concurrencyFactor),
		worker.WithRedisAddress(tp.redisConnectionAddress),
		worker.WithTaskHandler(tp.taskHandler),
		worker.WithInstrumentationClient(tp.instrumentationClient),
	}...)
	if err != nil {
		return nil, err
	}

	// set the worker
	tp.worker = worker

	// validate the task processor
	if err := tp.Validate(); err != nil {
		return nil, err
	}

	return tp, nil
}

func WithRedisAddressOpt(redisAddress string) Option {
	return func(tp *TaskProcessor) {
		tp.redisConnectionAddress = redisAddress
	}
}

func WithLoggerOpt(logger *zap.Logger) Option {
	return func(tp *TaskProcessor) {
		tp.logger = logger
	}
}

func WithInstrumentationClientOpt(instrumentationClient *instrumentation.Client) Option {
	return func(tp *TaskProcessor) {
		tp.instrumentationClient = instrumentationClient
	}
}

func WithConcurrencyFactorOpt(concurrencyFactor *int) Option {
	return func(tp *TaskProcessor) {
		tp.concurrencyFactor = concurrencyFactor
	}
}

func WithTaskHandlerOpt(taskHandler taskhandler.ITaskHandler) Option {
	return func(tp *TaskProcessor) {
		tp.taskHandler = taskHandler
	}
}
