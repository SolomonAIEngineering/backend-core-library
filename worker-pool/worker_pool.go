package workerpool // import "github.com/SimifiniiCTO/simfiny-core-lib/worker-pool"

// This implementation defines a WorkerPool struct that contains a channel
// of channels (workers) that can hold a maximum of `max-workers` worker channels
//
// To use the goroutine pool, you can create a WorkerPool instance and start it:
//
//	```go
//		pool := NewWorkerPool(10)
//		pool.Start()
//	```
type WorkerPool struct {
	workers chan chan func()
}

// NewWorkerPool creates a new WorkerPool with
// the given number of workers
func NewWorkerPool(maxWorkers int) *WorkerPool {
	workers := make(chan chan func(), maxWorkers)
	return &WorkerPool{workers: workers}
}

// Start creates a number of goroutines that listen
// for tasks on their worker channel and executes them
func (w *WorkerPool) Start() {
	for i := 0; i < cap(w.workers); i++ {
		go func() {
			for {
				select {
				case task := <-w.workers:
					taskFn := <-task
					taskFn()
				}
			}
		}()
	}
}

// ExecuteTask takes a function `task` as an argument. It creates a new channel `taskChan` and sends it to the `workers`
// channel of the `WorkerPool`. Then it sends the `task` function to the `taskChan` channel. This
// method is used to add a new task to the worker pool for execution.
func (w *WorkerPool) ExecuteTask(task func()) {
	taskChan := make(chan func())
	w.workers <- taskChan
	taskChan <- task
}
