package consumer // import "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/consumer"

import (
	"context"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	client "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"
	"go.uber.org/zap"
)

// MessageProcessorFunc serves as the logic used to process each incoming message from a msg queue
type MessageProcessorFunc = func(ctx context.Context, message *client.Message) error

// IConsumer defines an interface for a consumer that can process messages either concurrently or
// naively.
// @property ConcurrentConsumer - A method that takes a MessageProcessorFunc as an argument and
// processes messages concurrently. This means that multiple messages can be processed at the same
// time, potentially improving performance.
// @property NaiveConsumer - A method that takes a MessageProcessorFunc as an argument and processes
// messages in a single-threaded, sequential manner. This method is not optimized for performance and
// may not be suitable for high-throughput scenarios.
type IConsumer interface {
	// `StartConcurentConsumer(f MessageProcessorFunc)` is a method of the `IConsumer` interface that takes a
	// `MessageProcessorFunc` as an argument and processes messages concurrently. This means that multiple
	// messages can be processed at the same time, potentially improving performance. The implementation of
	// this method is not shown in the given code snippet.
	StartConcurentConsumer(f MessageProcessorFunc)
	// `NaiveConsumer(f MessageProcessorFunc)` is a method of the `IConsumer` interface that takes a
	// `MessageProcessorFunc` as an argument and processes messages in a single-threaded, sequential
	// manner. This means that messages will be processed one at a time, in the order in which they were
	// received. This method is not optimized for performance and may not be suitable for high-throughput
	// scenarios.
	NaiveConsumer(f MessageProcessorFunc)
}

// The ConsumerClient type contains various fields related to a client that consumes messages from a
// queue.
// @property Logger - A pointer to a zap logger, which is a logging library for Go. It is used to log
// messages and errors during the execution of the consumer client.
// @property NewRelicClient - NewRelicClient is a pointer to a newrelic.Application object, which is
// used for monitoring and tracing application performance. It allows for the collection of metrics and
// tracing of requests across distributed systems.
// @property Client - The `Client` property is a pointer to an instance of the `client.Client` struct.
// This struct is likely used to interact with some external service or API.
// @property QueueUrl - The URL of the queue from which the consumer client will receive messages.
// @property {int} ConcurrencyFactor - ConcurrencyFactor is an integer property that represents the
// number of concurrent message processing operations that can be performed by the ConsumerClient. It
// determines how many messages can be processed simultaneously from the queue.
// @property QueuePollingDuration - QueuePollingDuration is a property of the ConsumerClient struct
// that represents the duration for which the client will poll the queue for new messages. It is of
// type time.Duration, which is a built-in Go type that represents a duration of time. The value of
// this property is set by the user and
// @property MessageProcessTimeout - MessageProcessTimeout is a property of the ConsumerClient struct
// and it represents the maximum amount of time that a message can be processed before timing out. If
// the message processing takes longer than this duration, it will be considered as failed and the
// message will be returned to the queue for processing again.
type ConsumerClient struct {
	// `Logger` is a pointer to a `zap.Logger` object, which is a logging library for Go. It is used to log
	// messages and errors during the execution of the consumer client. This allows for easy debugging and
	// monitoring of the client's behavior.
	Logger *zap.Logger
	// `NewRelicClient` is a pointer to a `newrelic.Application` object, which is used for monitoring and
	// tracing application performance. It allows for the collection of metrics and tracing of requests
	// across distributed systems. This suggests that the consumer client is being monitored and tracked
	// for performance using New Relic.
	InstrumentationClient *instrumentation.Client
	// `Client` is a pointer to an instance of the `client.Client` struct. This struct is likely used to
	// interact with some external service or API. It is not clear from the given code what specific
	// service or API the `client.Client` is interacting with.
	SqsClient *client.Client
	// The `QueueUrl` property is a pointer to a string that represents the URL of the queue from which the
	// consumer client will receive messages. It is likely used by the `client.Client` to connect to the
	// queue and retrieve messages. The use of a pointer to a string allows for the value of the URL to be
	// easily updated or changed if necessary.
	QueueUrl *string
	// `ConcurrencyFactor` is an integer property of the `ConsumerClient` struct that represents the number
	// of concurrent message processing operations that can be performed by the client. It determines how
	// many messages can be processed simultaneously from the queue.
	ConcurrencyFactor int
	// `QueuePollingDuration` is a property of the `ConsumerClient` struct that represents the duration for
	// which the client will poll the queue for new messages. It is of type `time.Duration`, which is a
	// built-in Go type that represents a duration of time. The value of this property is set by the user
	// and determines how long the client will wait for new messages before checking the queue again.
	QueuePollingDuration time.Duration
	// `MessageProcessTimeout` is a property of the `ConsumerClient` struct that represents the maximum
	// amount of time that a message can be processed before timing out. If the message processing takes
	// longer than this duration, it will be considered as failed and the message will be returned to the
	// queue for processing again.
	MessageProcessTimeout time.Duration
}

// Checks if the `Consumer` struct implements
// the `IConsumer` interface. It creates a variable of type `IConsumer` and assigns it the value of
// `(*Consumer)(nil)`, which is a pointer to a `Consumer` struct with a nil value. If the `Consumer`
// struct implements all the methods of the `IConsumer` interface, this statement will compile without
// errors. If it does not, the compiler will throw an error indicating which methods are missing. This
// is a way to ensure that the `Consumer` struct satisfies the `IConsumer` interface at compile time.
var _ IConsumer = (*ConsumerClient)(nil)
