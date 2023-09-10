package consumer

import (
	"reflect"
	"testing"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"
	"go.uber.org/zap"
)

func TestWithMessageProcessTimeout(t *testing.T) {
	type args struct {
		messageProcessTimeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMessageProcessTimeout(tt.args.messageProcessTimeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMessageProcessTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithQueuePollingDuration(t *testing.T) {
	type args struct {
		queuePollingDuration time.Duration
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithQueuePollingDuration(tt.args.queuePollingDuration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithQueuePollingDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithConcurrencyFactor(t *testing.T) {
	type args struct {
		concurrencyFactor int
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithConcurrencyFactor(tt.args.concurrencyFactor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithConcurrencyFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLogger(t *testing.T) {
	type args struct {
		logger *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLogger(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithQueueUrl(t *testing.T) {
	type args struct {
		queueUrl *string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithQueueUrl(tt.args.queueUrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithQueueUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithAwsClient(t *testing.T) {
	type args struct {
		sqsClient *client.Client
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithAwsClient(tt.args.sqsClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithAwsClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConsumerClient_Validate(t *testing.T) {
	tests := []struct {
		name    string
		c       *ConsumerClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ConsumerClient.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *ConsumerClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
