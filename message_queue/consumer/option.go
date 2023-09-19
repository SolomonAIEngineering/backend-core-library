package consumer // import "github.com/SolomonAIEngineering/backend-core-library/message_queue/consumer"

import (
	"fmt"
	"time"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

type Option func(*ConsumerClient)

func WithLogger(logger *zap.Logger) Option {
	return func(opt *ConsumerClient) {
		opt.logger = logger
	}
}

func WithInstrumentationClient(ic *instrumentation.Client) Option {
	return func(opt *ConsumerClient) {
		opt.instrumentationClient = ic
	}
}

func WithSQSClient(client *sqs.SQS) Option {
	return func(opt *ConsumerClient) {
		opt.sqsClient = client
	}
}

func WithQueueURL(url *string) Option {
	return func(opt *ConsumerClient) {
		opt.queueUrl = url
	}
}

func WithDLQURL(dlqURL *string) Option {
	return func(opt *ConsumerClient) {
		opt.deadletterQueueUrl = dlqURL
	}
}

func WithConcurrencyFactor(factor int) Option {
	return func(opt *ConsumerClient) {
		opt.concurrencyFactor = factor
	}
}

func WithQueuePollingDuration(duration time.Duration) Option {
	return func(opt *ConsumerClient) {
		opt.queuePollingDuration = duration
	}
}

func WithMessageProcessTimeout(timeout time.Duration) Option {
	return func(opt *ConsumerClient) {
		opt.messageProcessTimeout = timeout
	}
}

func WithMessageHandler(handler MessageProcessorFunc) Option {
	return func(opt *ConsumerClient) {
		opt.handler = handler
	}
}

func WithBackoffDuration(duration time.Duration) Option {
	return func(opt *ConsumerClient) {
		opt.backoffDuration = duration
	}
}

func WithBatchSize(size int64) Option {
	return func(opt *ConsumerClient) {
		opt.batchSize = size
	}
}

func WithWaitTimeSecond(waitTime int64) Option {
	return func(opt *ConsumerClient) {
		opt.waitTimeSecond = waitTime
	}
}

// Validate validates whether all the required
// parameters for the consumer client have been set. It checks if the `SqsClient`, `Logger`,
// `NewRelicClient`, `QueueUrl`, `ConcurrencyFactor`, `MessageProcessTimeout`, and
// `QueuePollingDuration` fields are not nil or zero. If any of these fields are nil or zero, it
// returns an error indicating that the consumer client is invalid.
func (c *ConsumerClient) Validate() error {
	if c.logger == nil ||
		c.instrumentationClient == nil ||
		c.sqsClient == nil ||
		c.queueUrl == nil ||
		c.deadletterQueueUrl == nil ||
		c.concurrencyFactor == 0 ||
		c.queuePollingDuration == 0 ||
		c.messageProcessTimeout == 0 ||
		c.handler == nil ||
		c.backoffDuration == 0 ||
		c.batchSize == 0 ||
		c.waitTimeSecond == 0 {
		return fmt.Errorf("invalid consumer client. params: %v", c)
	}

	return nil
}

// The New function creates a new ConsumerClient instance with optional configuration options.
//
// example:
// consumer := NewConsumerClient(
//
//	WithLogger(yourLogger),
//	WithInstrumentationClient(yourInstrumentationClient),
//	WithSQSClient(yourSQSClient),
//	WithQueueURL(&yourQueueURL),
//	WithDLQURL(&yourDLQURL),
//	WithConcurrencyFactor(yourConcurrencyFactor),
//	WithQueuePollingDuration(yourPollingDuration),
//	WithMessageProcessTimeout(yourMessageTimeout),
//	WithMessageHandler(yourMessageHandlerFunc),
//	WithBackoffDuration(yourBackoffDuration),
//	WithBatchSize(yourBatchSize),
//	WithWaitTimeSecond(yourWaitTimeSeconds),
//
// )
func New(options ...Option) (*ConsumerClient, error) {
	c := &ConsumerClient{
		stopCh: make(chan bool),
	}

	for _, option := range options {
		option(c)
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	// enable infinite polls if initialized through api
	c.maxPolls = -1

	return c, nil
}
