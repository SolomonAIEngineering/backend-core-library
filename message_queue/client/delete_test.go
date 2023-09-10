package client

import (
	"context"
	"testing"
)

func TestClient_Delete(t *testing.T) {
	type args struct {
		ctx       context.Context
		queueURL  string
		rcvHandle string
	}
	tests := []struct {
		name    string
		h       *Client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Delete(tt.args.ctx, tt.args.queueURL, tt.args.rcvHandle); (err != nil) != tt.wantErr {
				t.Errorf("Client.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
