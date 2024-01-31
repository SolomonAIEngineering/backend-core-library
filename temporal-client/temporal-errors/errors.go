package temporalerrors // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client/temporal-errors"

import "errors"

var (
	// ErrInvalidInput is returned when the input provided to a workflow is invalid.
	ErrInvalidInput                   = errors.New("invalid input")
	ErrNilWorkerConfig                = errors.New("worker config is nil")
	ErrNilClient                      = errors.New("client is nil")
	ErrEmptyTaskQueue                 = errors.New("task queue is empty")
	ErrInvalidMaxConcurrentActivities = errors.New("max concurrent activities is invalid")
	ErrInvalidMaxConcurrentWorkflows  = errors.New("max concurrent workflows is invalid")
)
