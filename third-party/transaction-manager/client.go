package manager

import (
	"context"
	"fmt"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/database/mongo"
	"github.com/SimifiniiCTO/simfiny-core-lib/database/postgres"
	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	msq "github.com/SimifiniiCTO/simfiny-core-lib/message_queue/client"

	"go.temporal.io/api/serviceerror"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	ErrInvalidClientOptions                = fmt.Errorf("invalid input argument. client options cannot be null")
	ErrInvalidTransactionManagerConfig     = fmt.Errorf("invalid input argument. transaction manager config cannot be nil")
	ErrInvalidLogger                       = fmt.Errorf("invalid logger. logger cannot be nil")
	ErrInvalidTelemetrySDK                 = fmt.Errorf("invalid telemetry sdk config. config cannot be nil")
	ErrNilAccount                          = fmt.Errorf("account cannot be nil")
	ErrNilOldEmail                         = fmt.Errorf("old email cannot be nil")
	ErrInvalidUpdateAccountWorkflowRequest = fmt.Errorf("invalid update account workflow request. request cannot be nil or invalid")
	ErrInvalidRetryPolicy                  = fmt.Errorf("invalid input argument. retry policy cannot be nil or invalid")
	ErrInvalidRpcTimeout                   = fmt.Errorf("invalid input argument. rpc timeout cannot be nil or invalid")
	ErrInvalidConfigurations               = fmt.Errorf("invalid input argument. transaction manager configurations cannot be nil or invalid")
)

// `TransactionManager` is the single struct by which we manage and initiate all distributed transactions
// within the service. It provides wrapper facilities around the temporal sdk client as well in order
// to properly emit metrics and traces during rpc operations
//
// @property client - This is the client used to interact with a remote temporal cluster.
// @property TelemetrySDK - This is the telemetry SDK that we use to send telemetry data to newrelic
// @property Logger - This is the logger that will be used to log messages.
// @property AuthenticationServiceClient - This is the gRPC client for the Authentication Service.
// @property FinancialIntegrationServiceClient - This is the gRPC client for the Financial Integration
// Service.
// @property SocialServiceClient - This is the client for the Social Service.
// @property MessageQueueSDK - This is the message queue SDK that we will use to send messages to the
// queue.
// @property DatabaseConn - This is the database connection object that we will use to connect to the
// database.
type TransactionManager struct {
	temporalClient          client.Client
	temporalNamespaceClient client.NamespaceClient
	options                 *client.Options
	instrumentationClient   *instrumentation.Client
	logger                  *zap.Logger
	messageQueueClient      *msq.Client
	mongoClient             *mongo.Client
	postgresClient          *postgres.Client
	retryPolicy             *Policy
	rpcTimeout              time.Duration
	worker                  worker.Worker
	metricsEnabled          bool
}

// Policy outlines retry policies necessary for workflow and downstream service calls
type Policy struct {
	RetryInitialInterval    *time.Duration
	RetryBackoffCoefficient float64
	MaximumInterval         time.Duration
	MaximumAttempts         int
}

type WorkflowManager interface {
	Close()
	Start()
}

var _ WorkflowManager = &TransactionManager{}

// It creates a new instance of the TransactionManager struct and returns it
func NewTransactionManager(options ...Option) (*TransactionManager, error) {
	txm := &TransactionManager{}
	for _, o := range options {
		o(txm)
	}

	if txm.options != nil {
		client, err := configureTemporalClient(txm.options)
		if err != nil {
			return nil, err
		}

		namespaceClient, err := configureTemporalNamespaceClient(txm.options)
		if err != nil {
			return nil, err
		}

		txm.temporalClient = client
		txm.temporalNamespaceClient = namespaceClient
		txm.worker = worker.New(client, txm.options.Namespace, worker.Options{})
	}

	// initialize the temporal worker process to process items off the task queue
	return txm, nil
}

// Start enables the worker to start listening to a given task queue
// NOTE: This should be ran in a go routine otherwise the process will block
func (tx *TransactionManager) Start() {
	var (
		txmWorker = tx.worker
		log       = tx.logger
	)

	// register worklows and activities
	tx.registerWorkflowsAndActivities()

	// run the worker in a blocking fashion
	err := txmWorker.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal("unable to start Worker", zap.Error(err))
	}
}

// Close closes the client and all its underlying connections
// and clears up any associated reasources
func (t *TransactionManager) Close() {
	if t.temporalClient == nil {
		t.temporalClient.Close()
	}

	if t.temporalNamespaceClient == nil {
		t.temporalNamespaceClient.Close()
	}
}

// registerWorkflowsAndActivities registers all the workflows and activities that the worker will be
// responsible for processing
func (tx *TransactionManager) registerWorkflowsAndActivities() {
}

// Instantiates a new client
func configureTemporalClient(opts *client.Options) (client.Client, error) {
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
func configureTemporalNamespaceClient(opts *client.Options) (client.NamespaceClient, error) {
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
		OwnerEmail:                       "yoan@simfinii.com",
		WorkflowExecutionRetentionPeriod: &workflowRetentionPeriod,
	}); err != nil {
		// ignore the error if the namespace already exists
		if _, ok := err.(*serviceerror.NamespaceAlreadyExists); !ok {
			return nil, err
		}
	}

	return c, nil
}
