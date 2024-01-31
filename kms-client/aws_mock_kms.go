package kmsClient // import "github.com/SolomonAIEngineering/backend-core-library/kms-client"

import (
	"context"
	"crypto/rand"
	"math/big"
)

type MockKeyManagement struct{}

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// KeyManagement is the interface that defines the methods for encrypting and decrypting data
type KeyManagement interface {
	Encrypt(ctx context.Context, input []byte) (keyID, version string, result []byte, _ error)
	Decrypt(ctx context.Context, keyID, version string, input []byte) (result []byte, _ error)
}

var _ KeyManagement = (*MockKeyManagement)(nil)

func NewMockAwsKms() (KeyManagement, error) {
	return &MockKeyManagement{}, nil
}

// Decrypt implements KeyManagement
func (*MockKeyManagement) Decrypt(ctx context.Context, keyID string, version string, input []byte) (result []byte, _ error) {
	result, err := generateRandomBytes(10)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Encrypt implements KeyManagement
func (*MockKeyManagement) Encrypt(ctx context.Context, input []byte) (keyID string, version string, result []byte, _ error) {
	keyId, err := generateRandomString(10)
	if err != nil {
		return "", "", nil, err
	}

	version, err = generateRandomString(10)
	if err != nil {
		return "", "", nil, err
	}

	result, err = generateRandomBytes(10)
	if err != nil {
		return "", "", nil, err
	}

	return keyId, version, result, nil
}

func generateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

func generateRandomString(length int) (string, error) {
	randomString := make([]byte, length)

	// Generate random indices within the letterBytes string
	for i := range randomString {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		randomString[i] = letterBytes[n.Int64()]
	}

	return string(randomString), nil
}
