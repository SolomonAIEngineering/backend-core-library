package postgres // import "github.com/SolomonAIEngineering/backend-core-library/database/postgres"

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
		conn.MaxConnectionRetries = retries
	}
}

// WithMaxConnectionRetryTimeout sets the max connection retry timeout
func WithMaxConnectionRetryTimeout(timeout *time.Duration) Option {
	return func(conn *Client) {
		conn.MaxConnectionRetryTimeout = timeout
	}
}

// WithRetrySleep sets the retry sleep
func WithRetrySleep(sleep *time.Duration) Option {
	return func(conn *Client) {
		conn.RetrySleep = sleep
	}
}

// WithConnectionString sets the connection string
func WithConnectionString(connectionString *string) Option {
	return func(conn *Client) {
		conn.ConnectionString = connectionString
	}
}

// Validate validates the client
func (c *Client) Validate() error {
	if c.QueryTimeout == nil {
		return fmt.Errorf("query timeout is nil")
	}

	if c.MaxConnectionRetries == nil {
		return fmt.Errorf("max connection retries is nil")
	}

	if c.MaxConnectionRetryTimeout == nil {
		return fmt.Errorf("max connection retry timeout is nil")
	}

	if c.RetrySleep == nil {
		return fmt.Errorf("retry sleep is nil")
	}

	if c.ConnectionString == nil {
		return fmt.Errorf("connection string is nil")
	}

	if c.MaxIdleConnections == nil {
		return fmt.Errorf("max idle connections is nil")
	}

	if c.MaxOpenConnections == nil {
		return fmt.Errorf("max open connections is nil")
	}

	if c.MaxConnectionLifetime == nil {
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

// WithMaxIdleConnections sets the max idle connections
func WithMaxIdleConnections(maxIdleConnections *int) Option {
	return func(conn *Client) {
		conn.MaxIdleConnections = maxIdleConnections
	}
}

// WithMaxOpenConnections sets the max open connections
func WithMaxOpenConnections(maxOpenConnections *int) Option {
	return func(conn *Client) {
		conn.MaxOpenConnections = maxOpenConnections
	}
}

// WithMaxConnectionLifetime sets the max connection lifetime
func WithMaxConnectionLifetime(maxConnectionLifetime *time.Duration) Option {
	return func(conn *Client) {
		conn.MaxConnectionLifetime = maxConnectionLifetime
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
