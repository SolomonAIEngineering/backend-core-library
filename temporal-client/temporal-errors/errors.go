package temporalerrors // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client/temporal-errors"

import "errors"

var (
	// ErrInvalidInput is returned when the input provided to a workflow is invalid.
	ErrInvalidInput = errors.New("invalid input")
	// ErrInvalidWorkflowID is returned when the workflow ID is invalid.
	ErrNilWorkerConfig = errors.New("worker config is nil")
	// ErrInvalidWorkflowID is returned when the workflow ID is invalid.
	ErrNilClient = errors.New("client is nil")
	// ErrInvalidWorkflowID is returned when the workflow ID is invalid.
	ErrEmptyTaskQueue = errors.New("task queue is empty")
	// ErrInvalidWorkflowID is returned when the workflow ID is invalid.
	ErrInvalidMaxConcurrentActivities = errors.New("max concurrent activities is invalid")
	// ErrInvalidWorkflowID is returned when the workflow ID is invalid.
	ErrInvalidMaxConcurrentWorkflows = errors.New("max concurrent workflows is invalid")
)
