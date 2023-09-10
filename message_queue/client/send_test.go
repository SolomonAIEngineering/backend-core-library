package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func TestClient_Send(t *testing.T) {
	type args struct {
		ctx context.Context
		req *SendRequest
	}
	tests := []struct {
		name    string
		h       *Client
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.Send(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.Send() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SendMessage(t *testing.T) {
	type args struct {
		ctx context.Context
		msg *sqs.SendMessageInput
	}
	tests := []struct {
		name    string
		h       *Client
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.SendMessage(tt.args.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SendMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
