package producer

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type mockSQSClient struct {
	SQSAPI
	sendMessageErr         error
	sendMessageBatchErr    error
	sendMessageOutput      *sqs.SendMessageOutput
	sendMessageBatchOutput *sqs.SendMessageBatchOutput
}

func (m *mockSQSClient) SendMessageWithContext(ctx aws.Context, input *sqs.SendMessageInput, opts ...request.Option) (*sqs.SendMessageOutput, error) {
	return m.sendMessageOutput, m.sendMessageErr
}

func (m *mockSQSClient) SendMessageBatchWithContext(ctx aws.Context, input *sqs.SendMessageBatchInput, opts ...request.Option) (*sqs.SendMessageBatchOutput, error) {
	return m.sendMessageBatchOutput, m.sendMessageBatchErr
}

func TestSendMessage(t *testing.T) {
	mockClient := &mockSQSClient{
		sendMessageErr: awserr.New(sqs.ErrCodeOverLimit, "Over limit", nil),
	}
	logger, _ := zap.NewProduction()
	producer := NewProducerClient(mockClient, aws.String("fakeQueueURL"), logger)

	err := producer.SendMessage(context.TODO(), "test message", nil)
	assert.Error(t, err)
}

func TestSendMessagesBatch(t *testing.T) {
	entries := []*sqs.SendMessageBatchRequestEntry{
		{
			Id:          aws.String("testID1"),
			MessageBody: aws.String("test message 1"),
		},
		{
			Id:          aws.String("testID2"),
			MessageBody: aws.String("test message 2"),
		},
	}
	mockClient := &mockSQSClient{
		sendMessageBatchErr: awserr.New(sqs.ErrCodeOverLimit, "Over limit", nil),
	}
	logger, _ := zap.NewProduction()
	producer := NewProducerClient(mockClient, aws.String("fakeQueueURL"), logger)

	err := producer.SendMessagesBatch(context.TODO(), entries)
	assert.Error(t, err)
}

func TestHandleAWSError(t *testing.T) {
	mockClient := &mockSQSClient{}
	logger, _ := zap.NewProduction()
	producer := NewProducerClient(mockClient, aws.String("fakeQueueURL"), logger)

	err := awserr.New(sqs.ErrCodeOverLimit, "Over limit", nil)
	producer.handleAWSError(err)
	// Here you might want to check if your logger received the expected error, using an in-memory or hook-based logger. This is just a basic test.
}
