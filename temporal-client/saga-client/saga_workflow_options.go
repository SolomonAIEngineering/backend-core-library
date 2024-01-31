package sagaclient // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client/saga-client"

import (
	"context"

	"go.temporal.io/sdk/workflow"
)

type SagaActivityFunc = func(ctx context.Context, input any) (err error)
type WorkflowFunc = func(workflow.Context, ...interface{}) (err error)

type SagaWorkflow struct {
	ActionActivity       SagaActivityFunc
	CompensationActivity SagaActivityFunc
	Options              *workflow.ActivityOptions
	Identifier           string
}

// SagaWorkflowOption defines the signature for functions that can configure SagaWorkflow
type SagaWorkflowOption func(*SagaWorkflow)

// WithActionActivity configures the action activity and its options.
func WithActionActivity(activityFunc SagaActivityFunc) SagaWorkflowOption {
	return func(sw *SagaWorkflow) {
		sw.ActionActivity = activityFunc
	}
}

// WithCompensationActivity configures the compensation activity and its options.
func WithCompensationActivity(activityFunc SagaActivityFunc) SagaWorkflowOption {
	return func(sw *SagaWorkflow) {
		sw.CompensationActivity = activityFunc
	}
}

// WithActivityOptions configures the activity options.
func WithActivityOptions(options workflow.ActivityOptions) SagaWorkflowOption {
	return func(sw *SagaWorkflow) {
		sw.Options = &options
	}
}

// WithIdentifier configures the identifier.
func WithIdentifier(identifier string) SagaWorkflowOption {
	return func(sw *SagaWorkflow) {
		sw.Identifier = identifier
	}
}
