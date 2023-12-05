package instrumentation // import "github.com/SolomonAIEngineering/backend-core-library/instrumentation"

import (
	"fmt"

	"go.uber.org/zap"
)

// Option is a function that configures a Client
type Option func(*Client)

// WithServiceName configures the service name
func WithServiceName(name string) Option {
	return func(t *Client) {
		t.ServiceName = name
	}
}

// WithServiceVersion configures the service version
func WithServiceVersion(version string) Option {
	return func(t *Client) {
		t.ServiceVersion = version
	}
}

// WithServiceEnvironment configures the service environment
func WithServiceEnvironment(environment string) Option {
	return func(t *Client) {
		t.ServiceEnvironment = environment
	}
}

// WithEnabled configures wether or not instrumentation is enabled
func WithEnabled(enabled bool) Option {
	return func(t *Client) {
		t.Enabled = enabled
	}
}

// WithNewrelicKey configures the newrelic key
func WithNewrelicKey(key string) Option {
	return func(t *Client) {
		t.NewrelicKey = key
	}
}

// WithLogger configures the logger
func WithLogger(logger *zap.Logger) Option {
	return func(t *Client) {
		t.Logger = logger
	}
}

// WithEnableMetrics configures wether or not metrics are enabled
func WithEnableMetrics(enableMetrics bool) Option {
	return func(t *Client) {
		t.EnableMetrics = enableMetrics
	}
}

// WithEnableTracing configures wether or not tracing is enabled
func WithEnableTracing(enableTracing bool) Option {
	return func(t *Client) {
		t.EnableTracing = enableTracing
	}
}

// WithEnableEvents configures wether or not events are enabled
func WithEnableEvents(enableEvents bool) Option {
	return func(t *Client) {
		t.EnableEvents = enableEvents
	}
}

// WithEnableLogger configures wether or not logging is enabled
func WithEnableLogger(enableLogger bool) Option {
	return func(t *Client) {
		t.EnableLogs = enableLogger
	}
}


func (t *Client) Validate() error {
	if t.ServiceName == "" {
		return fmt.Errorf("service name not set")
	}

	if t.ServiceVersion == "" {
		return fmt.Errorf("service version not set")
	}

	if t.ServiceEnvironment == "" {
		return fmt.Errorf("service environment not set")
	}

	if t.NewrelicKey == "" {
		return fmt.Errorf("newrelic key not set")
	}

	if t.Logger == nil {
		return fmt.Errorf("logger not set")
	}

	if t.client == nil {
		return fmt.Errorf("client not set")
	}

	return nil
}
