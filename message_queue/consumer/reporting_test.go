package consumer

import "testing"

func TestConsumerClient_reportErrorEvent(t *testing.T) {
	type args struct {
		op  string
		err error
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
			tt.c.reportErrorEvent(tt.args.op, tt.args.err)
		})
	}
}

func TestConsumerClient_reportProcessedMessageCount(t *testing.T) {
	type args struct {
		op string
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
			tt.c.reportProcessedMessageCount(tt.args.op)
		})
	}
}
