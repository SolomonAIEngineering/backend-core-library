package instrumentation // import "github.com/SolomonAIEngineering/backend-core-library/instrumentation"

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	nrgorilla "github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// The `Client` type is a struct that holds information about a service, including its name, version,
// environment, telemetry settings, New Relic license key, and a logger object.
// @property {string} ServiceName - The name of the service that this client is associated with.
// @property {string} ServiceVersion - The version of the service.
// @property {string} ServiceEnvironment - The environment of the service, such as "production",
// "staging", or "development".
// @property {bool} Enabled - a boolean value indicating whether telemetry/metrics are enabled for the
// service.
// @property client - The `client` property is a pointer to a `newrelic.Application` object, which is
// used to interact with the New Relic API and send telemetry data.
// @property {string} NewrelicKey - `NewrelicKey` is a field in the `Client` struct that holds the New
// Relic license key. This key is used to authenticate and connect to the New Relic service for
// monitoring and reporting purposes.
// @property Logger - A pointer to a `zap.Logger` object, which is used for logging purposes and allows
// the `Client` object to log messages using the `zap` logging library. The `json:"-"` tag is used to
// indicate that this field should not be included when the struct is serialized to JSON.
type Client struct {
	// The name of the service
	ServiceName string `json:"service_name"`
	// The version of the service
	ServiceVersion string `json:"service_version"`
	// The environment of the service
	ServiceEnvironment string `json:"service_environment"`
	// wether telemetry is enabled
	Enabled bool `json:"metrics_enabled"`
	// wether metrics are enabled
	EnableMetrics bool `json:"enable_metrics"`
	// wether tracing is enabled
	EnableTracing bool `json:"enable_tracing"`
	// wether events are enabled
	EnableEvents bool `json:"enable_events"`
	// wether logs are enabled
	EnableLogs bool `json:"enable_logs"`
	// The New Relic Sdk Object
	client *newrelic.Application
	// `NewrelicKey` is a field in the `Client` struct that holds the New Relic license key. The
	// `json:"newrelic_key"` tag is used to specify the name of the field when it is serialized to JSON.
	NewrelicKey string `json:"newrelic_key"`
	// `Logger      *zap.Logger `json:"-"` is a field in the `Client` struct that holds a pointer
	// to a `zap.Logger` object. The `json:"-"` tag is used to indicate that this field should not be
	// included when the struct is serialized to JSON. This field is used for logging purposes and allows
	// the `Client` object to log messages using the `zap` logging library.
	Logger *zap.Logger `json:"-"`
	// base service metrics for the service using this library
	baseMetrics *serviceBaseMetrics
}

// The IServiceTelemetry interface defines methods for collecting telemetry data for a service,
// including tracing, segments, and roundtrippers.
// @property {string} GetServiceName - Returns the name of the service that is being monitored for
// telemetry data.
// @property {string} GetServiceVersion - This method returns the version of the service being
// monitored by the telemetry service.
// @property {string} GetServiceEnvironment - This method returns the environment in which the service
// is running, such as "production", "staging", or "development".
// @property {bool} IsEnabled - A boolean property that indicates whether the telemetry service is
// enabled or not.
// @property GetTraceFromContext - This method retrieves a newrelic.Transaction object from a given
// context.Context object, which can be used to trace a transaction across multiple services or
// components.
// @property GetTraceFromRequest - This method retrieves a newrelic.Transaction object from an HTTP
// request. This object can be used to instrument and trace the performance of the request and any
// subsequent operations.
// @property WithContext - A method that takes a context.Context and a newrelic.Transaction as input
// parameters and returns a new context.Context with the newrelic.Transaction added to it. This is
// useful for propagating the transaction across different function calls within the same request.
// @property WithRequest - WithRequest is a method that takes an HTTP request and a
// newrelic.Transaction object and returns a new HTTP request with the transaction information added to
// its context. This is useful for propagating tracing information across service boundaries.
// @property StartTransaction - StartTransaction is a method that starts a new transaction with the
// given name and returns a pointer to the newrelic.Transaction object. This method is used to
// instrument a new transaction in the code.
// @property NewChildSpan - Creates a new child span within a given transaction. This can be used to
// track a specific operation or function within a larger transaction. The parent parameter specifies
// the transaction to which the child span belongs.
// @property StartExternalSegment - StartExternalSegment is a method that starts a new external segment
// in a New Relic transaction. It takes in the transaction object and the HTTP request object as
// parameters and returns a newrelic.ExternalSegment object. This method is used to track external
// calls made by the service, such as API calls
// @property StartDatastoreSegment - StartDatastoreSegment is a method that starts a new segment for a
// datastore operation within a transaction. It takes in the transaction object and the name of the
// operation being performed as parameters and returns a newrelic.DatastoreSegment object. This segment
// can be used to record metrics and trace the performance of
// @property StartMessageQueueSegment - StartMessageQueueSegment is a method that starts a new segment
// for a message queue operation within a New Relic transaction. It takes in the transaction object and
// the name of the message queue as parameters and returns a newrelic.MessageProducerSegment object.
// This segment can be used to record metrics and trace
// @property StartSegment - StartSegment is a method that starts a new custom segment within a
// transaction. It takes in the transaction object and a name for the segment as parameters and returns
// a newrelic.Segment object. This method can be used to track specific parts of a transaction that may
// not be automatically instrumented by New
// @property NewRoundtripper - NewRoundtripper is a method that returns an http.RoundTripper that
// instruments outgoing HTTP requests with New Relic tracing. This allows for tracing of external
// service calls made by the service. The returned RoundTripper should be used in place of the default
// RoundTripper in the http.Client used to
// @property StartNosqlDatastoreSegment - This method starts a new segment for a NoSQL datastore
// operation in a transaction. It takes in the transaction object, the operation being performed, and
// the name of the collection being accessed. It returns a newrelic.DatastoreSegment object that can be
// used to add additional information and end the segment.
type IServiceTelemetry interface {
	GetServiceName() string
	GetServiceVersion() string
	GetServiceEnvironment() string
	RecordEvent(eventType string, params map[string]interface{})
	RecordMetric(metric string, metricValue float64)
	IsEnabled() bool
	GetTraceFromContext(ctx context.Context) *newrelic.Transaction
	GetTraceFromRequest(r *http.Request) *newrelic.Transaction
	WithContext(ctx context.Context, trace newrelic.Transaction) context.Context
	WithRequest(r *http.Request, trace newrelic.Transaction) *http.Request
	StartTransaction(name string) *newrelic.Transaction
	NewChildSpan(name string, parent newrelic.Transaction) *newrelic.Segment
	StartExternalSegment(txn *newrelic.Transaction, req *http.Request) *newrelic.ExternalSegment
	StartDatastoreSegment(txn *newrelic.Transaction, operation string) *newrelic.DatastoreSegment
	StartMessageQueueSegment(txn *newrelic.Transaction, queueName string) *newrelic.MessageProducerSegment
	StartSegment(txn *newrelic.Transaction, name string) *newrelic.Segment
	NewRoundtripper(txn *newrelic.Transaction) *http.RoundTripper
	StartNosqlDatastoreSegment(txn *newrelic.Transaction, operation, collectionName string) *newrelic.DatastoreSegment
	GetUnaryServerInterceptors() []grpc.UnaryServerInterceptor
	GetStreamServerInterceptors() []grpc.StreamServerInterceptor
	GetUnaryClientInterceptors() []grpc.UnaryClientInterceptor
	GetStreamClientInterceptors() []grpc.StreamClientInterceptor
	NewMuxRouter() *mux.Router
}

var _ IServiceTelemetry = &Client{}

// NewServiceTelemetry creates a new Client
func New(opts ...Option) (*Client, error) {
	telemetry := &Client{}

	for _, opt := range opts {
		opt(telemetry)
	}

	if err := telemetry.configureNrClient(); err != nil {
		return nil, err
	}

	if err := telemetry.Validate(); err != nil {
		return nil, err
	}

	// initialize base service metrics set
	baseMetrics, err := newServiceBaseMetrics(&telemetry.ServiceName)
	if err != nil {
		return nil, err
	}

	telemetry.baseMetrics = baseMetrics
	return telemetry, nil
}

// RecordMetric takes two parameters: `metric` of type `string` and `metricValue` of type `float64`. The purpose of
// this method is to record a metric with a given name and value.
func (s *Client) RecordMetric(metric string, metricValue float64) {
	s.client.RecordCustomMetric(metric, metricValue)
}

// RecordEvent takes two parameters: `eventType` of type `string` and `params` of type `map[string]interface{}`.
// The purpose of this method is to record a custom event with a given name and parameters.
func (s *Client) RecordEvent(eventType string, params map[string]interface{}) {
	s.client.RecordCustomEvent(eventType, params)
}

// Enabled implements IServiceTelemetry
func (s *Client) IsEnabled() bool {
	return s.Enabled
}

// GetServiceEnvironment implements IServiceTelemetry
func (s *Client) GetServiceEnvironment() string {
	return s.ServiceEnvironment
}

// GetServiceName implements IServiceTelemetry
func (s *Client) GetServiceName() string {
	return s.ServiceName
}

// GetServiceVersion implements IServiceTelemetry
func (s *Client) GetServiceVersion() string {
	return s.ServiceVersion
}

// GetTraceFromContext implements IServiceTelemetry
func (s *Client) GetTraceFromContext(ctx context.Context) *newrelic.Transaction {

	return newrelic.FromContext(ctx)

}

// GetTraceFromRequest implements IServiceTelemetry
func (s *Client) GetTraceFromRequest(r *http.Request) *newrelic.Transaction {

	return newrelic.FromContext(r.Context())

}

// GetTracingEnabled implements IServiceTelemetry
func (s *Client) GetTracingEnabled() bool {
	return s.Enabled
}

// NewChildSpan implements IServiceTelemetry
func (s *Client) NewChildSpan(name string, parent newrelic.Transaction) *newrelic.Segment {

	return parent.StartSegment(name)

}

// StartTransaction implements IServiceTelemetry
func (s *Client) StartTransaction(name string) *newrelic.Transaction {

	return s.client.StartTransaction(name)

}

// WithContext implements IServiceTelemetry
func (s *Client) WithContext(ctx context.Context, trace newrelic.Transaction) context.Context {

	return newrelic.NewContext(ctx, &trace)

}

// WithRequest implements IServiceTelemetry
func (s *Client) WithRequest(r *http.Request, trace newrelic.Transaction) *http.Request {

	return nil

}

// StartExternalSegment starts an external segment
func (s *Client) StartExternalSegment(txn *newrelic.Transaction, req *http.Request) *newrelic.ExternalSegment {
	return newrelic.StartExternalSegment(txn, req)

}

// NewRoundTripper allows you to instrument external calls without
// calling StartExternalSegment by modifying the http.Client's Transport
// field.  If the Transaction parameter is nil, the RoundTripper
// returned will look for a Transaction in the request's context (using
// FromContext). This is recommended because it allows you to reuse the
// same client for multiple transactions.
func (s *Client) NewRoundtripper(txn *newrelic.Transaction) *http.RoundTripper {
	client := &http.Client{}

	h := newrelic.NewRoundTripper(client.Transport)
	return &h

}

// StartDatastoreSegment starts a datastore segment
func (s *Client) StartDatastoreSegment(txn *newrelic.Transaction, operation string) *newrelic.DatastoreSegment {

	segment := newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		// Product, Collection, and Operation are the most important
		// fields to populate because they are used in the breakdown
		// metrics.
		Product:   newrelic.DatastorePostgres,
		Operation: operation,
	}

	return &segment

}

// StartRedisDatastoreSegment starts a redis datastore segment
func (s *Client) StartRedisDatastoreSegment(txn *newrelic.Transaction, operation string) *newrelic.DatastoreSegment {

	segment := newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product:   newrelic.DatastoreRedis,
		Operation: operation,
	}

	return &segment

}

// StartDatastoreSegment starts a datastore segment
func (s *Client) StartNosqlDatastoreSegment(txn *newrelic.Transaction, operation, collectionName string) *newrelic.DatastoreSegment {

	segment := newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		// Product, Collection, and Operation are the most important
		// fields to populate because they are used in the breakdown
		// metrics.
		Product:    newrelic.DatastoreMongoDB,
		Operation:  operation,
		Collection: collectionName,
	}

	return &segment

}

// StartMessageQueueSegment starts a message queue segment
func (s *Client) StartMessageQueueSegment(txn *newrelic.Transaction, queueName string) *newrelic.MessageProducerSegment {

	segment := newrelic.MessageProducerSegment{
		StartTime:       txn.StartSegmentNow(),
		Library:         "aws-sqs",
		DestinationType: newrelic.MessageQueue,
		DestinationName: queueName,
	}

	return &segment

}

// StartSegment implements IServiceTelemetry
func (s *Client) StartSegment(txn *newrelic.Transaction, name string) *newrelic.Segment {
	return txn.StartSegment(name)
}

// GetUnaryServerInterceptors implements IServiceTelemetry
func (s *Client) GetUnaryServerInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		nrgrpc.UnaryServerInterceptor(s.client),
	}
}

// GetStreamServerInterceptors implements IServiceTelemetry
func (s *Client) GetStreamServerInterceptors() []grpc.StreamServerInterceptor {
	return []grpc.StreamServerInterceptor{
		nrgrpc.StreamServerInterceptor(s.client),
	}
}

// GetUnaryClientInterceptors implements IServiceTelemetry
func (s *Client) GetUnaryClientInterceptors() []grpc.UnaryClientInterceptor {
	return []grpc.UnaryClientInterceptor{
		nrgrpc.UnaryClientInterceptor,
	}
}

// GetStreamClientInterceptors implements IServiceTelemetry
func (s *Client) GetStreamClientInterceptors() []grpc.StreamClientInterceptor {
	return []grpc.StreamClientInterceptor{
		nrgrpc.StreamClientInterceptor,
	}
}

// NewMuxRouter returns a new mux router
func (s *Client) NewMuxRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(nrgorilla.Middleware(s.client))
	return r
}

// configureNrClient configures the newrelic client
func (s *Client) configureNrClient() error {
	z, _ := zap.NewProduction()

	client, err := newrelic.NewApplication(
		// service name
		newrelic.ConfigAppName(s.ServiceName),
		// license
		newrelic.ConfigLicense(s.NewrelicKey),
		// enabel application log forwarding
		newrelic.ConfigAppLogForwardingEnabled(s.EnableLogs),
		// configure max log forwarding samples
		newrelic.ConfigAppLogForwardingMaxSamplesStored(1000),
		// custom insights samples stored
		newrelic.ConfigCustomInsightsEventsMaxSamplesStored(1000),
		// enable app logging
		newrelic.ConfigAppLogEnabled(s.EnableLogs),
		// enable distributed tracing
		newrelic.ConfigDistributedTracerEnabled(s.EnableTracing),
		// enable the agent
		newrelic.ConfigEnabled(s.Enabled),
		// configure the logger to be names to the service name and use the zap logger
		nrzap.ConfigLogger(z.Named(s.ServiceName)),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.RecordPanics = s.Enabled
			cfg.ErrorCollector.Enabled = s.Enabled
			cfg.TransactionEvents.Enabled = s.EnableTracing
			cfg.TransactionEvents.MaxSamplesStored = 1000
			cfg.Attributes.Enabled = s.Enabled
			cfg.TransactionTracer.Enabled = s.Enabled
			cfg.SpanEvents.Enabled = s.Enabled
			cfg.RuntimeSampler.Enabled = s.Enabled
			cfg.DistributedTracer.Enabled = s.Enabled
			cfg.AppName = s.ServiceName
			cfg.DatastoreTracer.InstanceReporting.Enabled = s.Enabled
			cfg.DatastoreTracer.QueryParameters.Enabled = s.Enabled
			cfg.DatastoreTracer.DatabaseNameReporting.Enabled = s.Enabled
			cfg.Labels = map[string]string{
				"Environment": s.ServiceEnvironment,
			}
		},
	)

	if err != nil {
		return err
	}

	s.client = client
	return nil
}
