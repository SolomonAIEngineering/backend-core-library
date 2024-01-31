package scheduleclient

import (
	"context"

	"go.temporal.io/sdk/client"
)

// Task represents a schedulable task with an ID, specification, and action.
type Task struct {
	Id     string                         `json:"id"`     // Unique identifier for the task.
	Spec   *client.ScheduleSpec           `json:"spec"`   // Specification of the scheduling.
	Action *client.ScheduleWorkflowAction `json:"action"` // Action to be executed by the task.
}

// Validate checks if the task has all the necessary information set.
// It returns an error if any required field is missing.
func (task *Task) Validate() error {
	if task.Id == "" {
		return ErrTaskIdNotSet
	}

	if task.Spec == nil {
		return ErrTaskSpecNotSet
	}

	if task.Action == nil {
		return ErrTaskActionNotSet
	}

	return nil
}

// ScheduleTask schedules a new task using the provided context and task details.
// It returns a handle to the scheduled task and any error encountered.
func (sc *ScheduleClient) ScheduleTask(ctx context.Context, task *Task) (*client.ScheduleHandle, error) {
	if err := task.Validate(); err != nil {
		return nil, err
	}

	hdl, err := sc.ScheduleClient.Create(ctx, client.ScheduleOptions{
		ID:     task.Id,
		Spec:   *task.Spec,
		Action: task.Action,
	})

	if err != nil {
		return nil, err
	}

	return &hdl, nil
}

// UpdateScheduledTask updates an existing scheduled task using the provided options.
// It takes a context, the task's handler, and update options as arguments.
func (c *ScheduleClient) UpdateScheduledTask(
	ctx context.Context,
	handler client.ScheduleHandle,
	options *client.ScheduleUpdateOptions) error {
	return handler.Update(ctx, *options)
}

// PauseScheduledTask pauses an existing scheduled task.
// It takes a context, the task's handler, and pause options as arguments.
func (c *ScheduleClient) PauseScheduledTask(
	ctx context.Context,
	handler client.ScheduleHandle,
	options *client.SchedulePauseOptions) error {
	return handler.Pause(ctx, *options)
}

// ResumeScheduledTask resumes a paused scheduled task.
// It takes a context, the task's handler, and unpause options as arguments.
func (c *ScheduleClient) ResumeScheduledTask(
	ctx context.Context,
	handler client.ScheduleHandle,
	options *client.ScheduleUnpauseOptions) error {
	return handler.Unpause(ctx, *options)
}

// DeleteScheduledTask deletes an existing scheduled task.
// It takes a context and the task's handler as arguments.
func (c *ScheduleClient) DeleteScheduledTask(
	ctx context.Context,
	handler client.ScheduleHandle) error {
	return handler.Delete(ctx)
}
