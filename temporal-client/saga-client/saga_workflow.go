package sagaclient

import (
	"go.temporal.io/sdk/workflow"
)

// NewSagaWorkflow creates a new SagaWorkflow.
func NewSagaWorkflow(options ...SagaWorkflowOption) (*SagaWorkflow, error) {
	sw := &SagaWorkflow{}

	for _, option := range options {
		option(sw)
	}

	// validate the SagaWorkflow
	if err := sw.validate(); err != nil {
		return nil, err
	}

	return sw, nil
}

// validate validates the SagaWorkflow.
func (s *SagaWorkflow) validate() error {
	if s.ActionActivity == nil {
		return ErrActionActivityNotSet
	}

	if s.CompensationActivity == nil {
		return ErrCompensationActivityNotSet
	}

	if s.Options == nil {
		return ErrActivityOptionsNotSet
	}

	if s.Identifier == "" {
		return ErrIdentifierNotSet
	}

	return nil
}

func (s *SagaWorkflow) toWorkflowFunc() WorkflowFunc {
	return func(ctx workflow.Context, input ...any) (err error) {
		activtyOptions := s.Options
		ctx = workflow.WithActivityOptions(ctx, *activtyOptions)

		// execute the action activity
		err = workflow.ExecuteActivity(ctx, s.ActionActivity, input).Get(ctx, nil)
		if err != nil {
			return err
		}

		// execute the compensation activity
		err = workflow.ExecuteActivity(ctx, s.CompensationActivity, input).Get(ctx, nil)
		if err != nil {
			return err
		}

		return nil
	}
}
