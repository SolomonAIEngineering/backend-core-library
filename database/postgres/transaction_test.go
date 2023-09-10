// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package postgres

import (
	"context"
	"reflect"
	"testing"
)

func TestClient_PerformTransaction(t *testing.T) {
	type args struct {
		ctx         context.Context
		transaction Tx
	}
	tests := []struct {
		name    string
		db      *Client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.PerformTransaction(tt.args.ctx, tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("Client.PerformTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_PerformComplexTransaction(t *testing.T) {
	type args struct {
		ctx         context.Context
		transaction CmplxTx
	}
	tests := []struct {
		name    string
		db      *Client
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.PerformComplexTransaction(tt.args.ctx, tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PerformComplexTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.PerformComplexTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
