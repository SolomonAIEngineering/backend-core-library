package clickhouse // import "github.com/SimifiniiCTO/simfiny-core-lib/database/clickhouse"

import (
	"errors"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/giantswarm/retry-go"
	"go.uber.org/zap"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

// `type Client struct {` is defining a new struct type named `Client`. This struct will have several
// fields that will be used to store various properties and settings related to a ClickHouse database
// connection. The `Client` struct will be used to create instances of a ClickHouse client that can be
// used to execute queries against the database.
type Client struct {
	// `Engine *gorm.DB` is defining a field named `Engine` of type `*gorm.DB` in the `Client` struct. This
	// field is used to store a reference to a GORM database connection.
	Engine *gorm.DB

	// `QueryTimeout` is a field in the `Client` struct that stores a pointer to a `time.Duration` value.
	// This value represents the maximum amount of time that a query can take before timing out. It is used
	// to set a timeout for queries executed by the `Client` instance.
	QueryTimeout *time.Duration

	// `maxConnectionRetries *int` is a field in the `Client` struct that stores a pointer to an `int`
	// value. This value represents the maximum number of times the `Client` instance will attempt to
	// establish a connection to the ClickHouse database in case of connection failure. If the connection
	// fails, the `Client` will retry the connection `maxConnectionRetries` number of times before giving
	// up. The pointer is used to allow the value to be nil, indicating that no retries should be
	// attempted.
	maxConnectionRetries *int

	// `maxConnectionRetry *time.Duration` is a field in the `Client` struct that stores a pointer to a
	// `time.Duration` value. This value represents the maximum amount of time that the `Client` instance
	// will wait before retrying a failed connection attempt to the ClickHouse database. If a connection
	// attempt fails, the `Client` will wait for `maxConnectionRetry` duration before attempting to
	// establish the connection again. The pointer is used to allow the value to be nil, indicating that no
	// retry delay should be applied.
	maxConnectionRetryTimeout *time.Duration

	// `retrySleep *time.Duration` is a field in the `Client` struct that stores a pointer to a
	// `time.Duration` value. This value represents the amount of time that the `Client` instance will wait
	// before retrying a failed query to the ClickHouse database. If a query fails, the `Client` will wait
	// for `retrySleep` duration before attempting to execute the query again. The pointer is used to allow
	// the value to be nil, indicating that no retry delay should be applied.
	retrySleep *time.Duration

	// `connectionURI *string` is a field in the `Client` struct that stores a pointer to a string value.
	// This string value represents the URI for the ClickHouse database connection. The pointer is used to
	// allow the value to be nil, indicating that no connection URI has been set.
	connectionURI *string

	// `maxIdleConnections` is a field in the `Client` struct that stores an integer value representing the
	// maximum number of idle connections in the connection pool. This value determines the maximum number
	// of connections that can be kept open and idle in the pool, waiting to be used for executing queries.
	// When the number of idle connections in the pool exceeds this value, the excess connections are
	// closed and removed from the pool. This helps to prevent resource wastage and optimize the use of
	// available connections.
	maxIdleConnections *int

	// `MaxOpenConnections` is a field of the `Client` struct that holds a pointer to an integer value
	// representing the maximum number of open connections in the connection pool. It is used to set the
	// maximum number of open connections that can be kept in the pool and is passed to the
	// `SetMaxOpenConns` function of the `sql.DB` object obtained from the `gorm.DB` object. This helps to
	// optimize the performance of the database connection by limiting the number of open connections that
	// are kept open.
	maxOpenConnections *int
	// `MaxConnectionLifetime` is a field of the `Client` struct that holds a pointer to a `time.Duration`
	// value representing the maximum amount of time a connection can remain open before it is closed and
	// removed from the connection pool. It is used to set the maximum lifetime of a connection and is
	// passed to the `SetConnMaxLifetime` function of the `sql.DB` object obtained from the `gorm.DB`
	// object. This helps to optimize the performance of the database connection by limiting the amount of
	// time a connection can remain open and reducing the risk of resource leaks or other issues that can
	// arise from long-lived connections.
	maxConnectionLifetime *time.Duration

	// `InstrumentationClient *instrumentation.Client` is defining a field in the `Client` struct that
	// stores a reference to an `instrumentation.Client` instance. This instance is used to collect and
	// report metrics and traces related to the ClickHouse database queries executed by the `Client`
	// instance. The `InstrumentationClient` field allows the `Client` instance to integrate with an
	// external instrumentation library for monitoring and analysis purposes.
	InstrumentationClient *instrumentation.Client

	// `Logger *zap.Logger` is defining a field in the `Client` struct that stores a reference to a
	// `zap.Logger` instance. This instance is used for logging purposes and allows the `Client` instance
	// to log messages related to the ClickHouse database queries executed by the `Client`. The `Logger`
	// field allows the `Client` instance to integrate with an external logging library for debugging and
	// analysis purposes.
	Logger *zap.Logger
}

// The New function creates a new client with optional configuration options.
func New(options ...Option) (*Client, error) {
	c := &Client{}

	for _, option := range options {
		option(c)
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	if err := c.connect(); err != nil {
		return nil, err
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := c.Engine.DB()
	if err != nil {
		panic("failed to obtain generic db connection")
	}

	sqlDB.SetMaxIdleConns(*c.maxIdleConnections)
	sqlDB.SetMaxOpenConns(*c.maxOpenConnections)
	sqlDB.SetConnMaxLifetime(*c.maxConnectionLifetime)

	// configure gorm
	c.Engine.FullSaveAssociations = true
	c.Engine.SkipDefaultTransaction = false
	c.Engine.PrepareStmt = true
	c.Engine.DisableAutomaticPing = false
	c.Engine = c.Engine.Set("gorm:auto_preload", true)

	// ping the database
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return c, nil
}

// Close closes the database connection
func (c *Client) Close() error {
	clickhouseDb, err := c.Engine.DB()
	if err != nil {
		return err
	}
	return clickhouseDb.Close()
}

// connect attempts to connect to the database using retries
func (c *Client) connect() error {
	var connection = make(chan *gorm.DB, 1)

	err := retry.Do(
		func(conn chan<- *gorm.DB) func() error {
			return func() error {
				newConn, err := gorm.Open(clickhouse.Open(*c.connectionURI), &gorm.Config{})
				conn <- newConn
				return err
			}
		}(connection),
		retry.MaxTries(*c.maxConnectionRetries),
		retry.Timeout(*c.maxConnectionRetryTimeout),
		retry.Sleep(*c.retrySleep),
	)

	if err != nil {
		return errors.New("exceeded retries")
	}

	c.Engine = <-connection
	return nil
}
