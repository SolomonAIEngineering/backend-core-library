package cacher // import "github.com/SolomonAIEngineering/backend-core-library/cacher"

import (
	"errors"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

type Option func(*Client)

// WithServiceName configures the service name
func WithServiceName(name string) Option {
	return func(c *Client) {
		c.serviceName = name
	}
}

// WithLogger configures the logger
func WithLogger(logger *zap.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithRedisConn configures the redis connection
func WithRedisConn(connPool *redis.Pool) Option {
	return func(c *Client) {
		c.pool = connPool
	}
}

// WithIntrumentationClient configures whether or not instrumentation is enabled
func WithIntrumentationClient(client *instrumentation.Client) Option {
	return func(c *Client) {
		c.instrumentationClient = client
	}
}

// WithCacheTTLInSeconds configures the cache TTL in seconds
func WithCacheTTLInSeconds(ttl int) Option {
	return func(c *Client) {
		c.cacheTTLInSeconds = ttl
	}
}

// Validate validates the cacher
func (c *Client) Validate() error {
	if c.pool == nil {
		return errors.New("redis connection is nil")
	}

	if c.logger == nil {
		return errors.New("logger is nil")
	}

	if c.serviceName == "" {
		return errors.New("service name is empty")
	}

	if c.cacheTTLInSeconds == 0 {
		return errors.New("cache TTL is 0")
	}

	if c.instrumentationClient == nil {
		c.logger.Warn("instrumentation client is nil")
	}

	return nil
}
