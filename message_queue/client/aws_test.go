// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package client

import (
	"reflect"
	"testing"
)

func TestWithEndpoint(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want AwsConfigOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithEndpoint(tt.args.endpoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRegion(t *testing.T) {
	type args struct {
		region string
	}
	tests := []struct {
		name string
		args args
		want AwsConfigOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithRegion(tt.args.region); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithProfile(t *testing.T) {
	type args struct {
		profile string
	}
	tests := []struct {
		name string
		args args
		want AwsConfigOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithProfile(tt.args.profile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want AwsConfigOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithSecret(t *testing.T) {
	type args struct {
		secret string
	}
	tests := []struct {
		name string
		args args
		want AwsConfigOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithSecret(tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithUseS3PathStyleAddressing(t *testing.T) {
	type args struct {
		useS3PathStyleAddressing bool
	}
	tests := []struct {
		name string
		args args
		want AwsConfigOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithUseS3PathStyleAddressing(tt.args.useS3PathStyleAddressing); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithUseS3PathStyleAddressing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAwsConfig_Apply(t *testing.T) {
	type args struct {
		opts []AwsConfigOption
	}
	tests := []struct {
		name string
		cfg  *AwsConfig
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cfg.Apply(tt.args.opts...)
		})
	}
}

func TestAwsConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *AwsConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("AwsConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewAwsConfig(t *testing.T) {
	type args struct {
		opts []AwsConfigOption
	}
	tests := []struct {
		name    string
		args    args
		want    *AwsConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAwsConfig(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAwsConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAwsConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
