package cacher

import (
	"testing"

	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

func TestClient_Validate(t *testing.T) {
	// Test case 1: Redis connection is nil
	c := &Client{}
	err := c.Validate()
	if err == nil || err.Error() != "redis connection is nil" {
		t.Errorf("Expected error: 'redis connection is nil', got: %v", err)
	}

	// Test case 2: Logger is nil
	c.pool = redis.NewPool(func() (redis.Conn, error) {
		return nil, nil
	}, 1)

	err = c.Validate()
	if err == nil || err.Error() != "logger is nil" {
		t.Errorf("Expected error: 'logger is nil', got: %v", err)
	}

	// Test case 3: Service name is empty
	c.logger = zap.L()
	err = c.Validate()
	if err == nil || err.Error() != "service name is empty" {
		t.Errorf("Expected error: 'service name is empty', got: %v", err)
	}

	// Test case 4: Cache TTL is 0
	c.serviceName = "my-service"
	err = c.Validate()
	if err == nil || err.Error() != "cache TTL is 0" {
		t.Errorf("Expected error: 'cache TTL is 0', got: %v", err)
	}

	// Test case 5: Instrumentation client is nil (warning logged)
	c.cacheTTLInSeconds = 60
	err = c.Validate()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}
