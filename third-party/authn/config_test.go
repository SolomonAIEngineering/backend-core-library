// Copyright (C) Solomon AI, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package authn // import "SolomonAIEngineering/backend-core-library/third-party/authn"

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigDefaults(t *testing.T) {
	c := Config{
		Issuer:   "test_issuer",
		Audience: "test_audience",
		Username: "test_username",
		Password: "test_password",
	}
	c.setDefaults()
	assert.Equal(t, c.Issuer, "test_issuer")
	assert.Equal(t, c.PrivateBaseURL, "test_issuer")
	assert.Equal(t, c.Audience, "test_audience")
	assert.Equal(t, c.Username, "test_username")
	assert.Equal(t, c.Password, "test_password")
	assert.Equal(t, c.KeychainTTL, DefaultKeychainTTL)
}

func TestConfigDefaultsOverride(t *testing.T) {
	c := Config{
		Issuer:         "test_issuer",
		Audience:       "test_audience",
		Username:       "test_username",
		Password:       "test_password",
		PrivateBaseURL: "test_private_url",
		KeychainTTL:    500,
	}
	c.setDefaults()
	assert.Equal(t, c.Issuer, "test_issuer")
	assert.Equal(t, c.PrivateBaseURL, "test_private_url")
	assert.Equal(t, c.Audience, "test_audience")
	assert.Equal(t, c.Username, "test_username")
	assert.Equal(t, c.Password, "test_password")
	assert.Equal(t, c.KeychainTTL, 500)
}
