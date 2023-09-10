package mongo

import (
	"reflect"
	"testing"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func TestWithMaxConnectionAttempts(t *testing.T) {
	type args struct {
		maxConnectionAttempts int
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
			if got := WithMaxConnectionAttempts(tt.args.maxConnectionAttempts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMaxConnectionAttempts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMaxRetriesPerOperation(t *testing.T) {
	type args struct {
		maxRetriesPerOperation int
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
			if got := WithMaxRetriesPerOperation(tt.args.maxRetriesPerOperation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMaxRetriesPerOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRetryTimeOut(t *testing.T) {
	type args struct {
		retryTimeOut time.Duration
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
			if got := WithRetryTimeOut(tt.args.retryTimeOut); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRetryTimeOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithOperationSleepInterval(t *testing.T) {
	type args struct {
		operationSleepInterval time.Duration
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
			if got := WithOperationSleepInterval(tt.args.operationSleepInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithOperationSleepInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithQueryTimeout(t *testing.T) {
	type args struct {
		queryTimeout time.Duration
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
			if got := WithQueryTimeout(tt.args.queryTimeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithQueryTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithDatabaseName(t *testing.T) {
	type args struct {
		databaseName string
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
			if got := WithDatabaseName(tt.args.databaseName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDatabaseName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithTelemetry(t *testing.T) {
	type args struct {
		telemetry *instrumentation.Client
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
			if got := WithTelemetry(tt.args.telemetry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTelemetry() = %v, want %v", got, tt.want)
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

func TestWithCollectionNames(t *testing.T) {
	type args struct {
		collectionNames []string
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
			if got := WithCollectionNames(tt.args.collectionNames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCollectionNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithClientOptions(t *testing.T) {
	type args struct {
		clientOptions *options.ClientOptions
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
			if got := WithClientOptions(tt.args.clientOptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithClientOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithConnectionURI(t *testing.T) {
	type args struct {
		connectionURI string
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
			if got := WithConnectionURI(tt.args.connectionURI); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithConnectionURI() = %v, want %v", got, tt.want)
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
