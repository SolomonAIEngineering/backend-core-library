package mongo // import "github.com/SolomonAIEngineering/backend-core-library/database/mongo"

import (
	"fmt"
	"time"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Option func(*Client)

// WithMaxConnectionAttempts sets the maximum connection attempts to initiate against the database
func WithMaxConnectionAttempts(maxConnectionAttempts int) Option {
	return func(c *Client) {
		c.MaxConnectionAttempts = maxConnectionAttempts
	}
}

// WithMaxRetriesPerOperation sets the maximum retries to attempt per failed database connection attempt
func WithMaxRetriesPerOperation(maxRetriesPerOperation int) Option {
	return func(c *Client) {
		c.MaxRetriesPerOperation = maxRetriesPerOperation
	}
}

// WithRetryTimeOut sets the maximum time until a retry operation is observed as a timed out operation
func WithRetryTimeOut(retryTimeOut time.Duration) Option {
	return func(c *Client) {
		c.RetryTimeOut = retryTimeOut
	}
}

// WithOperationSleepInterval sets the amount of time between retry operations that the system sleeps
func WithOperationSleepInterval(operationSleepInterval time.Duration) Option {
	return func(c *Client) {
		c.OperationSleepInterval = operationSleepInterval
	}
}

// WithQueryTimeout sets the maximal amount of time a query can execute before being cancelled
func WithQueryTimeout(queryTimeout time.Duration) Option {
	return func(c *Client) {
		c.QueryTimeout = queryTimeout
	}
}

// WithDatabaseName sets the database name
func WithDatabaseName(databaseName string) Option {
	return func(c *Client) {
		c.DatabaseName = &databaseName
	}
}

// WithTelemetry sets the object by which we will emit metrics, trace requests, and database operations
func WithTelemetry(telemetry *instrumentation.Client) Option {
	return func(c *Client) {
		c.Telemetry = telemetry
	}
}

// WithLogger sets the logging utility used by this object
func WithLogger(logger *zap.Logger) Option {
	return func(c *Client) {
		c.Logger = logger
	}
}

// WithCollectionNames sets the collection names
func WithCollectionNames(collectionNames []string) Option {
	return func(c *Client) {
		c.CollectionNames = collectionNames
	}
}

// WithClientOptions sets the client options
func WithClientOptions(clientOptions *options.ClientOptions) Option {
	return func(c *Client) {
		c.ClientOptions = clientOptions
	}
}

// WithConnectionURI sets the connection URI
func WithConnectionURI(connectionURI string) Option {
	return func(c *Client) {
		c.ConnectionURI = &connectionURI
	}
}

// Validate validates the client
func (c *Client) Validate() error {
	if c.MaxConnectionAttempts <= 0 {
		c.MaxConnectionAttempts = 3
	}

	if c.MaxRetriesPerOperation <= 0 {
		c.MaxRetriesPerOperation = 3
	}

	if c.RetryTimeOut <= 0 {
		c.RetryTimeOut = 30 * time.Second
	}

	if c.OperationSleepInterval <= 0 {
		c.OperationSleepInterval = 1 * time.Second
	}

	if c.QueryTimeout <= 0 {
		c.QueryTimeout = 2 * time.Second
	}

	if c.DatabaseName == nil {
		return fmt.Errorf("database name not set")
	}

	if c.Telemetry == nil {
		return fmt.Errorf("telemetry not set")
	}

	if c.Logger == nil {
		return fmt.Errorf("logger not set")
	}

	if len(c.CollectionNames) == 0 {
		return fmt.Errorf("collection names not set")
	}

	if c.ClientOptions == nil {
		return fmt.Errorf("client options not set")
	}

	if c.ConnectionURI == nil {
		return fmt.Errorf("connection URI not set")
	}

	return nil
}
