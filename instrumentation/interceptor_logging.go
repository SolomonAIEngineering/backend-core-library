package instrumentation // import "github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
)

// `InterceptorLogger` is a function that returns a `logging.Logger` which is used as an
// interceptor for gRPC requests. The function takes a `zap.Logger` as input and returns a
// `logging.LoggerFunc`. The returned `logging.LoggerFunc` takes in a context, logging level,
// message, and fields as input and logs the message and fields using the `zap.Logger` with the
// appropriate logging level. It also adds a caller skip to the `zap.Logger` to skip the logging
// function call in the stack trace.
func InterceptorLogger(l *zap.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)
		for i := 0; i < len(fields); i += 2 {
			i := logging.Fields(fields).Iterator()
			if i.Next() {
				k, v := i.At()
				f = append(f, zap.Any(k, v))
			}
		}
		l = l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			l.Debug(msg)
		case logging.LevelInfo:
			l.Info(msg)
		case logging.LevelWarn:
			l.Warn(msg)
		case logging.LevelError:
			l.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
