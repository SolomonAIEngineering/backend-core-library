package consumer

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockSQSClient struct {
	mock.Mock
	sqs.SQS
}

func (m *mockSQSClient) SendMessageWithContext(ctx aws.Context, input *sqs.SendMessageInput, opts ...request.Option) (*sqs.SendMessageOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*sqs.SendMessageOutput), args.Error(1)
}

func TestHandleAWSError(t *testing.T) {
	client := &ConsumerClient{}
	errOverLimit := awserr.New(sqs.ErrCodeOverLimit, "Over limit", errors.New("Over limit"))

	client.handleAWSError(errOverLimit)
	assert.Equal(t, 2*time.Second, client.backoffDuration, "Backoff duration should double on ErrCodeOverLimit")

	client.backoffDuration = 5 * time.Second
	client.handleAWSError(errOverLimit)
	assert.Equal(t, maxBackoffDuration, client.backoffDuration, "Backoff duration should not exceed maxBackoffDuration")

	errOther := awserr.New(sqs.ErrCodeInvalidMessageContents, "Invalid content", errors.New("Invalid content"))
	client.handleAWSError(errOther)
	assert.Equal(t, time.Second, client.backoffDuration, "Backoff duration should reset on other errors")
}

func TestMoveToDLQ(t *testing.T) {
	mockClient := new(mockSQSClient)
	client := &ConsumerClient{
		sqsClient:          mockClient,
		deadletterQueueUrl: aws.String("https://dlq.url"),
	}

	msg := &sqs.Message{
		Body: aws.String("test message"),
	}

	// Define behavior for mocked method
	mockClient.On("SendMessageWithContext", mock.Anything, mock.Anything).Return(&sqs.SendMessageOutput{}, nil)

	err := client.moveToDLQ(context.Background(), msg)
	assert.Nil(t, err, "moveToDLQ should not return an error with a valid DLQ URL")

	// Test with nil DLQ URL
	client.deadletterQueueUrl = nil
	err = client.moveToDLQ(context.Background(), msg)
	assert.NotNil(t, err, "moveToDLQ should return an error with a nil DLQ URL")
}
