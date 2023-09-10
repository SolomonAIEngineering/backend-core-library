package workerpool

import (
	"context"
	"errors"
	"time"

	"github.com/hibiken/asynq"
	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type (
	// The `TaskPoolProcessor` struct is defining a new type that represents a worker pool processor. It
	// contains fields for an `asynq.Client` instance, a Redis client instance, a logger instance, a task
	// result time-to-live duration, and an `asynq.Server` instance representing the worker. This struct is
	// used to manage the worker pool and process tasks asynchronously.
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

	// The `IJobPool` interface is defining a set of methods that must be implemented by any type that
	// wants to act as a job pool for the `TaskPoolProcessor`. It specifies the behavior that the job pool
	// must have, including creating new tasks, processing tasks, and enqueuing tasks for processing. By
	// defining this interface, the `TaskPoolProcessor` can be decoupled from any specific implementation
	// of the job pool and can work with any type that satisfies the `IJobPool` interface. This makes the
	// `TaskPoolProcessor` more flexible and extensible.
	IJobPool interface {
		// `NewTask` is a method defined in the `IJobPool` interface. It takes a `taskId` string and a pointer
		// to a `scheduledTime` of type `time.Time` as arguments and returns a pointer to an `asynq.Task`
		// instance and an error. This method is used to create a new task with the given `taskId` and
		// `scheduledTime` and return it as an `asynq.Task` instance. The `asynq.Task` instance can then be
		// used to enqueue the task for processing by the worker pool.
		NewTask(taskId, taskType string, scheduledTime *time.Time, taskPayload any) (*asynq.Task, error)
		// The `ProcessTask` method is defined in the `IJobPool` interface and takes a context and an
		// `asynq.Task` instance as arguments. It is used to process the given task asynchronously. The
		// implementation of this method will vary depending on the specific task being processed, but it
		// typically involves performing some kind of computation or I/O operation. The method returns an error
		// if there was a problem processing the task.
		ProcessTask(ctx context.Context, task *asynq.Task, f Executor, payload any) error
		// The `EnqueueTask` method is defined in the `IJobPool` interface and is used to enqueue a task for
		// processing by the worker pool. It takes a context, a pointer to an `asynq.Task` instance, and a
		// pointer to a `time.Duration` representing the delay before the task should be processed as
		// arguments. The method returns a pointer to an `asynq.TaskInfo` instance and an error. The
		// `asynq.TaskInfo` instance contains information about the enqueued task, such as its ID and its
		// scheduled time. This method is typically used to add a new task to the worker pool for processing at
		// a later time.
		EnqueueTask(ctx context.Context, task *asynq.Task, delay *time.Duration) (*asynq.TaskInfo, error)
	}

	// The `Executor` interface is defining a new type that represents an executor for a task. It specifies
	// a single method `Execute` that takes a context and a payload as arguments and returns an error. This
	// interface is used to define the behavior of the function that will be executed when a task is
	// processed by the worker pool. By defining this interface, the `TaskPoolProcessor` can be decoupled
	// from any specific implementation of the task executor and can work with any type that satisfies the
	// `Executor` interface. This makes the `TaskPoolProcessor` more flexible and extensible.
	Executor interface {
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
