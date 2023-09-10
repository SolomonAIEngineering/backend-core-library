package consumer

import (
	"context"
	"testing"

	client "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"
)

func TestConsumerClient_StartConcurentConsumer(t *testing.T) {
	type args struct {
		f MessageProcessorFunc
	}
	tests := []struct {
		name string
		c    *ConsumerClient
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.StartConcurentConsumer(tt.args.f)
		})
	}
}

func TestConsumerClient_processMessageConcurrently(t *testing.T) {
	type args struct {
		ctx     context.Context
		message *client.Message
		f       MessageProcessorFunc
		sync    chan bool
	}
	tests := []struct {
		name string
		c    *ConsumerClient
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.processMessageConcurrently(tt.args.ctx, tt.args.message, tt.args.f, tt.args.sync)
		})
	}
}
