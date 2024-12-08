package cacher // import "github.com/SolomonAIEngineering/backend-core-library/cacher"

import (
	"errors"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

// Option defines a function type that configures a Client.
// Options are used with New() to configure the client instance.
type Option func(*Client)

// WithServiceName sets the service name used for key prefixing.
// The service name is prepended to all cache keys to prevent collisions
// when multiple services share the same Redis instance.
//
// Example:
//
//	client, err := cacher.New(
//	    cacher.WithServiceName("user-service"),
//	)
func WithServiceName(name string) Option {
	return func(c *Client) {
		c.serviceName = name
	}
}

// WithLogger configures the logging instance for the cache client.
// The logger is used to record operational events and errors.
//
// Example:
//
//	logger := zap.NewProduction()
//	client, err := cacher.New(
//	    cacher.WithLogger(logger),
//	)
func WithLogger(logger *zap.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithRedisConn configures the Redis connection pool.
// The pool manages a set of Redis connections for optimal performance.
//
// Example:
//
//	pool := &redis.Pool{
//	    MaxIdle:     3,
//	    IdleTimeout: 240 * time.Second,
//	    Dial: func() (redis.Conn, error) {
//	        return redis.Dial("tcp", "localhost:6379")
//	    },
//	}
//	client, err := cacher.New(
//	    cacher.WithRedisConn(pool),
//	)
func WithRedisConn(connPool *redis.Pool) Option {
	return func(c *Client) {
		c.pool = connPool
	}
}

// WithIntrumentationClient enables distributed tracing for cache operations.
// When configured, each cache operation will create a trace span for monitoring.
//
// Example:
//
//	instrClient := instrumentation.NewClient()
//	client, err := cacher.New(
//	    cacher.WithIntrumentationClient(instrClient),
//	)
func WithIntrumentationClient(client *instrumentation.Client) Option {
	return func(c *Client) {
		c.instrumentationClient = client
	}
}

// WithCacheTTLInSeconds sets the default time-to-live for cache entries in seconds.
// This TTL is used for all cache writes unless overridden by WriteToCacheWithTTL.
//
// Example:
//
//	// Set default TTL to 1 hour
//	client, err := cacher.New(
//	    cacher.WithCacheTTLInSeconds(3600),
//	)
func WithCacheTTLInSeconds(ttl int) Option {
	return func(c *Client) {
		c.cacheTTLInSeconds = ttl
	}
}

// Validate checks that all required fields are properly configured.
// It returns an error if any required configuration is missing or invalid.
//
// Required configurations:
// - Redis connection pool
// - Logger
// - Service name
// - Cache TTL > 0
//
// The instrumentation client is optional - a warning is logged if not provided.
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
