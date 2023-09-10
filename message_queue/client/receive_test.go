package client

import (
	"context"
	"reflect"
	"testing"
)

func TestClient_Receive(t *testing.T) {
	type args struct {
		ctx      context.Context
		queueURL string
	}
	tests := []struct {
		name    string
		h       *Client
		args    args
		want    []*Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.Receive(tt.args.ctx, tt.args.queueURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Receive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Receive() = %v, want %v", got, tt.want)
			}
		})
	}
}
