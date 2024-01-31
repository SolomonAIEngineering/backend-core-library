package scheduleclient // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client/schedule-client"

import "errors"

var (
	ErrClientNotSet         = errors.New("client not set")
	ErrLoggerNotSet         = errors.New("logger not set")
	ErrTaskIdNotSet         = errors.New("task id not set")
	ErrTaskSpecNotSet       = errors.New("task spec not set")
	ErrTaskActionNotSet     = errors.New("task action not set")
	ErrScheduleClientNotSet = errors.New("schedule client not set")
)
