package signals // import "github.com/SolomonAIEngineering/backend-core-library/signals"

import (
	"context"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Shutdown struct {
	logger                *zap.Logger
	pool                  *redis.Pool
	serverShutdownTimeout time.Duration
}

func NewShutdown(serverShutdownTimeout time.Duration, logger *zap.Logger) (*Shutdown, error) {
	srv := &Shutdown{
		logger:                logger,
		serverShutdownTimeout: serverShutdownTimeout,
	}

	return srv, nil
}

func (s *Shutdown) Graceful(stopCh <-chan struct{}, httpServer *http.Server, httpsServer *http.Server, grpcServer *grpc.Server, healthy *int32, ready *int32) {
	ctx := context.Background()

	// wait for SIGTERM or SIGINT
	<-stopCh
	ctx, cancel := context.WithTimeout(ctx, s.serverShutdownTimeout)
	defer cancel()

	// all calls to /healthz and /readyz will fail from now on
	atomic.StoreInt32(healthy, 0)
	atomic.StoreInt32(ready, 0)

	// close cache pool
	if s.pool != nil {
		_ = s.pool.Close()
	}

	s.logger.Info("Shutting down HTTP/HTTPS server", zap.Duration("timeout", s.serverShutdownTimeout))

	// wait for Kubernetes readiness probe to remove this instance from the load balancer
	// the readiness check interval must be lower than the timeout
	if viper.GetString("level") != "debug" {
		time.Sleep(3 * time.Second)
	}

	// determine if the GRPC was started
	if grpcServer != nil {
		s.logger.Info("Shutting down GRPC server", zap.Duration("timeout", s.serverShutdownTimeout))
		grpcServer.GracefulStop()
	}

	// determine if the http server was started
	if httpServer != nil {
		if err := httpServer.Shutdown(ctx); err != nil {
			s.logger.Warn("HTTP server graceful shutdown failed", zap.Error(err))
		}
	}

	// determine if the secure server was started
	if httpsServer != nil {
		if err := httpsServer.Shutdown(ctx); err != nil {
			s.logger.Warn("HTTPS server graceful shutdown failed", zap.Error(err))
		}
	}
}
