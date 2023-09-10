package mongo

import (
	"os/exec"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestNewInMemoryTestDbClient(t *testing.T) {
	type args struct {
		collectionNames []string
	}
	tests := []struct {
		name    string
		args    args
		want    *InMemoryTestDbClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInMemoryTestDbClient(tt.args.collectionNames)
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

func TestInMemoryTestDbClient_Teardown(t *testing.T) {
	tests := []struct {
		name    string
		c       *InMemoryTestDbClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Teardown(); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryTestDbClient.Teardown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setupInMemoryMongoDB(t *testing.T) {
	tests := []struct {
		name    string
		want    *mongo.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setupInMemoryMongoDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("setupInMemoryMongoDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupInMemoryMongoDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stopConflictingProcesses(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stopConflictingProcesses(tt.args.port)
		})
	}
}

func Test_exec_cmd(t *testing.T) {
	type args struct {
		cmd *exec.Cmd
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exec_cmd(tt.args.cmd)
		})
	}
}
