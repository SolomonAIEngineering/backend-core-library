package kmsClient

import (
	"bytes"
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// MockKMSClient is a mock implementation of the AWS KMS client for testing purposes
type MockKMSClient struct {
	kmsiface.KMSAPI
}

// EncryptWithContext is a mock implementation of the EncryptWithContext method of the AWS KMS client
func (m *MockKMSClient) EncryptWithContext(ctx aws.Context, input *kms.EncryptInput, opts ...request.Option) (*kms.EncryptOutput, error) {
	// Implement the desired mock behavior here
	return &kms.EncryptOutput{
		KeyId:               input.KeyId,
		CiphertextBlob:      []byte("encryptedData"),
		EncryptionAlgorithm: input.EncryptionAlgorithm,
	}, nil
}

// DecryptWithContext is a mock implementation of the DecryptWithContext method of the AWS KMS client
func (m *MockKMSClient) DecryptWithContext(ctx aws.Context, input *kms.DecryptInput, opts ...request.Option) (*kms.DecryptOutput, error) {
	// Implement the desired mock behavior here
	return &kms.DecryptOutput{
		KeyId:     input.KeyId,
		Plaintext: []byte("decryptedData"),
	}, nil
}

func TestAWSKMS_Encrypt(t *testing.T) {
	config := AWSKMSConfig{
		Log:       zap.NewNop(),
		AccessKey: "accessKey",
		Region:    "us-west-2",
		SecretKey: "secretKey",
		KmsKeyID:  "kmsKeyID",
	}

	mockClient := &MockKMSClient{}
	awsKMS := &AWSKMS{
		log:    zap.NewNop(),
		config: config,
		client: mockClient,
	}

	// Prepare test data
	ctx := context.Background()
	input := []byte("Hello, world!")

	// Call the Encrypt method
	keyID, version, result, err := awsKMS.Encrypt(ctx, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify the results
	expectedKeyID := "kmsKeyID"
	if keyID != expectedKeyID {
		t.Errorf("Unexpected key ID. Got: %s, want: %s", keyID, expectedKeyID)
	}

	expectedVersion := ""
	if version != expectedVersion {
		t.Errorf("Unexpected version. Got: %s, want: %s", version, expectedVersion)
	}

	expectedResult := []byte("encryptedData")
	if !bytes.Equal(result, expectedResult) {
		t.Errorf("Unexpected result. Got: %v, want: %v", result, expectedResult)
	}
}

func TestAWSKMS_Decrypt(t *testing.T) {
	config := AWSKMSConfig{
		Log:       zap.NewNop(),
		AccessKey: "accessKey",
		Region:    "us-west-2",
		SecretKey: "secretKey",
		KmsKeyID:  "kmsKeyID",
	}

	mockClient := &MockKMSClient{}
	awsKMS := &AWSKMS{
		log:    zap.NewNop(),
		config: config,
		client: mockClient,
	}

	// Prepare test data
	ctx := context.Background()
	keyID := "keyID"
	version := ""
	input := []byte("encryptedData")

	// Call the Decrypt method
	result, err := awsKMS.Decrypt(ctx, keyID, version, input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify the result
	expectedResult := []byte("decryptedData")
	if !bytes.Equal(result, expectedResult) {
		t.Errorf("Unexpected result. Got: %v, want: %v", result, expectedResult)
	}
}

func TestNewAWSKMS(t *testing.T) {
	// Prepare test data
	config := AWSKMSConfig{
		Log:       zap.L(),
		AccessKey: "your-access-key",
		Region:    "your-region",
		SecretKey: "your-secret-key",
		Endpoint:  aws.String("your-endpoint"),
		KmsKeyID:  "random-key-id",
	}

	// Call the function to be tested
	km, err := NewAWSKMS(config)

	// Assert the expected behavior and results
	assert.NoError(t, err)                // Check that there are no errors
	assert.NotNil(t, km)                  // Check that a non-nil instance is returned
	assert.NotNil(t, km.(*AWSKMS).client) // Check that the client is set correctly
}

func TestAWSKMSConfig_Validate(t *testing.T) {
	// Prepare test data
	config := AWSKMSConfig{
		Log:       zap.L(), // Provide your own logger implementation
		AccessKey: "your-access-key",
		SecretKey: "your-secret-key",
		Region:    "your-region",
		Endpoint:  aws.String("your-endpoint"),
		KmsKeyID:  "your-kms-key-id",
	}

	// Call the method to be tested
	err := config.Validate()

	// Assert the expected behavior and results
	assert.NoError(t, err) // Check that there are no errors

	// Test the required fields individually
	config.Log = nil
	err = config.Validate()
	assert.EqualError(t, err, "log is required") // Check that the error message matches

	config.Log = zap.L() // Reset the logger for the next test

	config.AccessKey = ""
	err = config.Validate()
	assert.EqualError(t, err, "access key is required")

	config.AccessKey = "your-access-key" // Reset the access key for the next test

	config.Region = ""
	err = config.Validate()
	assert.EqualError(t, err, "region is required")

	config.Region = "your-region" // Reset the region for the next test

	config.SecretKey = ""
	err = config.Validate()
	assert.EqualError(t, err, "secret key is required")

	config.SecretKey = "your-secret-key" // Reset the secret key for the next test

	config.KmsKeyID = ""
	err = config.Validate()
	assert.EqualError(t, err, "kms key id is required")
}
