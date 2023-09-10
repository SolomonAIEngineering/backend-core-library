package mongo

import (
	"context"
	"reflect"
	"testing"
)

func TestClient_ComplexTransaction(t *testing.T) {
	type args struct {
		ctx      context.Context
		callback MongoTx
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ComplexTransaction(tt.args.ctx, tt.args.callback)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ComplexTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ComplexTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_StandardTransaction(t *testing.T) {
	type args struct {
		ctx      context.Context
		callback MongoTx
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.StandardTransaction(tt.args.ctx, tt.args.callback); (err != nil) != tt.wantErr {
				t.Errorf("Client.StandardTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
