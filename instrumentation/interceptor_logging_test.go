package instrumentation

import (
	"context"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestInterceptorLogger(t *testing.T) {
	// create a buffer for test logging output
	observedZapCore, observedLogs := observer.New(zap.DebugLevel)
	observedLogger := zap.New(observedZapCore)

	// create a logging.Logger using the InterceptorLogger function
	interceptorLogger := InterceptorLogger(observedLogger)

	// test logging at different levels
	interceptorLogger.Log(context.Background(), logging.LevelDebug, "debug message", "key", "value")
	interceptorLogger.Log(context.Background(), logging.LevelInfo, "info message", "key", "value")
	interceptorLogger.Log(context.Background(), logging.LevelWarn, "warn message", "key", "value")
	interceptorLogger.Log(context.Background(), logging.LevelError, "error message", "key", "value")

	// check that log output was captured by the logger
	assert.Equal(t, observedLogs.Len(), 4)
}
