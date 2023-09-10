package mongo

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestNew(t *testing.T) {
	type args struct {
		options []Option
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

func TestClient_GetCollection(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *mongo.Collection
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetCollection(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetConnection(t *testing.T) {
	tests := []struct {
		name string
		c    *Client
		want *mongo.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetConnection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_connect(t *testing.T) {
	tests := []struct {
		name    string
		c       *Client
		want    *mongo.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_createCollections(t *testing.T) {
	tests := []struct {
		name    string
		c       *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.createCollections(); (err != nil) != tt.wantErr {
				t.Errorf("Client.createCollections() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		collectionSet []string
		collection    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.collectionSet, tt.args.collection); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
