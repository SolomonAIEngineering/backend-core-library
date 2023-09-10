package workerpool

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
