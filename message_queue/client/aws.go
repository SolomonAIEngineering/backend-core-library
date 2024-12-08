// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package client // import "github.com/SolomonAIEngineering/backend-core-library/message_queue/client"

import (
	"errors"
)

// The `AwsConfig` struct represents the configuration options for an AWS service, including endpoint,
// region, authentication credentials, and S3 addressing style.
// @property {string} Endpoint - represents the endpoint for the AWS service.
// @property {string} Region - The `Region` field in the `AwsConfig` struct represents the region for
// the AWS service. It is a string type and can be set using the `WithRegion` function, which is an
// `AwsConfigOption` that sets the value of this field when passed as an argument to the `
// @property {string} Profile - The `Profile` field in the `AwsConfig` struct represents the AWS
// profile to use for authentication. It is a string type and can be set using the `WithProfile`
// function, which is an `AwsConfigOption` that sets the value of this field when passed as an argument
// to the
// @property {string} ID - The `ID` field is a string type field in the `AwsConfig` struct that
// represents the AWS access key ID to use for authentication. It can be set using the `WithID`
// function, which is an `AwsConfigOption` that sets the value of this field when passed as an
// @property {string} Secret - The `Secret` field is a string type field in the `AwsConfig` struct that
// represents the AWS secret access key to use for authentication. It can be set using the `WithSecret`
// function, which is an `AwsConfigOption` that sets the value of this field when passed as an
// @property {bool} UseS3PathStyleAddressing - `UseS3PathStyleAddressing` is a boolean field in the
// `AwsConfig` struct that determines whether to use path-style or virtual-hosted style addressing when
// making requests to Amazon S3. If it is set to `true`, path-style addressing will be used, and if it
// is
type AwsConfig struct {
	// `Endpoint` is a field in the `AwsConfig` struct that represents the endpoint for the AWS service. It
	// is a string type. The `WithEndpoint` function is an `AwsConfigOption` that sets the value of this
	// field when passed as an argument to `NewAwsConfig` function.
	Endpoint string
	// `Region` is a field in the `AwsConfig` struct that represents the region for the AWS service. It is
	// a string type. The `WithRegion` function is an `AwsConfigOption` that sets the value of this field
	// when passed as an argument to `NewAwsConfig` function.
	Region string
	// `Profile` is a field in the `AwsConfig` struct that represents the AWS profile to use for
	// authentication. It is a string type. The `WithProfile` function is an `AwsConfigOption` that sets
	// the value of this field when passed as an argument to `NewAwsConfig` function.
	Profile string
	// The `ID` field is a string type field in the `AwsConfig` struct that represents the AWS access key
	// ID to use for authentication. It can be set using the `WithID` function, which is an
	// `AwsConfigOption` that sets the value of this field when passed as an argument to the `NewAwsConfig`
	// function.
	ID string
	// The `Secret` field is a string type field in the `AwsConfig` struct that represents the AWS secret
	// access key to use for authentication. It can be set using the `WithSecret` function, which is an
	// `AwsConfigOption` that sets the value of this field when passed as an argument to the `NewAwsConfig`
	// function.
	Secret string
	// `UseS3PathStyleAddressing` is a boolean field in the `AwsConfig` struct that determines whether to
	// use path-style or virtual-hosted style addressing when making requests to Amazon S3. If it is set
	// to `true`, path-style addressing will be used, and if it is set to `false`, virtual-hosted style
	// addressing will be used. This option is set using the `WithUseS3PathStyleAddressing` function,
	// which is an `AwsConfigOption` that sets the value of this field when passed as an argument to the
	// `NewAwsConfig` function.
	UseS3PathStyleAddressing bool
}

// AwsConfigOption is a function that configures a AwsConfig.
type AwsConfigOption func(*AwsConfig)

// WithEndpoint sets an endpoint.
func WithEndpoint(endpoint string) AwsConfigOption {
	return func(c *AwsConfig) {
		c.Endpoint = endpoint
	}
}

// WithRegion sets a region.
func WithRegion(region string) AwsConfigOption {
	return func(c *AwsConfig) {
		c.Region = region
	}
}

// WithProfile sets a profile.
func WithProfile(profile string) AwsConfigOption {
	return func(c *AwsConfig) {
		c.Profile = profile
	}
}

// WithID sets an ID.
func WithID(id string) AwsConfigOption {
	return func(c *AwsConfig) {
		c.ID = id
	}
}

// WithSecret sets a secret.
func WithSecret(secret string) AwsConfigOption {
	return func(c *AwsConfig) {
		c.Secret = secret
	}
}

// WithUseS3PathStyleAddressing sets a useS3PathStyleAddressing.
func WithUseS3PathStyleAddressing(useS3PathStyleAddressing bool) AwsConfigOption {
	return func(c *AwsConfig) {
		c.UseS3PathStyleAddressing = useS3PathStyleAddressing
	}
}

// Apply applies the given AwsConfigOptions to this AwsConfig.
func (cfg *AwsConfig) Apply(opts ...AwsConfigOption) {
	for _, opt := range opts {
		opt(cfg)
	}
}

// Validate validates this AwsConfig.
func (cfg *AwsConfig) Validate() error {
	if cfg.ID == "" {
		return errors.New("aws id is required")
	}
	if cfg.Secret == "" {
		return errors.New("aws secret is required")
	}
	if cfg.Region == "" {
		return errors.New("aws region is required")
	}

	if cfg.Endpoint == "" {
		return errors.New("aws endpoint is required")
	}

	if cfg.Profile == "" {
		return errors.New("aws profile is required")
	}

	return nil
}

// NewAwsConfig creates a new AwsConfig.
func NewAwsConfig(opts ...AwsConfigOption) (*AwsConfig, error) {
	cfg := &AwsConfig{}
	cfg.Apply(opts...)
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}
