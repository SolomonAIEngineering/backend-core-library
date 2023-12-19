package kmsClient

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// AWSKMSConfig is the configuration for the AWS KMS implementation of the KeyManagement interface
type AWSKMSConfig struct {
	// Log is the logger to use for this implementation
	Log *zap.Logger
	// KeyID is the ID of the key to use for encryption and decryption
	AccessKey string
	// Region is the AWS region to use for the KMS client
	Region string
	// SecretKey is the secret key to use for the KMS client
	SecretKey string
	// Endpoint is the endpoint to use for the KMS client
	Endpoint *string
	// KmsKeyID is the ID of the key to use for encryption and decryption
	KmsKeyID string
}

// `func (c *AWSKMSConfig) Validate() error` is a method defined on the `AWSKMSConfig` struct. It
// validates that all the required fields in the `AWSKMSConfig` struct are present and returns an error
// if any of them are missing. This method is called when creating a new instance of the `AWSKMS`
// struct to ensure that the configuration is valid.
func (c *AWSKMSConfig) Validate() error {
	if c.Log == nil {
		return errors.New("log is required")
	}

	if c.AccessKey == "" {
		return errors.New("access key is required")
	}

	if c.Region == "" {
		return errors.New("region is required")
	}

	if c.SecretKey == "" {
		return errors.New("secret key is required")
	}

	if c.KmsKeyID == "" {
		return errors.New("kms key id is required")
	}

	return nil
}

// AWSKMS is the implementation of the KeyManagement interface for AWS KMS
type AWSKMS struct {
	// log is the logger to use for this implementation
	log *zap.Logger
	// config is the configuration for this implementation
	config AWSKMSConfig
	// client is the AWS KMS client to use for this implementation
	client kmsiface.KMSAPI
}

// NewAWSKMS creates a new AWS KMS implementation of the KeyManagement interface
func NewAWSKMS(config AWSKMSConfig) (KeyManagement, error) {
	if err := config.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid configuration")
	}

	options := session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
			Region:      aws.String(config.Region),
		},
	}

	if config.Endpoint != nil && *config.Endpoint != "" {
		options.Config.Endpoint = config.Endpoint
	}

	awsSession, err := session.NewSessionWithOptions(options)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create aws session")
	}

	client := kms.New(awsSession)
	return &AWSKMS{
		log:    config.Log,
		config: config,
		client: client,
	}, nil
}

// Encrypt encrypts the input using the AWS KMS client
func (a *AWSKMS) Encrypt(ctx context.Context, input []byte) (keyID string, version string, result []byte, _ error) {
	request := &kms.EncryptInput{
		EncryptionAlgorithm: aws.String("RSAES_OAEP_SHA_256"),
		KeyId:               aws.String(a.config.KmsKeyID),
		Plaintext:           input,
	}

	response, err := a.client.EncryptWithContext(ctx, request)
	if err != nil {
		return "", "", nil, errors.Wrap(err, "failed to encrypt data using AWS KMS")
	}

	return *response.KeyId, "", response.CiphertextBlob, nil
}

// Decrypt decrypts the input using the AWS KMS client
func (a *AWSKMS) Decrypt(ctx context.Context, keyID string, version string, input []byte) (result []byte, _ error) {

	request := &kms.DecryptInput{
		CiphertextBlob:      input,
		EncryptionAlgorithm: aws.String("RSAES_OAEP_SHA_256"), // TODO Maybe make this a config thing?
		KeyId:               aws.String(keyID),
	}

	response, err := a.client.DecryptWithContext(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decrypt data using AWS KMS")
	}

	return response.Plaintext, nil
}
