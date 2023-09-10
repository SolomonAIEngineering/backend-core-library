package consumer // import "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/consumer"

import (
	"fmt"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"
	"go.uber.org/zap"
)

type Option func(*ConsumerClient)

// WithMessageProcessTimeout sets the message process timeout for the consumer
func WithMessageProcessTimeout(messageProcessTimeout time.Duration) Option {
	return func(c *ConsumerClient) {
		c.MessageProcessTimeout = messageProcessTimeout
	}
}

// WithQueuePollingDuration sets the queue polling duration for the consumer
func WithQueuePollingDuration(queuePollingDuration time.Duration) Option {
	return func(c *ConsumerClient) {
		c.QueuePollingDuration = queuePollingDuration
	}
}

// WithConcurrencyFactor sets the concurrency factor for the consumer
func WithConcurrencyFactor(concurrencyFactor int) Option {
	return func(c *ConsumerClient) {
		c.ConcurrencyFactor = concurrencyFactor
	}
}

// WithLogger sets the logger for the consumer
func WithLogger(logger *zap.Logger) Option {
	return func(c *ConsumerClient) {
		c.Logger = logger
	}
}

// WithNewRelicClient sets the new relic client for the consumer
func WithInstrumentationClient(instrumentationClient *instrumentation.Client) Option {
	return func(c *ConsumerClient) {
		c.InstrumentationClient = instrumentationClient
	}
}

// WithQueueUrl sets the queue url for the consumer
func WithQueueUrl(queueUrl *string) Option {
	return func(c *ConsumerClient) {
		c.QueueUrl = queueUrl
	}
}

// WithAwsClient sets the aws client for the consumer
func WithAwsClient(sqsClient *client.Client) Option {
	return func(c *ConsumerClient) {
		c.SqsClient = sqsClient
	}
}

// Validate validates whether all the required
// parameters for the consumer client have been set. It checks if the `SqsClient`, `Logger`,
// `NewRelicClient`, `QueueUrl`, `ConcurrencyFactor`, `MessageProcessTimeout`, and
// `QueuePollingDuration` fields are not nil or zero. If any of these fields are nil or zero, it
// returns an error indicating that the consumer client is invalid.
func (c *ConsumerClient) Validate() error {
	if c.SqsClient == nil ||
		c.Logger == nil ||
		c.InstrumentationClient == nil ||
		c.QueueUrl == nil ||
		c.ConcurrencyFactor == 0 ||
		c.MessageProcessTimeout == 0 ||
		c.QueuePollingDuration == 0 {
		return fmt.Errorf("invalid consumer client. params: %v", c)
	}

	return nil
}

// The New function creates a new ConsumerClient instance with optional configuration options.
func New(options ...Option) (*ConsumerClient, error) {
	c := &ConsumerClient{}
	for _, option := range options {
		option(c)
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}
