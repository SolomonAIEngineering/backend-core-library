package kmsClient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockKeyManagement_Decrypt(t *testing.T) {
	keyManagement := &MockKeyManagement{}
	ctx := context.Background()
	keyID := "keyID"
	version := "version"
	input := []byte("encryptedData")

	result, err := keyManagement.Decrypt(ctx, keyID, version, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NotNil(t, keyID)
	assert.NotNil(t, version)
	assert.NotNil(t, result)
}

func TestMockKeyManagement_Encrypt(t *testing.T) {
	keyManagement := &MockKeyManagement{}
	ctx := context.Background()
	input := []byte("plaintextData")

	keyID, version, result, err := keyManagement.Encrypt(ctx, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.NotNil(t, keyID)
	assert.NotNil(t, version)
	assert.NotNil(t, result)
}
