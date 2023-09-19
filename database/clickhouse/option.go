package clickhouse // import "github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"

import (
	"fmt"
	"time"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"go.uber.org/zap"
)

// Option is a function that sets a parameter for the client
type Option func(*Client)

// WithQueryTimeout sets the query timeout
func WithQueryTimeout(timeout *time.Duration) Option {
	return func(conn *Client) {
		conn.QueryTimeout = timeout
	}
}

// WithMaxConnectionRetries sets the max connection retries
func WithMaxConnectionRetries(retries *int) Option {
	return func(conn *Client) {
		conn.maxConnectionRetries = retries
	}
}

// WithMaxConnectionRetryTimeout sets the max connection retry timeout
func WithMaxConnectionRetryTimeout(timeout *time.Duration) Option {
	return func(conn *Client) {
		conn.maxConnectionRetryTimeout = timeout
	}
}

// WithRetrySleep sets the retry sleep
func WithRetrySleep(sleep *time.Duration) Option {
	return func(conn *Client) {
		conn.retrySleep = sleep
	}
}

// WithConnectionString sets the connection string
func WithConnectionString(connectionString *string) Option {
	return func(conn *Client) {
		conn.connectionURI = connectionString
	}
}

// WithMaxIdleConnections sets the max idle connections
func WithMaxIdleConnections(maxIdleConnections *int) Option {
	return func(conn *Client) {
		conn.maxIdleConnections = maxIdleConnections
	}
}

// WithMaxOpenConnections sets the max open connections
func WithMaxOpenConnections(maxOpenConnections *int) Option {
	return func(conn *Client) {
		conn.maxOpenConnections = maxOpenConnections
	}
}

// WithMaxConnectionLifetime sets the max connection lifetime
func WithMaxConnectionLifetime(maxConnectionLifetime *time.Duration) Option {
	return func(conn *Client) {
		conn.maxConnectionLifetime = maxConnectionLifetime
	}
}

// WithInstrumentationClient sets the instrumentation client
func WithInstrumentationClient(instrumentationClient *instrumentation.Client) Option {
	return func(conn *Client) {
		conn.InstrumentationClient = instrumentationClient
	}
}

// WithLogger sets the logger
func WithLogger(logger *zap.Logger) Option {
	return func(conn *Client) {
		conn.Logger = logger
	}
}

// Validate validates the client
func (c *Client) Validate() error {
	if c.QueryTimeout == nil {
		return fmt.Errorf("query timeout is nil")
	}

	if c.maxConnectionRetries == nil {
		return fmt.Errorf("max connection retries is nil")
	}

	if c.maxConnectionRetryTimeout == nil {
		return fmt.Errorf("max connection retry timeout is nil")
	}

	if c.retrySleep == nil {
		return fmt.Errorf("retry sleep is nil")
	}

	if c.connectionURI == nil {
		return fmt.Errorf("connection string is nil")
	}

	if c.maxIdleConnections == nil {
		return fmt.Errorf("max idle connections is nil")
	}

	if c.maxOpenConnections == nil {
		return fmt.Errorf("max open connections is nil")
	}

	if c.maxConnectionLifetime == nil {
		return fmt.Errorf("max connection lifetime is nil")
	}

	if c.InstrumentationClient == nil {
		return fmt.Errorf("instrumentation client is nil")
	}

	if c.Logger == nil {
		return fmt.Errorf("logger is nil")
	}

	return nil
}
