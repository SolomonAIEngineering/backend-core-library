// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package postgres // import "github.com/SolomonAIEngineering/backend-core-library/database/postgres"

import (
	"errors"
	"time"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/giantswarm/retry-go"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Client is defining a new struct type called `Client` which will be used to create
// instances of a client for connecting to a PostgreSQL database. The struct contains various fields
// that can be set using options, such as the database engine, query timeout, connection string, and
// instrumentation client. The struct also has a `connect` method that attempts to connect to the
// database using retries.
type Client struct {
	// `Engine` is a field of the `Client` struct that holds a pointer to a `gorm.DB` object, which is a
	// database connection object provided by the GORM library. This field is used to store the connection
	// to the PostgreSQL database and is set when the `connect` method is called during the creation of a
	// new `Client` instance.
	Engine *gorm.DB
	// `QueryTimeout` is a field of the `Client` struct that holds a pointer to a `time.Duration` value
	// representing the maximum amount of time to wait for a query to complete before timing out. This
	// field can be set using options when creating a new `Client` instance and is used to set the query
	// timeout for the database connection. If a query takes longer than the specified timeout, an error
	// will be returned.
	QueryTimeout *time.Duration
	// `MaxConnectionRetries` is a field of the `Client` struct that holds a pointer to an integer value
	// representing the maximum number of times to retry connecting to the PostgreSQL database in case of a
	// connection failure. It is used in the `connect` method to set the maximum number of retries for the
	// retry mechanism.
	MaxConnectionRetries *int
	// `MaxConnectionRetryTimeout` is a field of the `Client` struct that holds a pointer to a
	// `time.Duration` value representing the maximum amount of time to wait for a successful connection to
	// the PostgreSQL database before giving up and returning an error. It is used in the `connect` method
	// to set the maximum timeout for the retry mechanism.
	MaxConnectionRetryTimeout *time.Duration
	// `RetrySleep` is a field of the `Client` struct that holds a pointer to a `time.Duration` value
	// representing the amount of time to wait between retries when attempting to connect to the PostgreSQL
	// database. It is used in the `connect` method to set the sleep time for the retry mechanism.
	RetrySleep *time.Duration
	// `ConnectionString` is a field of the `Client` struct that holds a pointer to a string value
	// representing the connection string used to connect to the PostgreSQL database. It is used in the
	// `connect` method to open a new connection to the database using the `gorm.Open` function provided by
	// the GORM library. The connection string typically includes information such as the database name,
	// host, port, username, and password required to establish a connection to the database.
	ConnectionString *string
	// `MaxIdleConnections` is a field of the `Client` struct that holds a pointer to an integer value
	// representing the maximum number of idle connections in the connection pool. It is used to set the
	// maximum number of idle connections that can be kept in the pool and is passed to the
	// `SetMaxIdleConns` function of the `sql.DB` object obtained from the `gorm.DB` object. This helps to
	// optimize the performance of the database connection by limiting the number of idle connections that
	// are kept open.
	MaxIdleConnections *int
	// `MaxOpenConnections` is a field of the `Client` struct that holds a pointer to an integer value
	// representing the maximum number of open connections in the connection pool. It is used to set the
	// maximum number of open connections that can be kept in the pool and is passed to the
	// `SetMaxOpenConns` function of the `sql.DB` object obtained from the `gorm.DB` object. This helps to
	// optimize the performance of the database connection by limiting the number of open connections that
	// are kept open.
	MaxOpenConnections *int
	// `MaxConnectionLifetime` is a field of the `Client` struct that holds a pointer to a `time.Duration`
	// value representing the maximum amount of time a connection can remain open before it is closed and
	// removed from the connection pool. It is used to set the maximum lifetime of a connection and is
	// passed to the `SetConnMaxLifetime` function of the `sql.DB` object obtained from the `gorm.DB`
	// object. This helps to optimize the performance of the database connection by limiting the amount of
	// time a connection can remain open and reducing the risk of resource leaks or other issues that can
	// arise from long-lived connections.
	MaxConnectionLifetime *time.Duration
	// `InstrumentationClient` is a field of the `Client` struct that holds a pointer to an instance of the
	// `instrumentation.Client` struct. This field is used to pass an instrumentation client to the
	// `Client` instance, which can be used to collect metrics and traces related to the database
	// operations performed by the client. The `instrumentation.Client` struct provides methods for
	// collecting metrics and traces, which can be used to monitor the performance and behavior of the
	// database connection.
	InstrumentationClient *instrumentation.Client
	// `Logger *zap.Logger` is a field of the `Client` struct that holds a pointer to an instance of the
	// `zap.Logger` struct. This field is used to pass a logger to the `Client` instance, which can be used
	// to log messages related to the database operations performed by the client. The `zap.Logger` struct
	// provides methods for logging messages at different levels of severity, which can be used to monitor
	// the behavior of the database connection and diagnose issues that may arise.
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
		return nil, err
	}

	sqlDB.SetMaxIdleConns(*c.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(*c.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(*c.MaxConnectionLifetime)

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
	sqlDB, err := c.Engine.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// connect attempts to connect to the database using retries
func (c *Client) connect() error {
	var connection = make(chan *gorm.DB, 1)

	err := retry.Do(
		func(conn chan<- *gorm.DB) func() error {
			return func() error {
				newConn, err := gorm.Open(postgres.Open(*c.ConnectionString), &gorm.Config{
					CreateBatchSize: 500,
				})
				conn <- newConn
				return err
			}
		}(connection),
		retry.MaxTries(*c.MaxConnectionRetries),
		retry.Timeout(*c.MaxConnectionRetryTimeout),
		retry.Sleep(*c.RetrySleep),
	)

	if err != nil {
		return errors.New("exceeded retries")
	}

	c.Engine = <-connection
	return nil
}
