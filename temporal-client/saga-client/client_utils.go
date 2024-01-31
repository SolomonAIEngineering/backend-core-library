package sagaclient

import (
	"context"

	"go.temporal.io/sdk/client"
)

type ExecuteSagaResult struct {
	WorkflowID string
	RunID      string
}

type ExecuteSagaParams struct {
	Identifier string
	Options    *client.StartWorkflowOptions
	Input      []interface{}
}

func (p *ExecuteSagaParams) validate() error {
	if p.Identifier == "" {
		return ErrIdentifierNotSet
	}

	if p.Options == nil {
		return ErrStartWorkflowOptionsNotSet
	}

	if p.Input == nil || len(p.Input) == 0 {
		return ErrInputNotSet
	}

	return nil
}

func (c *SagaClient) Execute(ctx context.Context, param *ExecuteSagaParams) (*ExecuteSagaResult, error) {
	if err := param.validate(); err != nil {
		return nil, err
	}

	identifier := param.Identifier
	options := param.Options
	input := param.Input

	// get the workflow from the set of registered custom workflows
	wf, ok := c.RegisteredSagas[identifier]
	if !ok {
		return nil, ErrSagaWorkflowNotFound
	}

	we, err := c.Client.ExecuteWorkflow(ctx, *options, wf, input...)
	if err != nil {
		return nil, err
	}

	return &ExecuteSagaResult{
		WorkflowID: we.GetID(),
		RunID:      we.GetRunID(),
	}, nil
}
