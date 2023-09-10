package manager

import (
	"context"
	"time"

	"go.temporal.io/api/serviceerror"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"google.golang.org/grpc"
)

// Instantiates a new worker
func NewWorker(client client.Client, taskQueue string, options worker.Options) worker.Worker {
	return worker.New(client, taskQueue, options)
}

// Instantiates a new client
func NewClient(opts *client.Options) (client.Client, error) {
	if opts == nil {
		return nil, ErrInvalidClientOptions
	}
	// Create the client object just once per process
	c, err := client.Dial(*opts)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// NewNamespaceClient creates an instance of a namespace client, to manage
// lifecycle of namespaces.
func NewNamespaceClient(opts *client.Options) (client.NamespaceClient, error) {
	if opts == nil {
		return nil, ErrInvalidClientOptions
	}

	// calls the initialize a new namespace will not attempt to connect to the
	// temporal cluster eagerly hence the call may not fail even if the server is unreachable
	// we need to pass grpc.WithBlock as a gRPC dial option to connection options to eagerly connect
	connectionOptions := opts.ConnectionOptions.DialOptions
	connectionOptions = append(connectionOptions, grpc.WithBlock())
	opts.ConnectionOptions.DialOptions = connectionOptions

	// Create the client object just once per process
	c, err := client.NewNamespaceClient(*opts)
	if err != nil {
		return nil, err
	}

	// upon successfully creating the namespace client, we ensure to also create the namespace
	// on which our workers will process tasks
	// TODO: read this from env variables
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	workflowRetentionPeriod := time.Hour * 24
	if err := c.Register(ctx, &workflowservice.RegisterNamespaceRequest{
		Namespace:                        opts.Namespace,
		OwnerEmail:                       "admin@simfinii.com",
		WorkflowExecutionRetentionPeriod: &workflowRetentionPeriod,
	}); err != nil {
		// ignore the error if the namespace already exists
		if _, ok := err.(*serviceerror.NamespaceAlreadyExists); !ok {
			return nil, err
		}
	}

	return c, nil
}
