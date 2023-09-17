package redis

import (
	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"go.uber.org/zap"
)

// Option configures the redis client properly
type Option func(*Client)

// WithLogger sets the logger for the Redis client.
func WithLogger(logger *zap.Logger) Option {
	return func(c *Client) {
		c.Logger = logger
	}
}

// WithTelemetrySdk sets the telemetry SDK for the Redis client.
func WithTelemetrySdk(sdk *instrumentation.Client) Option {
	return func(c *Client) {
		c.telemetrySdk = sdk
	}
}

// WithURI sets the URI for the Redis client.
func WithURI(uri string) Option {
	return func(c *Client) {
		c.URI = uri
	}
}

// WithServiceName sets the service name for the Redis client.
func WithServiceName(serviceName string) Option {
	return func(c *Client) {
		c.serviceName = serviceName
	}
}

// WithCacheTTLInSeconds sets the cache TTL in seconds for the Redis client.
func WithCacheTTLInSeconds(cacheTTLInSeconds int) Option {
	return func(c *Client) {
		c.cacheTTLInSeconds = cacheTTLInSeconds
	}
}

func WithTlsEnabled(enabled bool) Option {
	return func(c *Client) {
		c.tlsEnabled = enabled
	}
}

// Validate validates the configuration of the Redis client.
func (c *Client) Validate() error {
	if c.URI == "" {
		return ErrInvalidURI
	}

	if c.serviceName == "" {
		return ErrInvalidServiceName
	}

	if c.Logger == nil {
		return ErrInvalidLogger
	}

	if c.telemetrySdk == nil {
		return ErrInvalidTelemetrySdk
	}

	if c.pool == nil {
		return ErrInvalidPool
	}

	if c.cacheTTLInSeconds <= 0 {
		return ErrInvalidCacheTTLInSeconds
	}

	return nil
}
