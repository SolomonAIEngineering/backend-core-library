package sagaclient // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client/saga-client"

import "errors"

var (
	ErrConfigNotSet               = errors.New("config not set")
	ErrClientNotSet               = errors.New("client not set")
	ErrSagasNotSet                = errors.New("sagas not set")
	ErrLoggerNotSet               = errors.New("logger not set")
	ErrRetryPolicyNotSet          = errors.New("retry policy not set")
	ErrActionActivityNotSet       = errors.New("action activity not set")
	ErrCompensationActivityNotSet = errors.New("compensation activity not set")
	ErrActivityOptionsNotSet      = errors.New("activity options not set")
	ErrIdentifierNotSet           = errors.New("identifier not set")
	ErrSagaWorkflowNotFound       = errors.New("saga workflow not found")
	ErrStartWorkflowOptionsNotSet = errors.New("start workflow options not set")
	ErrInputNotSet                = errors.New("input not set")
)
