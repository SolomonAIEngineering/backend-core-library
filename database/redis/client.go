package redis

import (
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

// The `Client` type contains various fields related to a Redis client in a Go program, including a
// connection pool, cache TTL, telemetry SDK, URI, and logger.
// @property pool - A redis connection pool used to manage connections to a Redis server.
// @property {string} serviceName - The `serviceName` property is a string that represents the name of
// the service that is using this `Client` struct.
// @property {int} cacheTTLInSeconds - cacheTTLInSeconds is a property of the Client struct that
// represents the time-to-live (TTL) value in seconds for cached data. This property is used to
// determine how long cached data should be considered valid before it needs to be refreshed or
// retrieved again from the data source.
// @property telemetrySdk - The `telemetrySdk` property is a pointer to an instance of the `Client`
// struct from the `instrumentation` package. This is likely used for collecting and reporting
// telemetry data related to the Redis client's usage and performance.
// @property {string} URI - The URI property is likely a string that represents the connection URI for
// the Redis database that the client will be interacting with. It could include information such as
// the host, port, and authentication credentials needed to establish a connection.
// @property Logger - The Logger property is a pointer to a zap.Logger instance, which is a structured,
// leveled logging library for Go. It is used to log messages and events related to the Client struct
// and its operations.
type Client struct {
	// `pool *redis.Pool` is a field of the `Client` struct that represents a connection pool used to
	// manage connections to a Redis server. The `redis.Pool` type is a built-in type in the
	// `gomodule/redigo/redis` package that provides a thread-safe connection pool for Redis connections.
	// By using a connection pool, the client can reuse existing connections to the Redis server instead of
	// creating new connections for each request, which can improve performance and reduce resource usage.
	pool *redis.Pool
	// The `serviceName` field is a string that represents the name of the service that is using the
	// `Client` struct. It is likely used for logging and telemetry purposes to identify which service is
	// making requests to the Redis database.
	serviceName string
	// `cacheTTLInSeconds` is a property of the `Client` struct that represents the time-to-live (TTL)
	// value in seconds for cached data. This property is used to determine how long cached data should be
	// considered valid before it needs to be refreshed or retrieved again from the data source.
	cacheTTLInSeconds int
	// `telemetrySdk` is a pointer to an instance of the `Client` struct from the `instrumentation`
	// package. This is likely used for collecting and reporting telemetry data related to the Redis
	// client's usage and performance.
	telemetrySdk *instrumentation.Client
	// The `URI` field is a string that represents the connection URI for the Redis database that the
	// client will be interacting with. It could include information such as the host, port, and
	// authentication credentials needed to establish a connection.
	URI string
	// The `Logger` field is a pointer to a `zap.Logger` instance, which is a structured, leveled logging
	// library for Go. It is used to log messages and events related to the `Client` struct and its
	// operations. This allows for more detailed and organized logging, which can be helpful for debugging
	// and troubleshooting issues with the Redis client.
	Logger *zap.Logger
	// The `tlsEnabled` field is a boolean value that is likely used to indicate whether or not the Redis
	// client should use Transport Layer Security (TLS) encryption when connecting to the Redis server. If
	// `tlsEnabled` is set to `true`, the client will use TLS encryption to secure the connection to the
	// Redis server. If it is set to `false`, the client will not use TLS encryption.
	tlsEnabled bool
}

// CloseCacheConn closes the redis connection
// NOTE: this function should be ran as soon as the server is initialized
//
// `defer s.Close()`
func (c *Client) Close() {
	if c.pool != nil {
		_ = c.pool.Close()
	}
}

// New creates a new client with optional configurations and a stop channel.
func New(stopCh <-chan struct{}, opts ...Option) (*Client, error) {
	client := &Client{}
	for _, opt := range opts {
		opt(client)
	}

	// start redis connection pool
	ticker := time.NewTicker(30 * time.Second)
	client.startCachePool(ticker, stopCh)

	if err := client.Validate(); err != nil {
		return nil, err
	}

	return client, nil
}
