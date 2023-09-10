package consumer // import "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/consumer"

import "fmt"

// reportErrorEvent reports an error event to the logger and New
// Relic client. It takes two parameters: `op` which is a string representing the operation that caused
// the error, and `err` which is the error object. The method logs the error message using the logger's
// `Error` method and records a custom event using the New Relic client's `RecordCustomEvent` method.
// The custom event has a type of "message_queue.error" and includes a map of additional data with keys
// "operation" and "error" that contain the values of `op` and `err.Error()` respectively.
func (c *ConsumerClient) reportErrorEvent(op string, err error) {
	c.Logger.Error(err.Error())
	// TODO: ideally we should emit a custom event here
}

// reportProcessedMessageCount increments the processed message count metric
func (c *ConsumerClient) reportProcessedMessageCount(op string) {
	c.Logger.Info(fmt.Sprintf("Processed message %s", op))
	// TODO: ideally we should increment the processed message count metric here
}
