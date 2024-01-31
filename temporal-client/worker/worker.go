package worker // import "github.com/SolomonAIEngineering/backend-core-library/temporal-client/worker"

import (
	"log"

	temporalWorkerClient "go.temporal.io/sdk/worker"
)

type Worker struct {
	config   *WorkerConfig
	executor temporalWorkerClient.Worker
}

// New creates and returns a new Worker based on the provided WorkerConfig.
// This function will validate the provided configuration and create the Temporal worker.
// In case of invalid configuration, the application will terminate with a log.Fatal.
//
// Example:
//
//  config := worker.NewWorkerConfig(worker.WithTaskQueue("exampleQueue"), ...)
//  w := worker.New(config)
//  err = w.Start(worker.InterruptCh())
//  if err != nil {
// 	 log.Fatalln("Unable to start the Worker Process", err)
//  }

func New(config *WorkerConfig) *Worker {
	// validate the worker configuration
	if config == nil {
		log.Fatal("worker config is nil")
	}

	if err := config.validate(); err != nil {
		log.Fatal(err)
	}

	// Create the worker
	executor := temporalWorkerClient.New(config.Client, config.TaskQueue, config.WorkerOptions)

	return &Worker{
		config:   config,
		executor: executor,
	}
}

// Start initiates the execution of the Temporal worker, allowing it to start processing workflows and activities.
// It will block until an interrupt signal is received.
// In case of failure to start the worker, the application will terminate with a log.Fatal.
func (c *Worker) Start() error {
	// Register the activities and workflows
	// NOTE: must register workflows before activities
	c.registerWorkflows(c.config.Workflows)
	c.registerActivities(c.config.Activities)

	err := c.executor.Run(temporalWorkerClient.InterruptCh())
	if err != nil {
		return err
	}

	return nil
}

// registerWorkflow registers a single workflow with the Temporal worker.
// The workflow parameter should be a function that adheres to Temporal's workflow signature requirements.
//
// Example:
//
//	w.registerWorkflow(myWorker, myWorkflow)
func (c *Worker) registerWorkflow(workflow interface{}) {
	c.executor.RegisterWorkflow(workflow)
}

// registerWorkflows registers multiple workflows with the Temporal worker.
// Each workflow should be a function that adheres to Temporal's workflow signature requirements.
// Workflows are provided as a slice of interface{}.
//
// Example:
//
//	workflows := []interface{}{workflow1, workflow2, ...}
//	w.registerWorkflows(workflows)
func (w *Worker) registerWorkflows(workflows []interface{}) {
	for _, workflow := range workflows {
		w.executor.RegisterWorkflow(workflow)
	}
}

// registerActivity registers a single activity with the Temporal worker.
// The activity parameter should be a function that adheres to Temporal's activity signature requirements.
//
// Example:
//
//	w.registerActivity(myWorker, myActivity)
func (c *Worker) registerActivity(worker temporalWorkerClient.Worker, activity interface{}) {
	worker.RegisterActivity(activity)
}

// registerActivities registers multiple activities with the Temporal worker.
// Each activity should be a function that adheres to Temporal's activity signature requirements.
// Activities are provided as a slice of interface{}.
//
// Example:
//
//	activities := []interface{}{activity1, activity2, ...}
//	w.registerActivities(activities)
func (w *Worker) registerActivities(activities []interface{}) {
	for _, activity := range activities {
		w.executor.RegisterActivity(activity)
	}
}
