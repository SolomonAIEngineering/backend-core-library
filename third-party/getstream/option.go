package getstream // import "github.com/SolomonAIEngineering/backend-core-library/third-party/getstream"

import (
	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"go.uber.org/zap"
)

type Option func(*Client)

// The function WithLogger takes a pointer to a zap.Logger and returns an Option.
func WithLogger(logger *zap.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

// The function WithKey takes a string and returns an Option.
func WithKey(key string) Option {
	return func(c *Client) {
		c.getstreamKey = key
	}
}

// The function WithSecret takes a string and returns an Option.
func WithSecret(secret string) Option {
	return func(c *Client) {
		c.getstreamSecret = secret
	}
}

// The function WithInstrumentationClient takes a pointer to a instrumentation.Client and returns an Option.
func WithInstrumentationClient(instrumentationClient *instrumentation.Client) Option {
	return func(c *Client) {
		c.instrumentationClient = instrumentationClient
	}
}

// The function WithClient takes a pointer to a stream.Client and returns an Option.
func (c *Client) Validate() error {
	if c.getstreamKey == "" || c.getstreamSecret == "" {
		return ErrInvalidKeyOrSecret
	}

	if c.logger == nil {
		return ErrInvalidLogger
	}

	if c.instrumentationClient == nil {
		return ErrInvalidInstrumentationClient
	}

	return nil
}

// The function ApplyOptions takes a variadic number of Options and applies them to the Client.
func (c *Client) ApplyOptions(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}
