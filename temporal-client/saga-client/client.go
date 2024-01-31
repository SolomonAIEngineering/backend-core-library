package sagaclient

import (
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

type SagaClient struct {
	Client          client.Client
	RegisteredSagas map[string]WorkflowFunc
	Logger          *zap.Logger
}

func NewSagaClient(options ...SagaClientOption) (*SagaClient, error) {
	config := &SagaClient{}

	for _, option := range options {
		option(config)
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}
