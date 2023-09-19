package manager // import "github.com/SimifiniiCTO/simfiny-core-lib/third-party/transaction-manager"

import (
	"time"

	"go.temporal.io/sdk/client"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-core-lib/database/mongo"
	"github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	msq "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"
)

// Option is a function that takes a TransactionManager pointer and sets some option on it
type Option = func(*TransactionManager)

// WithClientOptions defines the options to use when connecting to the database
func WithClientOptions(option *client.Options) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.options = option
	}
}

// The function returns a function that takes a TransactionManager pointer and sets its instrumentation
// client to the provided one.
func WithInstrumentationClient(telemetry *instrumentation.Client) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.instrumentationClient = telemetry
	}
}

// The function returns a function that takes a TransactionManager pointer and sets its logger to the
// provided zap logger.
func WithLogger(logger *zap.Logger) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.logger = logger
	}
}

// The function returns a function that takes a TransactionManager pointer and sets its Postgres client
// field to the provided client.
func WithPostgres(postgres *postgres.Client) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.postgresClient = postgres
	}
}

// The function returns a closure that takes a TransactionManager pointer and sets its mongoConn field
// to the provided mongo.Client pointer.
func WithMongo(mongoConn *mongo.Client) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.mongoClient = mongoConn
	}
}

// The function returns a function that takes a TransactionManager and sets its retry policy to the
// provided policy.
func WithRetryPolicy(policy *Policy) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.retryPolicy = policy
	}
}

// The function returns a function that takes a TransactionManager pointer and sets its RPC timeout to
// the specified duration.
func WithRpcTimeout(timeout time.Duration) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.rpcTimeout = timeout
	}
}

// The function returns a function that takes a TransactionManager pointer and enables or disables
// metrics based on a boolean input.
func WithMetricsEnabled(enabled bool) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.metricsEnabled = enabled
	}
}

// This function returns a function that takes a TransactionManager pointer and sets its message queue
// client to the provided client.
func WithMessageQueueClient(client *msq.Client) func(*TransactionManager) {
	return func(t *TransactionManager) {
		t.messageQueueClient = client
	}
}
