package temporalclient // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client"

import (
	"go.uber.org/zap"
)

// TemporalConfig holds the configuration necessary for setting up a Temporal client.
// This includes the paths to the TLS certificate and key, the Temporal namespace,
// the host and port information, and a flag to enable mTLS.
type TemporalConfig struct {
	ClientKeyPath  string      // Path to the TLS key file
	ClientCertPath string      // Path to the TLS certificate file
	Namespace      string      // Temporal namespace
	HostPort       string      // Host and port of the Temporal service
	MtlsEnabled    bool        // Flag to enable mTLS
	Logger         *zap.Logger // Logger
}

// TemporalOption defines a type for functional options.
type TemporalOption func(*TemporalConfig)

// WithCertPath sets the client certificate path.
func WithCertPath(certPath string) TemporalOption {
	return func(tc *TemporalConfig) {
		tc.ClientCertPath = certPath
	}
}

// WithKeyPath sets the client key path.
func WithKeyPath(keyPath string) TemporalOption {
	return func(tc *TemporalConfig) {
		tc.ClientKeyPath = keyPath
	}
}

// WithNamespace sets the Temporal namespace.
func WithNamespace(namespace string) TemporalOption {
	return func(tc *TemporalConfig) {
		tc.Namespace = namespace
	}
}

// WithPort sets the Temporal port.
func WithHostPort(hostPort string) TemporalOption {
	return func(tc *TemporalConfig) {
		tc.HostPort = hostPort
	}
}

func WithMtlsEnabled(mtlsEnabled bool) TemporalOption {
	return func(tc *TemporalConfig) {
		tc.MtlsEnabled = mtlsEnabled
	}
}

func WithLogger(logger *zap.Logger) TemporalOption {
	return func(tc *TemporalConfig) {
		tc.Logger = logger
	}
}

// validate the configuration
func (tc *TemporalConfig) validate() error {
	if tc.MtlsEnabled {
		if tc.ClientCertPath == "" {
			return ErrCertPathNotSet
		}

		if tc.ClientKeyPath == "" {
			return ErrKeyPathNotSet
		}
	}

	if tc.Namespace == "" {
		return ErrNamespaceNotSet
	}

	if tc.HostPort == "" {
		return ErrHostPortNotSet
	}

	if tc.Logger == nil {
		return ErrLoggerNotSet
	}

	return nil
}

// NewTemporalConfig creates a new instance of TemporalConfig with the provided parameters.
// This function is a convenient way to create a configuration with all necessary details.
//
// Example:
//
//	config := NewTemporalConfig("./ca.key", "./ca.pem", "example_namespace", "localhost:7233", true)
func NewTemporalConfig(opts ...TemporalOption) (*TemporalConfig, error) {
	config := &TemporalConfig{}
	for _, opt := range opts {
		opt(config)
	}

	// validate the configuration
	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}
