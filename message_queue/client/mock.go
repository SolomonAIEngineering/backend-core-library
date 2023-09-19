package client // import "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type MockClient struct{}

// Delete implements MessageClientInterface.
func (*MockClient) Delete(ctx context.Context, queueURL string, rcvHandle string) error {
	return nil
}

// Receive implements MessageClientInterface.
func (*MockClient) Receive(ctx context.Context, queueURL string) ([]*Message, error) {
	return []*Message{
		{
			ID:            "test-id",
			ReceiptHandle: "test-receipt-handle",
			Body:          "test-body",
			Attributes: map[string]string{
				"test-key": "test-value",
			},
		},
	}, nil
}

// Send implements MessageClientInterface.
func (*MockClient) Send(ctx context.Context, req *SendRequest) (string, error) {
	return "test-id", nil
}

// SendMessage implements MessageClientInterface.
func (*MockClient) SendMessage(ctx context.Context, msg *sqs.SendMessageInput) (*string, error) {
	str := "test-id"
	return &str, nil
}

var _ MessageClientInterface = (*MockClient)(nil)

func NewMockClient() (MessageClientInterface, error) {
	return &MockClient{}, nil
}
