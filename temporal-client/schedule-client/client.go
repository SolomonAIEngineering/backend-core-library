// Package scheduleclient provides functionalities to interact with a scheduling service,
// such as creating and managing scheduled tasks.
package scheduleclient

import (
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

// ScheduleClient is a struct that encapsulates clients and logger necessary for scheduling tasks.
// It includes the Temporal client, a logger, and a specific schedule client for task scheduling.
type ScheduleClient struct {
	client         *client.Client        // Temporal client for basic operations.
	Logger         *zap.Logger           // Logger for logging messages.
	ScheduleClient client.ScheduleClient // Specific client for scheduling-related operations.
}

// NewScheduleClient creates a new instance of ScheduleClient with the provided options.
// It applies each option to modify the ScheduleClient and validates it before returning.
// Returns an error if the validation fails.
func NewScheduleClient(opts ...ScheduleOption) (*ScheduleClient, error) {
	sc := &ScheduleClient{}

	// Applying provided options to the ScheduleClient instance.
	for _, opt := range opts {
		opt(sc)
	}

	// Validating the configured ScheduleClient.
	if err := sc.validate(); err != nil {
		return nil, err
	}

	return sc, nil
}

// validate checks the completeness and validity of the ScheduleClient configuration.
// It ensures that all necessary components (client, logger, and schedule client) are set.
// Returns an error if any component is missing.
func (sc *ScheduleClient) validate() error {
	if sc.client == nil {
		return ErrClientNotSet // Error when the Temporal client is not set.
	}

	if sc.Logger == nil {
		return ErrLoggerNotSet // Error when the logger is not set.
	}

	if sc.ScheduleClient == nil {
		return ErrScheduleClientNotSet // Error when the schedule client is not set.
	}

	return nil
}
