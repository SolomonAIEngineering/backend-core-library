package worker

import "errors"

var (
	ErrNoActivities                   = errors.New("no activities provided")
	ErrNoWorkflows                    = errors.New("no workflows provided")
	ErrWorkerRegistrationConfigNotSet = errors.New("workflows and activities not set")
)
