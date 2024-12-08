package workerpool // import "github.com/SolomonAIEngineering/backend-core-library/worker-pool"

import (
	"context"
	"errors"
	"time"

	"github.com/hibiken/asynq"
	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type (
	// TaskPoolProcessor manages asynchronous task processing using Redis-backed queues.
	// It provides a robust worker pool implementation for handling distributed tasks
	// with features like retry mechanisms, timeouts, and concurrent processing.
	//
	// The processor uses asynq for task queue management and requires Redis for
	// persistence and message broker capabilities.
	TaskPoolProcessor struct {
		client            *asynq.Client
		redisClient       *redis.Client
		logger            *zap.Logger
		taskResultTTL     *time.Duration
		worker            *asynq.Server
		maxRetry          *int
		taskTimeout       *time.Duration
		RedisAddress      *string
		RedisUserName     *string
		RedisPassword     *string
		ConcurrencyFactor *int
	}

	// `Option` is a functional option pattern used to modify the behavior of the `TaskPoolProcessor`
	// struct. It is a function that takes a pointer to a `TaskPoolProcessor` instance as its argument and
	// modifies its fields. This pattern is commonly used in Go to provide flexible and extensible APIs.
	Option func(processor *TaskPoolProcessor)

	// IJobPool defines the interface for task pool operations.
	// Implementations must provide methods for creating, processing,
	// and enqueueing tasks in a distributed task processing system.
	IJobPool interface {
		// NewTask creates a new task with the specified parameters.
		//
		// Parameters:
		//   - taskId: unique identifier for the task
		//   - taskType: classification of the task (e.g., "email", "notification")
		//   - scheduledTime: optional time when the task should be executed
		//   - taskPayload: data required for task execution
		//
		// Returns:
		//   - *asynq.Task: the created task object
		//   - error: if task creation fails
		NewTask(taskId, taskType string, scheduledTime *time.Time, taskPayload any) (*asynq.Task, error)

		// ProcessTask executes the given task using the provided executor.
		//
		// Parameters:
		//   - ctx: context for task execution
		//   - task: the task to be processed
		//   - f: executor implementing the task logic
		//   - payload: data needed for task execution
		//
		// Returns:
		//   - error: if task processing fails
		ProcessTask(ctx context.Context, task *asynq.Task, f Executor, payload any) error

		// EnqueueTask schedules a task for future execution.
		//
		// Parameters:
		//   - ctx: context for task enqueuing
		//   - task: the task to be scheduled
		//   - delay: optional duration to delay task execution
		//
		// Returns:
		//   - *asynq.TaskInfo: information about the enqueued task
		//   - error: if task enqueuing fails
		EnqueueTask(ctx context.Context, task *asynq.Task, delay *time.Duration) (*asynq.TaskInfo, error)
	}

	// Executor defines the interface for task execution logic.
	// Implementations should contain the actual business logic
	// for processing specific types of tasks.
	Executor interface {
		// Execute runs the task-specific logic with the given payload.
		//
		// Parameters:
		//   - ctx: context for execution
		//   - payload: task-specific data needed for execution
		//
		// Returns:
		//   - error: if execution fails
		Execute(ctx context.Context, payload any) error
	}
)

// `var _ IJobPool = (*TaskPoolProcessor)(nil)` is a compile-time check that ensures that the
// `TaskPoolProcessor` struct satisfies the `IJobPool` interface. It creates a variable of type
// `IJobPool` and assigns it a pointer to a `TaskPoolProcessor` instance that is `nil`. If the
// `TaskPoolProcessor` struct does not implement all the methods defined in the `IJobPool` interface,
// the code will not compile and an error will be thrown. This is a way to catch errors early in the
// development process and ensure that the code is correct.
var _ IJobPool = (*TaskPoolProcessor)(nil)

// NewTaskPoolProcessor creates a new TaskPoolProcessor with the provided options.
// It initializes the Redis connection, worker pool, and task processing configuration.
//
// Parameters:
//   - opts: variadic list of Option functions to configure the processor
//
// Returns:
//   - *TaskPoolProcessor: configured processor instance
//   - error: if initialization fails
func NewTaskPoolProcessor(opts ...Option) (*TaskPoolProcessor, error) {
	processor := &TaskPoolProcessor{}
	for _, opt := range opts {
		opt(processor)
	}

	if err := processor.Validate(); err != nil {
		return nil, err
	}

	processor.worker = asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     *processor.RedisAddress,
			Username: *processor.RedisUserName,
			Password: *processor.RedisPassword,
		},
		// redis has a connection limit hence this factor must be below that.
		// understand for example if the concurrency factor is 10, and you have
		// 3 instances of a service with asynq processing enabled, you will have initiated
		// 30 redis connections .... NOTE: super important to keep this in mind
		asynq.Config{Concurrency: *processor.ConcurrencyFactor},
	)

	processor.client = asynq.NewClient(
		asynq.RedisClientOpt{
			Addr:     *processor.RedisAddress,
			Username: *processor.RedisUserName,
			Password: *processor.RedisPassword,
		})

	// start the worker pool
	pool := NewWorkerPool(*processor.ConcurrencyFactor)
	pool.Start()

	return processor, nil
}

func (p *TaskPoolProcessor) Run(handler asynq.Handler) error {
	return p.worker.Run(handler)
}

func (p *TaskPoolProcessor) Validate() error {
	if p.client == nil {
		return errors.New("invalid task pool client")
	}

	if p.logger == nil {
		return errors.New("invalid task pool logger")
	}

	if p.maxRetry == nil {
		return errors.New("invalid task pool max retry")
	}

	if p.redisClient == nil {
		return errors.New("invalid task pool redis client")
	}

	if p.taskResultTTL == nil {
		return errors.New("invalid task task result TTL")
	}

	if p.taskTimeout == nil {
		return errors.New("invalid task task timeout")
	}

	if p.RedisAddress == nil {
		return errors.New("invalid redis address")
	}

	if p.RedisUserName == nil {
		return errors.New("invalid redis user name")
	}

	if p.RedisPassword == nil {
		return errors.New("invalid redis password")
	}

	if p.ConcurrencyFactor == nil {
		return errors.New("invalid concurrency factor")
	}

	return nil
}

// This function takes an asynq client as input and returns an option.
func WithJobPoolClient(client *asynq.Client) Option {
	return func(p *TaskPoolProcessor) {
		p.client = client
	}
}

// WithRedisClient returns a new TaskPoolProcessor that uses the given
// Redis connection
func WithRedisClient(client *redis.Client) Option {
	return func(p *TaskPoolProcessor) {
		p.redisClient = client
	}
}

// WithLogger returns a new TaskPoolProcessor with the given
// logger
func WithLogger(logger *zap.Logger) Option {
	return func(p *TaskPoolProcessor) {
		p.logger = logger
	}
}

// WithTaskResultTTL returns a new TaskPoolProcessor with the given
// task result
func WithTaskResultTTL(ttl *time.Duration) Option {
	return func(p *TaskPoolProcessor) {
		p.taskResultTTL = ttl
	}
}

func WithRedisAddress(addr *string) Option {
	return func(p *TaskPoolProcessor) {
		p.RedisAddress = addr
	}
}

func WithRedisUsername(username *string) Option {
	return func(p *TaskPoolProcessor) {
		p.RedisUserName = username
	}
}

func WithRedisPassword(password *string) Option {
	return func(p *TaskPoolProcessor) {
		p.RedisPassword = password
	}
}

func WithConcurrencyFactor(factor *int) Option {
	return func(p *TaskPoolProcessor) {
		p.ConcurrencyFactor = factor
	}
}

// WithMaxRetry sets the maximum number of retry attempts for a given task
func WithMaxRetry(maxRetry *int) Option {
	return func(p *TaskPoolProcessor) {
		p.maxRetry = maxRetry
	}
}

func WithTaskTimeout(timeout *time.Duration) Option {
	return func(p *TaskPoolProcessor) {
		p.taskTimeout = timeout
	}
}
