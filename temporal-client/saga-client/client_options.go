package sagaclient // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client/saga-client"

import (
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

type SagaClientOption func(*SagaClient)

func WithClient(client client.Client) SagaClientOption {
	return func(sc *SagaClient) {
		sc.Client = client
	}
}

func WithLogger(logger *zap.Logger) SagaClientOption {
	return func(sc *SagaClient) {
		sc.Logger = logger
	}
}

func WithSagaCustomWorkflows(sagaWf ...SagaWorkflow) SagaClientOption {
	return func(sc *SagaClient) {
		for _, wf := range sagaWf {
			if sc.RegisteredSagas == nil {
				sc.RegisteredSagas = make(map[string]WorkflowFunc)
			}

			sc.RegisteredSagas[wf.Identifier] = wf.toWorkflowFunc()
		}
	}
}

func (sc *SagaClient) validate() error {
	if sc.Client == nil {
		return ErrClientNotSet
	}

	if sc.Logger == nil {
		return ErrLoggerNotSet
	}

	if sc.RegisteredSagas == nil {
		return ErrSagasNotSet
	}

	return nil
}
