package taskhandler

import "github.com/SimifiniiCTO/asynq"

// `type ITaskHandler interface` is defining an interface that specifies the methods that must be
// implemented by a task handler. This interface is likely used by the `TaskProcessor` to ensure that
// any task handler that is used with the processor implements the required methods. The `ITaskHandler`
// interface may include methods such as `RegisterTaskHandler()` and `Validate()`, which are used to
// register the task handler with the Asynq task queue and validate that the handler is properly
// configured, respectively. By defining this interface, the `TaskProcessor` can ensure that any task
// handler used with the processor conforms to a specific set of requirements, making it easier to swap
// out different task handlers as needed.
type ITaskHandler interface {
	// The `RegisterTaskHandler()` method is likely a method that needs to be implemented by any task
	// handler that is used with the `TaskProcessor`. It returns a pointer to an instance of the
	// `asynq.ServeMux` struct, which is used to register the task handler with the Asynq task queue
	// system. The `ServeMux` struct is a multiplexer that allows multiple task handlers to be registered
	// with the Asynq task queue, each handling a different type of task. By registering the task handler
	// with the `ServeMux`, the Asynq task queue system knows which handler to use for each type of task
	// that is enqueued. This method is likely called during the initialization of the `TaskProcessor` to
	// ensure that the task handler is properly registered with the Asynq task queue system.
	RegisterTaskHandler() *asynq.ServeMux
	// The `Validate()` method is a method of the `TaskProcessor` struct that checks if all the required
	// properties of the struct have been set. It returns an error if any of the required properties are
	// not set, indicating that the `TaskProcessor` is not ready to start processing tasks. This method is
	// typically called before starting the worker to ensure that all the necessary dependencies are in
	// place.
	Validate() error
}
