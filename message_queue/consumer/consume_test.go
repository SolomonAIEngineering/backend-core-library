package consumer

import (
	"context"
	"testing"
	"time"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type mockSQSClient struct {
	mock.Mock
	sqs.SQS
	messages []*sqs.Message
}

func (m *mockSQSClient) SendMessageWithContext(ctx aws.Context, input *sqs.SendMessageInput, opts ...request.Option) (*sqs.SendMessageOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*sqs.SendMessageOutput), args.Error(1)
}

func (m *mockSQSClient) ReceiveMessageWithContext(ctx aws.Context, input *sqs.ReceiveMessageInput, opts ...request.Option) (*sqs.ReceiveMessageOutput, error) {
	return &sqs.ReceiveMessageOutput{
		Messages: m.messages,
	}, nil
}

func (m *mockSQSClient) DeleteMessageWithContext(ctx aws.Context, input *sqs.DeleteMessageInput, opts ...request.Option) (*sqs.DeleteMessageOutput, error) {
	return &sqs.DeleteMessageOutput{}, nil
}

func TestMoveToDLQ(t *testing.T) {
	mockClient := new(mockSQSClient)
	client := &ConsumerClient{
		sqsClient:          mockClient,
		deadletterQueueUrl: aws.String("https://dlq.url"),
		logger:             zap.NewNop(),
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

func TestConsumerClient_Start(t *testing.T) {
	// Define a dummy message for our mock SQS client to return
	dummyMessage := &sqs.Message{
		Body: aws.String("test message"),
	}

	mockClient := &mockSQSClient{
		messages: []*sqs.Message{dummyMessage},
	}

	// Channel to signal when the dummy message is processed
	messageProcessed := make(chan bool)

	// Use a channel to signal when polling is done
	doneCh := make(chan bool)

	client := &ConsumerClient{
		logger:                zap.NewNop(),
		instrumentationClient: &instrumentation.Client{},
		sqsClient:             mockClient,
		queueUrl:              aws.String("test-queue-url"),
		deadletterQueueUrl:    aws.String("test-dlq-url"),
		concurrencyFactor:     1,
		queuePollingDuration:  time.Millisecond * 100,
		messageProcessTimeout: 0,
		handler: func(ctx context.Context, message *sqs.Message) error {
			if *message.Body == "test message" {
				messageProcessed <- true
			}
			return nil
		},
		stopCh:          make(chan bool),
		backoffDuration: 3 * time.Second,
		batchSize:       10,
		waitTimeSecond:  1,
	}

	go func() {
		client.Start()
		close(doneCh) // Signal the main test routine that polling has finished
	}()

	select {
	case <-doneCh:
		// Finished successfully
	case <-time.After(25 * time.Second):
		t.Fatal("Test timed out")
	}

	mockClient.AssertExpectations(t)
}
