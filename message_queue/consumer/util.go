package consumer // import "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/consumer"

// `createFullBufferedChannel` creates a buffered channel of boolean values
// with a specified capacity. It initializes the channel with `capacity` number of `true` values
// and returns the channel.
func createFullBufferedChannel(capacity int) chan bool {
	sync := make(chan bool, capacity)

	for i := 0; i < capacity; i++ {
		sync <- true
	}
	return sync
}
