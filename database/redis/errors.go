package redis // import "github.com/SimifiniiCTO/simfiny-core-lib/database/redis"

import "errors"

var (
	ErrInvalidURI               = errors.New("invalid redis URI")
	ErrInvalidServiceName       = errors.New("invalid service name")
	ErrInvalidLogger            = errors.New("invalid logger")
	ErrInvalidTelemetrySdk      = errors.New("invalid telemetry SDK")
	ErrInvalidPool              = errors.New("invalid pool")
	ErrInvalidCacheTTLInSeconds = errors.New("invalid cache TTL in seconds")
)
