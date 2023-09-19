package redis

import (
	"os"
	"testing"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestMain(m *testing.M) {
	// Run all tests in the package.
	result := m.Run()

	// Exit with the same status code as the tests.
	os.Exit(result)
}

func TestClient_Close(t *testing.T) {
	redisTestServer := miniredis.RunT(t)
	defer redisTestServer.Close()

	// Set up test stop channel
	stopCh := make(chan struct{})

	// Set up test options
	logger := zap.NewNop()
	opts := []Option{
		WithLogger(logger),
		WithURI(redisTestServer.Addr()),
		WithCacheTTLInSeconds(60),
		WithServiceName("test-service"),
		WithTelemetrySdk(&instrumentation.Client{}),
	}

	// Call New() to create a new Redis client
	client, err := New(stopCh, opts...)
	assert.NoError(t, err)

	// Ensure pool is not nil
	assert.NotNil(t, client.pool)

	// Close the pool
	client.Close()

	// Ensure trying to close the pool again results in a nil value
	assert.True(t, client.pool.ActiveCount() == 0)
}

func TestNew(t *testing.T) {
	redisTestServer := miniredis.RunT(t)
	defer redisTestServer.Close()

	// Set up test stop channel
	stopCh := make(chan struct{})

	// Set up test options
	logger := zap.NewNop()
	opts := []Option{
		WithLogger(logger),
		WithURI(redisTestServer.Addr()),
		WithCacheTTLInSeconds(60),
		WithServiceName("test-service"),
		WithTelemetrySdk(&instrumentation.Client{}),
	}

	// Call New() to create a new Redis client
	client, err := New(stopCh, opts...)
	assert.NoError(t, err)

	// Ensure client is not nil
	assert.NotNil(t, client)

	// Ensure pool is not nil
	assert.NotNil(t, client.pool)

	// Ensure telemetrySdk is not nil
	assert.NotNil(t, client.telemetrySdk)

	// Ensure serviceName is set correctly
	assert.Equal(t, "test-service", client.serviceName)

	// Ensure cacheTTLInSeconds is set correctly
	assert.Equal(t, 60, client.cacheTTLInSeconds)

	// Ensure URI is set correctly
	assert.Equal(t, redisTestServer.Addr(), client.URI)

	// Ensure logger is set correctly
	assert.Equal(t, logger, client.Logger)
}
