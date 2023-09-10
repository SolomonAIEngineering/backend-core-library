package consumer // import "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/consumer"

import (
	"context"
	"sync"
	"time"

	client "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"
)

// With an SQS message subscriber we will be receiving messages in small batches out of the box,
// In order for our message consumer to achieve a high throughput, we will process the messages in parallel,
// and in order this to be robust, we should impose a limit on how many messages we should process simultaneously.

// As standard aws sqs receive call gives us maximum of 10 messages, the naive approach will be to process them
//
//	in parallel, then call the next batch.
//
// With approach like this we will be limited to the
// 1 minute / slowest message processing in batch * 10, for example having the slowest message being processed in 50ms
// it will give us (1000 ms / 50ms) * 10 = 200 messages per second of processing time minus network latency, that can
// eat up most of the projected capacity.
func (c *ConsumerClient) NaiveConsumer(f MessageProcessorFunc) {
	for {
		ctx := context.Background()
		messages, err := c.SqsClient.Receive(ctx, *c.QueueUrl)
		if err != nil {
			c.reportErrorEvent("receive_message", err)
			continue
		}

		wg := sync.WaitGroup{}
		for _, message := range messages {
			wg.Add(1)
			go c.processMessage(ctx, message, f, &wg)
		}

		if len(messages) == 0 { // add aditional sleep if queue is empty
			time.Sleep(c.QueuePollingDuration)
			continue
		}
	}
}

// processMessage processes a single message received
// from an SQS queue. It takes in the context, the message to be processed, a function to process the
// message, and a wait group to synchronize the processing of multiple messages. It calls the
// processing function with the message and if there is an error, it reports it. If the processing is
// successful, it deletes the message from the queue. The wait group is used to ensure that all
// messages in a batch are processed before moving on to the next batch.
func (c *ConsumerClient) processMessage(ctx context.Context, message *client.Message, f MessageProcessorFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	err := f(ctx, message)
	if err != nil {
		c.reportErrorEvent("process_message", err)
		return
	}

	err = c.SqsClient.Delete(ctx, *c.QueueUrl, message.ReceiptHandle)
	if err != nil {
		c.reportErrorEvent("delete_message", err)
		return
	}
}
