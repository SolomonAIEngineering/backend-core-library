// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package authn // import "github.com/SimifiniiCTO/simfiny-core-lib/third-party/authn"

import (
	"errors"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"

	jose "gopkg.in/square/go-jose.v2"
)

// Mock JWKProvider for tests
// Stores keys locally.
type mockJwkProvider struct {
	keyMap   map[string]jose.JSONWebKey
	hitCount int
}

func newMockJwkProvider() *mockJwkProvider {
	return &mockJwkProvider{
		keyMap: map[string]jose.JSONWebKey{
			"kid1": {
				KeyID: "kid1",
				Key:   "test_key1",
			},
			"kid2": {
				KeyID: "kid2",
				Key:   "test_key2",
			},
		},
	}
}

func (m *mockJwkProvider) Key(kid string) ([]jose.JSONWebKey, error) {
	m.hitCount = m.hitCount + 1

	if kid == "kidError" {
		return []jose.JSONWebKey{}, errors.New("testing error")
	}

	if jwk, ok := m.keyMap[kid]; ok {
		return []jose.JSONWebKey{jwk}, nil
	}
	return []jose.JSONWebKey{}, nil
}

func TestKeychainCacheHit(t *testing.T) {
	mockProvider := newMockJwkProvider()
	keychainCache := newKeychainCache(time.Minute, mockProvider)

	keys1, err := keychainCache.Key("kid1")
	assert.NoError(t, err)
	assert.Len(t, keys1, 1)
	assert.Equal(t, "kid1", keys1[0].KeyID)
	assert.Equal(t, "test_key1", keys1[0].Key)
	assert.Equal(t, 1, mockProvider.hitCount)

	keys1Again, err := keychainCache.Key("kid1")
	assert.NoError(t, err)
	assert.Len(t, keys1Again, 1)
	assert.Equal(t, "kid1", keys1Again[0].KeyID)
	assert.Equal(t, "test_key1", keys1Again[0].Key)
	assert.Equal(t, 1, mockProvider.hitCount) //Because we cached it

	keys2, err := keychainCache.Key("kid2")
	assert.NoError(t, err)
	assert.Len(t, keys2, 1)
	assert.Equal(t, "kid2", keys2[0].KeyID)
	assert.Equal(t, "test_key2", keys2[0].Key)
	assert.Equal(t, 2, mockProvider.hitCount) //Because key2 wasnt cached
}

func TestKeychainCacheMissing(t *testing.T) {
	mockProvider := newMockJwkProvider()
	keychainCache := newKeychainCache(time.Minute, mockProvider)

	keysNone, err := keychainCache.Key("kidNone")
	assert.NoError(t, err)
	assert.Len(t, keysNone, 0)
	assert.Equal(t, 1, mockProvider.hitCount)

	keysNoneAgain, err := keychainCache.Key("kidNone")
	assert.NoError(t, err)
	assert.Len(t, keysNoneAgain, 0)
	assert.Equal(t, 2, mockProvider.hitCount) //Because missing keys are not cached
}

func TestKeychainCacheTTL(t *testing.T) {
	mockProvider := newMockJwkProvider()
	keychainCache := newKeychainCache(time.Minute, mockProvider)
	// Minimum TTL is 1 min. But we cant wait that long to test.
	// TODO: Go is not flexible enough to accept both decimal and integer ttl. Consider using seconds?
	// Hacky test because we are screwing with internals
	keychainCache.keyCache = cache.New(time.Second, time.Second)

	_, _ = keychainCache.Key("kid1")
	assert.Equal(t, 1, mockProvider.hitCount)
	_, _ = keychainCache.Key("kid1")
	assert.Equal(t, 1, mockProvider.hitCount) //Because we cached itached

	// Wait for cache to expire
	time.Sleep(time.Second)
	keys1, err := keychainCache.Key("kid1")
	assert.NoError(t, err)
	// Assert values post-expiry
	assert.Len(t, keys1, 1)
	assert.Equal(t, "kid1", keys1[0].KeyID)
	assert.Equal(t, "test_key1", keys1[0].Key)
	assert.Equal(t, 2, mockProvider.hitCount) //Because cache expired
}

func TestKeychainCacheError(t *testing.T) {
	mockProvider := newMockJwkProvider()
	keychainCache := newKeychainCache(time.Minute, mockProvider)

	keysError, err := keychainCache.Key("kidError")
	assert.EqualError(t, err, "testing error")
	assert.Len(t, keysError, 0)
	assert.Equal(t, 1, mockProvider.hitCount)

	keysErrorAgain, err := keychainCache.Key("kidError")
	assert.EqualError(t, err, "testing error")
	assert.Len(t, keysErrorAgain, 0)
	assert.Equal(t, 2, mockProvider.hitCount) //Because keys are not cached in case of error
}
