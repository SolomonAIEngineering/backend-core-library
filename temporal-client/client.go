// Package temporalclient provides functionality to create and initialize
// a client for Temporal, a distributed, scalable, durable, and highly available
// orchestration engine. This package supports configuration with optional
// mutual TLS (mTLS) for enhanced security.
//
// The main components are the TemporalConfig struct and the NewTemporalClient
// function, which uses the configuration to create a Temporal client.
package temporalclient // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client"

import (
	"crypto/tls"

	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

type TemporalExecutorClient struct {
	Client client.Client
	Logger *zap.Logger
}

// New creates and initializes a new Temporal client using the provided TemporalConfig.
// This function handles the setup of mTLS based on the configuration and returns an initialized client.
//
// Example:
//
//	config := NewTemporalConfig("./ca.key", "./ca.pem", "example_namespace", "localhost:7233", true)
//	client, err := NewTemporalClient(config)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer client.Close()
//
//	// Use the client...
func New(config *TemporalConfig) (*TemporalExecutorClient, error) {
	var tlsConfig *tls.Config

	// Load the certificate and key only if mTLS is enabled
	if config.MtlsEnabled {
		cert, err := tls.LoadX509KeyPair(config.ClientCertPath, config.ClientKeyPath)
		if err != nil {
			return nil, err
		}
		tlsConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
	}

	// Create the Temporal client
	temporalClient, err := client.Dial(client.Options{
		HostPort:  config.HostPort,
		Namespace: config.Namespace,
		ConnectionOptions: client.ConnectionOptions{
			TLS: tlsConfig,
		},
	})
	if err != nil {
		return nil, err
	}

	return &TemporalExecutorClient{
		Client: temporalClient,
		Logger: config.Logger,
	}, nil
}

// Close closes the Temporal client.
func (c *TemporalExecutorClient) Close() {
	c.Client.Close()
}
