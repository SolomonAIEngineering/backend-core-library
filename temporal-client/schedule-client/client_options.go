package scheduleclient

import (
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

type ScheduleOption func(*ScheduleClient)

func WithClient(client *client.Client) ScheduleOption {
	return func(sc *ScheduleClient) {
		sc.client = client
	}
}

func WithLogger(logger *zap.Logger) ScheduleOption {
	return func(sc *ScheduleClient) {
		sc.Logger = logger
	}
}

func WithScheduleClient(scheduleClient client.ScheduleClient) ScheduleOption {
	return func(sc *ScheduleClient) {
		sc.ScheduleClient = scheduleClient
	}
}
