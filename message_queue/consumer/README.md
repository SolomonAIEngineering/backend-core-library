## Consumer 
To build a super performant SQS consumer in Golang, there are several best practices to follow. Here are some recommendations:

- Use long polling: Use long polling instead of short polling to receive messages from SQS. Long polling reduces the number of requests and therefore reduces the cost and latency. Set the WaitTimeSeconds parameter to a high value (up to 20 seconds).

- Use multiple goroutines: Consume messages from SQS using multiple goroutines to process messages concurrently. Create a worker pool and dispatch incoming messages to the workers. This approach can improve performance by taking advantage of multiple cores and minimizing idle time.

- Reduce message processing time: Keep message processing time to a minimum to avoid blocking the worker pool. If possible, move the heavy processing to another service or to a batch processing job.

- Use batch deletion: Use batch deletion to delete messages from SQS. This approach reduces the number of API calls and can significantly improve performance.

- Enable buffering: Use buffering to reduce the number of API calls. Buffer messages in memory or on disk before processing them. This approach can be especially useful when processing large messages or when processing messages at high volumes.

```go
package main

import (
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	maxNumberOfMessages = 10
	waitTimeSeconds     = 20
)

func main() {
	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	// Create an SQS client
	sqsClient := sqs.New(sess)

	// Create a WaitGroup to synchronize worker goroutines
	var wg sync.WaitGroup

	// Create a channel to receive messages
	msgChan := make(chan *sqs.Message)

	// Create a worker pool
	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, msgChan, sqsClient, &wg)
	}

	// Create a receive message input
	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String("<QUEUE_URL>"),
		MaxNumberOfMessages: aws.Int64(maxNumberOfMessages),
		WaitTimeSeconds:     aws.Int64(waitTimeSeconds),
	}

	// Continuously receive messages from SQS and dispatch to worker pool
	for {
		output, err := sqsClient.ReceiveMessage(input)
		if err != nil {
			log.Printf("Failed to receive messages: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		if len(output.Messages) == 0 {
			continue
		}

		// Dispatch messages to worker pool
		for _, message := range output.Messages {
			msgChan <- message
		}
	}

	// Wait for worker pool to complete
	close(msgChan)
	wg.Wait()
}

// Worker function to process messages
func worker(id int, msgChan <-chan *sqs.Message, sqsClient *sqs.SQS, wg *sync.WaitGroup) {
	defer wg.Done()

	// Process incoming messages
	for message := range msgChan {
	    log.Printf("Worker %d received message: %s", id, *message.Body)

	    // Process message here...

	    // Delete the message from the queue
	    if _, err := sqsClient.DeleteMessage(&s,sqs.DeleteMessageInput{
                QueueUrl: aws.String("<QUEUE_URL>"),
                ReceiptHandle: message.ReceiptHandle,
                }); err != nil {
        
            log.Printf("Failed to delete message: %v", err)
        }
    }
}
```