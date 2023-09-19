package workerpool // import "github.com/SimifiniiCTO/simfiny-core-lib/worker-pool"

type (
	TaskType   string
	TaskStatus string
)

const (
	TaskStatusWaiting   TaskStatus = "task:waiting"
	TaskStatusRunning   TaskStatus = "task:running"
	TaskStatusCompleted TaskStatus = "task:completed"
	TaskStatusFailed    TaskStatus = "task:failed"
)
