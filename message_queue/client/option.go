package client // import "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// Option is a function that configures the client.
type Option func(*Client)

// WithConfig sets the Config for the client
func WithConfig(config *Config) Option {
	return func(c *Client) {
		c.Config = config
	}
}

// WithTimeout sets the timeout for the client
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.Timeout = timeout
	}
}

// WithReadTimeout sets the read timeout for the client
func WithReadTimeout(readTimeout time.Duration) Option {
	return func(c *Client) {
		c.ReadTimeout = readTimeout
	}
}

// WithWriteTimeout sets the write timeout for the client
func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(c *Client) {
		c.WriteTimeout = writeTimeout
	}
}

// WithClient sets the client for the client
func WithSqsClient(client sqsiface.SQSAPI) Option {
	return func(c *Client) {
		c.SqsClient = client
	}
}

// WithAwsConfig sets the aws config for the client
func WithAwsConfig(config *AwsConfig) Option {
	return func(c *Client) {
		c.AwsConfig = config
	}
}

// Validate validates the client
func (c *Client) Validate() error {
	if c.SqsClient == nil {
		return fmt.Errorf("sqs client not set")
	}

	if c.Timeout == 0 {
		return fmt.Errorf("timeout not set")
	}

	if c.ReadTimeout == 0 {
		return fmt.Errorf("read timeout not set")
	}

	if c.WriteTimeout == 0 {
		return fmt.Errorf("write timeout not set")
	}

	if c.AwsConfig == nil {
		return fmt.Errorf("aws config not set")
	}

	if c.Config == nil {
		return fmt.Errorf("config not set")
	}

	return nil
}

// The Config type defines parameters for receiving messages from Amazon SQS.
// @property MaxNumberOfMessages - This property specifies the maximum number of messages that can be
// returned by an Amazon SQS (Simple Queue Service) request. The value can range from 1 to 10, and the
// default value is 1.
// @property MaxWaitTimeInSeconds - The duration (in seconds) that the received messages are hidden
// from subsequent retrieve requests after being retrieved by a ReceiveMessage request. This means that
// if a message is retrieved by a ReceiveMessage request, it will not be visible to other requests for
// the duration specified in this property. After the duration has elapsed
// @property Attributes - A list of attributes that need to be returned along with each message. These
// attributes provide additional information about the message, such as its unique identifier,
// timestamp, and message body. The list of available attributes includes MessageId, ReceiptHandle,
// MD5OfBody, Body, and others. By default,
type Config struct {
	// The maximum number of messages to return. Amazon SQS never returns
	// more messages than this value (however, fewer messages might be returned).
	// Valid values: 1 to 10. Default: 1.
	MaxNumberOfMessages *int
	// The duration (in seconds) that the received messages are hidden from subsequent
	// retrieve requests after being retrieved by a ReceiveMessage request.
	MaxWaitTimeInSeconds *time.Duration
	// A list of attributes that need to be returned along with each message.
	Attributes *[]string
}

// ConfigOption is a function that sets a configuration option such as MaxNumberOfMessages
type ConfigOption func(*Config)

// WithMaxNumberOfMessages sets the MaxNumberOfMessages configuration option
func WithMaxNumberOfMessages(maxNumberOfMessages int) ConfigOption {
	return func(c *Config) {
		c.MaxNumberOfMessages = &maxNumberOfMessages
	}
}

// WithMaxWaitTimeSeconds sets the MaxWaitTimeSeconds configuration option
func WithMaxWaitTimeSeconds(MaxWaitTimeInSeconds time.Duration) ConfigOption {
	return func(c *Config) {
		c.MaxWaitTimeInSeconds = &MaxWaitTimeInSeconds
	}
}

// WithAttributes sets the Attributes configuration option
func WithAttributes(attributes []string) ConfigOption {
	return func(c *Config) {
		c.Attributes = &attributes
	}
}

// NewConfigOptions returns a new Config with the given options
func NewConfigOptions(opts ...ConfigOption) *Config {
	c := &Config{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
