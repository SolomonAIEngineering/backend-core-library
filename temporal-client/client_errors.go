package temporalclient // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client"

import "errors"

var (
	ErrCertPathNotSet  = errors.New("client certificate path not set")
	ErrKeyPathNotSet   = errors.New("client key path not set")
	ErrNamespaceNotSet = errors.New("namespace not set")
	ErrHostPortNotSet  = errors.New("host port not set")
	ErrLoggerNotSet    = errors.New("logger not set")
)
