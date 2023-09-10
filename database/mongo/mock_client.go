package mongo // import "github.com/SimifiniiCTO/simfiny-core-lib/database/mongo"

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	mim "github.com/tryvium-travels/memongo"
)

const (
	mAX_DATABASE_CONNECTION_ATTEMPTS       = 3
	mAX_DATABASE_RETRIES_PER_OPERATION     = 3
	rETRY_TIMEOUT                          = 30 * time.Second
	rETRY_SLEEP_INTERVAL                   = 5 * time.Second
	qUERY_TIMEOUT                          = 1 * time.Minute
	dATABASE_NAME                          = "database"
	dEFAULT_PORT                       int = 27107
)

type InMemoryTestDbClient struct {
	DatabaseName string
	Client       *Client
	Server       *mim.Server
}

// NewInMemoryTestDbClient creates a new in-memory MongoDB database and returns a client to it.
func NewInMemoryTestDbClient(collectionNames []string) (*InMemoryTestDbClient, error) {
	stopConflictingProcesses(dEFAULT_PORT)
	server, err := mim.StartWithOptions(&mim.Options{MongoVersion: "6.0.0", ShouldUseReplica: false, DownloadURL: "https://fastdl.mongodb.org/osx/mongodb-macos-arm64-6.0.0.tgz"})
	if err != nil {
		return nil, err
	}
	uri := server.URI()

	opts := []Option{
		WithDatabaseName(dATABASE_NAME),
		WithRetryTimeOut(rETRY_TIMEOUT),
		WithOperationSleepInterval(rETRY_SLEEP_INTERVAL),
		WithMaxConnectionAttempts(mAX_DATABASE_CONNECTION_ATTEMPTS),
		WithMaxRetriesPerOperation(mAX_DATABASE_RETRIES_PER_OPERATION),
		WithQueryTimeout(qUERY_TIMEOUT),
		WithTelemetry(&instrumentation.Client{}),
		WithLogger(zap.L()),
		WithCollectionNames(collectionNames),
		WithConnectionURI(uri),
		WithClientOptions(options.Client().ApplyURI(uri).SetDirect(true)),
	}
	conn, err := New(opts...)
	if err != nil {
		return nil, err
	}

	return &InMemoryTestDbClient{
		DatabaseName: dATABASE_NAME,
		Client:       conn,
		Server:       server,
	}, nil

}

// Teardown cleans up resources initialized by Setup.
// This function must be called once after all tests have finished running.
func (c *InMemoryTestDbClient) Teardown() error {
	// Dropping the test database causes an error against Atlas Data Lake.
	conn := c.Client.Conn
	table := c.DatabaseName
	inMemoryServer := c.Server

	defer inMemoryServer.Stop()

	if err := conn.Database(table).Drop(context.Background()); err != nil {
		return fmt.Errorf("error dropping test database: %v", err)
	}

	if err := conn.Disconnect(context.Background()); err != nil {
		return fmt.Errorf("error disconnecting test client: %v", err)
	}

	return nil
}

func setupInMemoryMongoDB() (*mongo.Client, error) {
	// Create an in-memory MongoDB database
	mongoURI := "mongodb://localhost/test?authSource=$external&authMechanism=MONGODB-X509"
	clientOptions := options.Client().ApplyURI(mongoURI).SetDirect(true)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %v", err)
	}

	return client, nil
}

// stopConflictingProcesses stops any conflicting process running on a the defined port
func stopConflictingProcesses(port int) {
	if runtime.GOOS == "windows" {
		command := fmt.Sprintf("(Get-NetTCPConnection -LocalPort %d).OwningProcess -Force", port)
		exec_cmd(exec.Command("Stop-Process", "-Id", command))
	} else {
		command := fmt.Sprintf("lsof -i tcp:%d | grep LISTEN | awk '{print $2}' | xargs kill -9", port)
		exec_cmd(exec.Command("bash", "-c", command))
	}
}

// Execute command and return exited code.
func exec_cmd(cmd *exec.Cmd) {
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		}
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			fmt.Printf("Error during killing (exit code: %s)\n", []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Printf("Port successfully killed (exit code: %s)\n", []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
}
