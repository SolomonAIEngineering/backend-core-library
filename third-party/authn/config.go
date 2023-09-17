// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package authn // import "github.com/SolomonAIEngineering/backend-core-library/third-party/authn"

const (
	// DefaultKeychainTTL is the default TTL for a key in keychain in minutes.
	DefaultKeychainTTL = 60
)

// Config is a configuration struct for Client.
type Config struct {
	Issuer         string //the base url of the service handling authentication
	PrivateBaseURL string //overrides the base url for private endpoints
	Audience       string //the domain (host) of the main application
	Username       string //the http basic auth username for accessing private endpoints of the lib issuer
	Password       string //the http basic auth password for accessing private endpoints of the lib issuer
	KeychainTTL    int    //TTL for a key in keychain in minutes
}

// `setDefaults()` is a method of the `Config` struct that sets default values for some of its fields
// if they are not already set. Specifically, if `KeychainTTL` is not set, it is set to the value of
// `DefaultKeychainTTL` (which is 60). If `PrivateBaseURL` is not set, it is set to the value of
// `Issuer`. This method is called to ensure that the `Config` struct has all the necessary fields set
// before it is used to create a new `Client` instance.
func (c *Config) setDefaults() {
	if c.KeychainTTL == 0 {
		c.KeychainTTL = DefaultKeychainTTL
	}
	if c.PrivateBaseURL == "" {
		c.PrivateBaseURL = c.Issuer
	}
}
