package consumer

import (
	"context"
	"sync"
	"time"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

const maxBackoffDuration = 30 * time.Second

var wg sync.WaitGroup

// MessageProcessorFunc defines the function type that processes messages from SQS.
// An error should be returned if processing fails, leading the message to be moved to the DLQ.
type MessageProcessorFunc = func(ctx context.Context, message *sqs.Message) error

// IConsumer provides an interface for consuming messages either concurrently or in a naive, sequential manner.
// Start: Begins processing messages concurrently.
// Stop: Halts the processing of messages.
type IConsumer interface {
	Start()
	Stop()
}

type SQSAPI interface {
	SendMessageWithContext(ctx aws.Context, input *sqs.SendMessageInput, opts ...request.Option) (*sqs.SendMessageOutput, error)
	ReceiveMessageWithContext(ctx aws.Context, input *sqs.ReceiveMessageInput, opts ...request.Option) (*sqs.ReceiveMessageOutput, error)
	DeleteMessageWithContext(ctx aws.Context, input *sqs.DeleteMessageInput, opts ...request.Option) (*sqs.DeleteMessageOutput, error)
}

// ConsumerClient encapsulates fields related to a client that consumes messages from a queue.
// It uses a logger for debugging, an instrumentation client for performance monitoring, and an SQS client to interact with the message queue.
type ConsumerClient struct {
	logger                *zap.Logger             // Logger instance for debugging and monitoring.
	instrumentationClient *instrumentation.Client // New Relic client for application performance monitoring.
	sqsClient             SQSAPI                  // SQS client to interact with the external service or API.
	queueUrl              *string                 // URL of the queue to receive messages.
	deadletterQueueUrl    *string                 // URL of the dead letter queue.
	concurrencyFactor     int                     // Maximum number of concurrent message processing operations.
	queuePollingDuration  time.Duration           // Duration for polling the queue for new messages.
	messageProcessTimeout time.Duration           // Maximum duration for message processing before considering it as timed out.

	handler         MessageProcessorFunc // Function to process the received messages.
	stopCh          chan bool            // Channel to signal stopping the client.
	backoffDuration time.Duration        // Duration before attempting to re-process a failed message.
	batchSize       int64                // Number of messages to fetch in one call.
	waitTimeSecond  int64                // Duration to wait between polling attempts.
	pollCount       int
	maxPolls        int // -1 for infinite, any other positive number for a max count
}

// Static check to ensure ConsumerClient implements the IConsumer interface.
var _ IConsumer = (*ConsumerClient)(nil)

// Start initiates the polling of the SQS queue in a separate goroutine.
//
// Example:
// poller := NewSQSPoller("us-west-2", "https://sqs.us-west-2.amazonaws.com/1234567890/myqueue",
//
//	"https://sqs.us-west-2.amazonaws.com/1234567890/mydlq", exampleHandler)
//
// poller.Start()
func (c *ConsumerClient) Start() {
	go c.poll()
}

// Stop sends a signal to the poller to stop polling.
//
// Example:
// poller := NewSQSPoller("us-west-2", "https://sqs.us-west-2.amazonaws.com/1234567890/myqueue",
//
//	"https://sqs.us-west-2.amazonaws.com/1234567890/mydlq", exampleHandler)
//
// poller.Start()
// time.Sleep(10 * time.Second)  // Let it poll for 10 seconds
// poller.Stop()
func (c *ConsumerClient) Stop() {
	c.stopCh <- true
	wg.Wait() // wait for all messages to finish processing
}

// poll creates a limited parallel queue, and continues to poll AWS until all the limit is reached.
// This is performed by implementing a token bucket” using a buffered channel hence this approach is only limited by aws throughput
//
// Some scenarios will require a different set of resources consumed, depending on the message type (Lets say you want your handler to be able to process from 1 to N emails in 1 message).
//
//	To maintain our limitations, we could introduce the timely based token bucket algorithm , which will ensure we don’t process more than N emails over a period of time (like 1 minute),
//	by grabbing the exact amount of “worker tokens” from the pool, depending on emails count in message. Also, if your code can be timed out, there is a good approach to impose timeout and cancellation,
//	based on golang context.WithCancel function. Check out the golang semaphore library to build the nuclear-resistant solution. (the mechanics are the same as in our example, abstracted to library,
//
// so instead of using channel for limiting our operation we will call semaphore.Acquire, which will also block our execution until “worker tokens” will be refilled).
//
// LINK - Ref: https://docs.microsoft.com/en-us/azure/architecture/microservices/model/domain-analysis
// LINK - Ref: https://docs.microsoft.com/en-us/azure/architecture/microservices/design/interservice-communication
func (c *ConsumerClient) poll() {
	workerTokens := make(chan bool, c.concurrencyFactor)
	for i := 0; i < c.concurrencyFactor; i++ {
		workerTokens <- true
	}

	for {
		select {
		case <-c.stopCh:
			return
		default:
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			if c.queueUrl == nil {
				c.logger.Error("Queue URL is nil")
				return
			}

			result, err := c.sqsClient.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
				QueueUrl: c.queueUrl,
				AttributeNames: []*string{
					aws.String(sqs.QueueAttributeNameAll),
				},
				MessageAttributeNames: []*string{
					aws.String(sqs.QueueAttributeNameAll),
				},
			})
			if err != nil {
				c.handleAWSError(err)
				c.logger.Error("Error while receiving message", zap.Error(err))
				continue
			}

			for _, message := range result.Messages {
				wg.Add(1)
				<-workerTokens

				go func(msg *sqs.Message) {
					defer wg.Done()
					c.process(ctx, msg, workerTokens)
				}(message)
			}

			c.pollCount++
			if c.maxPolls != -1 && c.pollCount > c.maxPolls {
				return
			}
		}
	}
}

// handleAWSError checks for AWS-specific errors and adjusts the backoff duration.
//
// Not intended to be called directly by users, hence no public-facing example.
func (c *ConsumerClient) handleAWSError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case sqs.ErrCodeOverLimit:
			c.backoffDuration = c.backoffDuration * 2
			if c.backoffDuration > maxBackoffDuration {
				c.backoffDuration = maxBackoffDuration
			}
		default:
			c.backoffDuration = time.Second
		}
	}
	time.Sleep(c.backoffDuration)
}

// moveToDLQ sends a failed message to the Dead Letter Queue (DLQ).
//
// Not intended to be called directly by users, hence no public-facing example.
func (c *ConsumerClient) moveToDLQ(ctx context.Context, message *sqs.Message) error {
	if c.deadletterQueueUrl == nil {
		c.logger.Error("DLQ URL is nil")
		return awserr.New(sqs.ErrCodeQueueDoesNotExist, "DLQ URL is nil", nil)
	}

	req := &sqs.SendMessageInput{
		QueueUrl:    c.deadletterQueueUrl,
		MessageBody: message.Body,
	}

	_, err := c.sqsClient.SendMessageWithContext(ctx, req)
	if err != nil {
		c.logger.Error("Failed to move message to DLQ", zap.Error(err))
	}
	return err
}

// process a single message .
// It takes in a context, a message to process, and a channel of
// booleans to synchronize the processing of messages. The function calls the provided message
// processing function with the given context and message, and if there is an error, it reports it. If
// there is no error, it deletes the message from the queue. Finally, it returns a boolean value to the
// synchronization channel to indicate that the worker is available to process another message.
func (c *ConsumerClient) process(ctx context.Context, message *sqs.Message, sync chan bool) {
	err := c.handler(ctx, message)
	if err != nil {
		c.moveToDLQ(ctx, message)
		c.reportErrorEvent("process_message", err)
	} else {
		if c.queueUrl == nil {
			c.logger.Error("Queue URL is nil")
			return
		}

		// delete message from queue if no error was encountered
		if _, err := c.sqsClient.DeleteMessageWithContext(ctx, &sqs.DeleteMessageInput{
			QueueUrl:      c.queueUrl,
			ReceiptHandle: message.ReceiptHandle,
		}); err != nil {
			c.reportErrorEvent("delete_message", err)
		}
	}

	// return "worker" to the "pool"
	sync <- true
}
