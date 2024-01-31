package worker // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client/worker"

import "errors"

var (
	ErrNoActivities                   = errors.New("no activities provided")
	ErrNoWorkflows                    = errors.New("no workflows provided")
	ErrWorkerRegistrationConfigNotSet = errors.New("workflows and activities not set")
)
