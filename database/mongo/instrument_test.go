package mongo

import (
	"context"
	"reflect"
	"testing"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func TestClient_StartDbSegment(t *testing.T) {
	type args struct {
		ctx            context.Context
		name           string
		collectionName string
	}
	tests := []struct {
		name string
		c    *Client
		args args
		want *newrelic.DatastoreSegment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.StartDbSegment(tt.args.ctx, tt.args.name, tt.args.collectionName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.StartDbSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}
