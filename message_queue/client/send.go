package client // import "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Send sends a message to a target queue present in the request object
func (h *Client) Send(ctx context.Context, req *SendRequest) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, h.WriteTimeout)
	defer cancel()

	attrs := make(map[string]*sqs.MessageAttributeValue, len(req.Attributes))
	for _, attr := range req.Attributes {
		attrs[attr.Key] = &sqs.MessageAttributeValue{
			StringValue: aws.String(attr.Value),
			DataType:    aws.String(attr.Type),
		}
	}

	res, err := h.SqsClient.SendMessageWithContext(ctx, &sqs.SendMessageInput{
		MessageAttributes: attrs,
		MessageBody:       aws.String(req.Body),
		QueueUrl:          aws.String(req.QueueURL),
	})
	if err != nil {
		return "", fmt.Errorf("send: %w", err)
	}

	return *res.MessageId, nil
}

// SendMessage sends a message to a queue.
func (h *Client) SendMessage(ctx context.Context, msg *sqs.SendMessageInput) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, h.WriteTimeout)
	defer cancel()

	res, err := h.SqsClient.SendMessageWithContext(ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("could not send message to queue %v: %v", msg.QueueUrl, err)
	}

	return res.MessageId, nil
}
