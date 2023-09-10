package client // import "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Receive receives messages from an SQS queue specified by the
// `queueURL` parameter. It takes a `context.Context` parameter for handling timeouts and returns a
// slice of `Message` pointers and an error. It uses the AWS SDK to make a `ReceiveMessageWithContext`
// API call to the SQS service, passing in the `queueURL` and other parameters to retrieve a single
// message from the queue. It then converts the received message into a `Message` struct and appends it
// to the `messages` slice. Finally, it returns the `messages` slice or an error if there was a problem
// receiving the message.
func (h *Client) Receive(ctx context.Context, queueURL string) ([]*Message, error) {
	ctx, cancel := context.WithTimeout(ctx, h.ReadTimeout)
	defer cancel()

	res, err := h.SqsClient.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(1),
		AttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("receive: %w", err)
	}

	if len(res.Messages) == 0 {
		return nil, nil
	}

	messages := make([]*Message, 0)
	for _, message := range res.Messages {
		attrs := make(map[string]string)
		for key, attr := range message.MessageAttributes {
			attrs[key] = *attr.StringValue
		}

		messages = append(messages, &Message{
			ID:            *message.MessageId,
			ReceiptHandle: *message.ReceiptHandle,
			Body:          *message.Body,
			Attributes:    attrs,
		})
	}
	return messages, nil
}
