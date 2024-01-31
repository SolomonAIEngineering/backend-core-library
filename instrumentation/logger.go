package instrumentation // import "github.com/SolomonAIEngineering/backend-core-library/instrumentation"

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
)

type contextKey string

const (
	userIDKey    contextKey = "userID"
	deviceIDKey  contextKey = "deviceID"
	requestIDKey contextKey = "requestID"
)

func (c contextKey) String() string {
	return string(c)
}

const (
	traceIDKey    = "trace.id"
	spanIDKey     = "span.id"
	entityGUIDKey = "entity.guid"
	entityNameKey = "entity.name"
	entityTypeKey = "entity.type"
	hostnameKey   = "hostname"
)

func (s *Client) LoggerWithCtxValue(ctx context.Context) *zap.Logger {
	return s.Logger.With(keyAndValueFromContext(ctx)...)
}

func keyAndValueFromContext(ctx context.Context) []zap.Field {
	metadata := newrelic.FromContext(ctx).GetLinkingMetadata()

	return []zap.Field{
		zap.String(userIDKey.String(), getUserID(ctx)),
		zap.String(deviceIDKey.String(), getDeviceID(ctx)),
		zap.String(requestIDKey.String(), getRequestID(ctx)),
		zap.String(traceIDKey, metadata.TraceID),
		zap.String(spanIDKey, metadata.SpanID),
		zap.String(entityGUIDKey, metadata.EntityGUID),
		zap.String(entityNameKey, metadata.EntityName),
		zap.String(entityTypeKey, metadata.EntityType),
		zap.String(hostnameKey, metadata.Hostname),
	}
}

// setUserID stores the user ID in the context
func setUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// getUserID retrieves the user ID from the context
func getUserID(ctx context.Context) string {
	val := ctx.Value(userIDKey)
	if val == nil {
		return ""
	}
	return val.(string)
}

// setDeviceID stores the device ID in the context
func setDeviceID(ctx context.Context, deviceID string) context.Context {
	return context.WithValue(ctx, deviceIDKey, deviceID)
}

// getDeviceID retrieves the device ID from the context
func getDeviceID(ctx context.Context) string {
	val := ctx.Value(deviceIDKey)
	if val == nil {
		return ""
	}
	return val.(string)
}

// setRequestID stores the request ID in the context
func setRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// getRequestID retrieves the request ID from the context
func getRequestID(ctx context.Context) string {
	val := ctx.Value(requestIDKey)
	if val == nil {
		return ""
	}
	return val.(string)
}
