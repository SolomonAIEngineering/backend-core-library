package consumer

import (
	"context"
	"sync"
	"testing"

	client "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"
)

func TestConsumerClient_NaiveConsumer(t *testing.T) {
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
			tt.c.NaiveConsumer(tt.args.f)
		})
	}
}

func TestConsumerClient_processMessage(t *testing.T) {
	type args struct {
		ctx     context.Context
		message *client.Message
		f       MessageProcessorFunc
		wg      *sync.WaitGroup
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
			tt.c.processMessage(tt.args.ctx, tt.args.message, tt.args.f, tt.args.wg)
		})
	}
}
