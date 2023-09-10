package postgres

import (
	"reflect"
	"testing"
)

func TestNewInMemoryTestDbClient(t *testing.T) {
	type args struct {
		models []any
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInMemoryTestDbClient(tt.args.models...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInMemoryTestDbClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemoryTestDbClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ConfigureNewTxCleanupHandlerForUnitTests(t *testing.T) {
	tests := []struct {
		name string
		c    *Client
		want *TestTxCleanupHandlerForUnitTests
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ConfigureNewTxCleanupHandlerForUnitTests(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ConfigureNewTxCleanupHandlerForUnitTests() = %v, want %v", got, tt.want)
			}
		})
	}
}
