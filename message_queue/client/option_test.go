package client

import (
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

func TestWithConfig(t *testing.T) {
	type args struct {
		config *Config
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
			if got := WithConfig(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithTimeout(t *testing.T) {
	type args struct {
		timeout time.Duration
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
			if got := WithTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithReadTimeout(t *testing.T) {
	type args struct {
		readTimeout time.Duration
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
			if got := WithReadTimeout(tt.args.readTimeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithReadTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithWriteTimeout(t *testing.T) {
	type args struct {
		writeTimeout time.Duration
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
			if got := WithWriteTimeout(tt.args.writeTimeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithWriteTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithSqsClient(t *testing.T) {
	type args struct {
		client sqsiface.SQSAPI
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
			if got := WithSqsClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithSqsClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithAwsConfig(t *testing.T) {
	type args struct {
		config *AwsConfig
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
			if got := WithAwsConfig(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithAwsConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Validate(t *testing.T) {
	tests := []struct {
		name    string
		c       *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWithMaxNumberOfMessages(t *testing.T) {
	type args struct {
		maxNumberOfMessages int
	}
	tests := []struct {
		name string
		args args
		want ConfigOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMaxNumberOfMessages(tt.args.maxNumberOfMessages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMaxNumberOfMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMaxWaitTimeSeconds(t *testing.T) {
	type args struct {
		MaxWaitTimeInSeconds time.Duration
	}
	tests := []struct {
		name string
		args args
		want ConfigOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMaxWaitTimeSeconds(tt.args.MaxWaitTimeInSeconds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMaxWaitTimeSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithAttributes(t *testing.T) {
	type args struct {
		attributes []string
	}
	tests := []struct {
		name string
		args args
		want ConfigOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithAttributes(tt.args.attributes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfigOptions(t *testing.T) {
	type args struct {
		opts []ConfigOption
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigOptions(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
