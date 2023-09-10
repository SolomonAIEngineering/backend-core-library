package postgres

import (
	"reflect"
	"testing"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
)

func TestWithQueryTimeout(t *testing.T) {
	type args struct {
		timeout *time.Duration
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
			if got := WithQueryTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithQueryTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMaxConnectionRetries(t *testing.T) {
	type args struct {
		retries *int
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
			if got := WithMaxConnectionRetries(tt.args.retries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMaxConnectionRetries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMaxConnectionRetryTimeout(t *testing.T) {
	type args struct {
		timeout *time.Duration
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
			if got := WithMaxConnectionRetryTimeout(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMaxConnectionRetryTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRetrySleep(t *testing.T) {
	type args struct {
		sleep *time.Duration
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
			if got := WithRetrySleep(tt.args.sleep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRetrySleep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithConnectionString(t *testing.T) {
	type args struct {
		connectionString *string
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
			if got := WithConnectionString(tt.args.connectionString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithConnectionString() = %v, want %v", got, tt.want)
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

func TestWithMaxIdleConnections(t *testing.T) {
	type args struct {
		maxIdleConnections *int
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
			if got := WithMaxIdleConnections(tt.args.maxIdleConnections); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMaxIdleConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMaxOpenConnections(t *testing.T) {
	type args struct {
		maxOpenConnections *int
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
			if got := WithMaxOpenConnections(tt.args.maxOpenConnections); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMaxOpenConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMaxConnectionLifetime(t *testing.T) {
	type args struct {
		maxConnectionLifetime *time.Duration
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
			if got := WithMaxConnectionLifetime(tt.args.maxConnectionLifetime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMaxConnectionLifetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithInstrumentationClient(t *testing.T) {
	type args struct {
		instrumentationClient *instrumentation.Client
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
			if got := WithInstrumentationClient(tt.args.instrumentationClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithInstrumentationClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
