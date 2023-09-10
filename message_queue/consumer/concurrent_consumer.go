package consumer // import "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/consumer"

import (
	"context"
	"time"

	client "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"
)

// ConcurrentConsumer creates a limited parallel queue, and continues to poll AWS until all the limit is reached.
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
func (c *ConsumerClient) StartConcurentConsumer(f MessageProcessorFunc) {
	var (
		messages []*client.Message
		err      error
	)
	sync := createFullBufferedChannel(c.ConcurrencyFactor)
	for {
		ctx := context.Background()
		messages, err = c.SqsClient.Receive(ctx, *c.QueueUrl)
		if err != nil {
			c.reportErrorEvent("receive_message", err)
			continue
		}

		if len(messages) == 0 {
			time.Sleep(c.QueuePollingDuration)
		} else {
			for _, message := range messages {
				// request the exact amount of "workers" from pool.
				// Again, empty buffer will block this operation
				<-sync

				// process message concurrently
				go c.processMessageConcurrently(ctx, message, f, sync)
			}
		}
	}
}

// processMessageConcurrentlyprocesses a single message concurrently.
// It takes in a context, a message to process, a function to process the message, and a channel of
// booleans to synchronize the processing of messages. The function calls the provided message
// processing function with the given context and message, and if there is an error, it reports it. If
// there is no error, it deletes the message from the queue. Finally, it returns a boolean value to the
// synchronization channel to indicate that the worker is available to process another message.
func (c *ConsumerClient) processMessageConcurrently(ctx context.Context, message *client.Message, f MessageProcessorFunc, sync chan bool) {
	err := f(ctx, message)
	if err != nil {
		c.reportErrorEvent("process_message", err)
	} else {
		// delete message from queue if no error was encountered
		if err := c.SqsClient.Delete(ctx, *c.QueueUrl, message.ReceiptHandle); err != nil {
			c.reportErrorEvent("delete_message", err)
		}
	}

	// return "worker" to the "pool"
	sync <- true
}
