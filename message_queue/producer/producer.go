package producer // import "github.com/SolomonAIEngineering/backend-core-library/message_queue/producer"

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

const (
	maxRetries         = 3
	initialBackoffTime = 1 * time.Second
	maxBackoffTime     = 10 * time.Second
)

type SQSAPI interface {
	SendMessageWithContext(ctx aws.Context, input *sqs.SendMessageInput, opts ...request.Option) (*sqs.SendMessageOutput, error)
	SendMessageBatchWithContext(ctx aws.Context, input *sqs.SendMessageBatchInput, opts ...request.Option) (*sqs.SendMessageBatchOutput, error)
}

type ProducerClient struct {
	sqsClient SQSAPI
	queueUrl  *string
	logger    *zap.Logger
}

func NewProducerClient(sqsClient SQSAPI, queueUrl *string, logger *zap.Logger) *ProducerClient {
	return &ProducerClient{
		sqsClient: sqsClient,
		queueUrl:  queueUrl,
		logger:    logger,
	}
}

// Send a single message
func (p *ProducerClient) SendMessage(ctx context.Context, body string, attributes map[string]*sqs.MessageAttributeValue) error {
	req := &sqs.SendMessageInput{
		QueueUrl:          p.queueUrl,
		MessageBody:       aws.String(body),
		MessageAttributes: attributes,
	}

	var lastErr error
	for i := 0; i < maxRetries; i++ {
		_, err := p.sqsClient.SendMessageWithContext(ctx, req)
		if err != nil {
			lastErr = err
			p.handleAWSError(err)

			// Calculate backoff time
			backoffTime := time.Duration(int64(initialBackoffTime) * (1 << uint(i)))
			if backoffTime > maxBackoffTime {
				backoffTime = maxBackoffTime
			}
			time.Sleep(backoffTime)
		} else {
			// Success
			return nil
		}
	}
	return lastErr
}

// Send messages in batches
func (p *ProducerClient) SendMessagesBatch(ctx context.Context, messages []*sqs.SendMessageBatchRequestEntry) error {
	req := &sqs.SendMessageBatchInput{
		QueueUrl: p.queueUrl,
		Entries:  messages,
	}

	var lastErr error
	for i := 0; i < maxRetries; i++ {
		_, err := p.sqsClient.SendMessageBatchWithContext(ctx, req)
		if err != nil {
			lastErr = err
			p.handleAWSError(err)

			// Calculate backoff time
			backoffTime := time.Duration(int64(initialBackoffTime) * (1 << uint(i)))
			if backoffTime > maxBackoffTime {
				backoffTime = maxBackoffTime
			}
			time.Sleep(backoffTime)
		} else {
			// Success
			return nil
		}
	}
	return lastErr
}

// Handle AWS-specific errors
func (p *ProducerClient) handleAWSError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case sqs.ErrCodeOverLimit:
			p.logger.Warn("Hit message limit on SQS", zap.Error(aerr))
		case sqs.ErrCodeMessageNotInflight:
			p.logger.Warn("Message not in flight", zap.Error(aerr))
		//... Add more specific error codes as required.
		default:
			p.logger.Error("Encountered error with SQS", zap.Error(aerr))
		}
	}
}
