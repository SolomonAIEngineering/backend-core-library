package instrumentation

import (
	"testing"

	"go.uber.org/zap"
)

func TestWithServiceName(t *testing.T) {
	c := &Client{}

	// Use the WithServiceName option to set the service name
	WithServiceName("my-service")(c)

	// Check that the service name was set correctly
	if c.ServiceName != "my-service" {
		t.Errorf("Service name was not set correctly, expected 'my-service' but got '%s'", c.ServiceName)
	}
}

func TestWithServiceVersion(t *testing.T) {
	c := &Client{}

	WithServiceVersion("1.2.3")(c)

	if c.ServiceVersion != "1.2.3" {
		t.Errorf("Service version was not set correctly, expected '1.2.3' but got '%s'", c.ServiceVersion)
	}
}

func TestWithServiceEnvironment(t *testing.T) {
	c := &Client{}

	WithServiceEnvironment("production")(c)

	if c.ServiceEnvironment != "production" {
		t.Errorf("Service environment was not set correctly, expected 'production' but got '%s'", c.ServiceEnvironment)
	}
}

func TestWithEnabled(t *testing.T) {
	c := &Client{}

	WithEnabled(true)(c)

	if !c.Enabled {
		t.Errorf("Enabled was not set correctly, expected 'true' but got 'false'")
	}
}

func TestWithNewrelicKey(t *testing.T) {
	c := &Client{}

	WithNewrelicKey("abc123")(c)

	if c.NewrelicKey != "abc123" {
		t.Errorf("New Relic key was not set correctly, expected 'abc123' but got '%s'", c.NewrelicKey)
	}
}

func TestWithLogger(t *testing.T) {
	c := &Client{}

	logger := zap.NewExample()
	WithLogger(logger)(c)

	if c.Logger != logger {
		t.Errorf("Logger was not set correctly")
	}
}

func TestWithEnableMetrics(t *testing.T) {
	c := &Client{}

	WithEnableMetrics(true)(c)

	if !c.EnableMetrics {
		t.Errorf("EnableMetrics was not set correctly, expected 'true' but got 'false'")
	}
}

func TestWithEnableTracing(t *testing.T) {
	c := &Client{}

	WithEnableTracing(true)(c)

	if !c.EnableTracing {
		t.Errorf("EnableTracing was not set correctly, expected 'true' but got 'false'")
	}
}

func TestWithEnableEvents(t *testing.T) {
	c := &Client{}

	WithEnableEvents(true)(c)

	if !c.EnableEvents {
		t.Errorf("EnableEvents was not set correctly, expected 'true' but got 'false'")
	}
}