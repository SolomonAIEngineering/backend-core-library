package mongo // import "github.com/SimifiniiCTO/simfiny-core-lib/database/mongo"

import (
	"context"
	"fmt"
	"time"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type IClient interface{}

type Client struct {
	// Conn serves as the database co
	Conn *mongo.Client
	// Logger is the logging utility used by this object
	Logger *zap.Logger
	// MaxConnectionAttempts outlines the maximum connection attempts
	// to initiate against the database
	MaxConnectionAttempts int
	// MaxRetriesPerOperation defines the maximum retries to attempt per failed database
	// connection attempt
	MaxRetriesPerOperation int
	// RetryTimeOut defines the maximum time until a retry operation is observed as a
	// timed out operation
	RetryTimeOut time.Duration
	// OperationSleepInterval defines the amount of time between retry operations
	// that the system sleeps
	OperationSleepInterval time.Duration
	// QueryTimeout defines the maximal amount of time a query can execute before being cancelled
	QueryTimeout time.Duration
	// Telemetry defines the object by which we will emit metrics, trace requests, and database operations
	Telemetry *instrumentation.Client
	// DatabaseName database name
	DatabaseName *string
	// Collection that operations will be performed against
	collectionNameToCollectionObjectMap map[string]*mongo.Collection
	// CollectionNames is a list of collection names
	CollectionNames []string
	// ClientOptions defines the options to use when connecting to the database
	ClientOptions *options.ClientOptions
	// ConnectionURI defines the connection string to use when connecting to the database
	ConnectionURI *string
}

var _ IClient = (*Client)(nil)

// New creates a new instance of the mongo client
func New(options ...Option) (*Client, error) {
	client := &Client{}

	for _, option := range options {
		option(client)
	}

	if err := client.Validate(); err != nil {
		return nil, err
	}

	// connect to the database
	conn, err := client.connect()
	if err != nil {
		return nil, err
	}

	client.Conn = conn

	// create the collections
	if err := client.createCollections(); err != nil {
		return nil, err
	}

	return client, nil
}

// Close closes the database connection
func (c *Client) Close() error {
	return c.Conn.Disconnect(context.Background())
}

// GetCollection returns a collection object by name
func (c *Client) GetCollection(name string) (*mongo.Collection, error) {
	if value, ok := c.collectionNameToCollectionObjectMap[name]; ok {
		return value, nil
	}

	return nil, fmt.Errorf("collection %s does not exist", name)
}

// GetConnection returns the database connection
func (c *Client) GetConnection() *mongo.Client {
	return c.Conn
}

// connect establishes a connection to the database
func (c *Client) connect() (*mongo.Client, error) {
	nrMon := nrmongo.NewCommandMonitor(nil)

	serverAPI := options.
		ServerAPI(options.ServerAPIVersion1).
		SetStrict(true).
		SetDeprecationErrors(true)

	client, err := mongo.Connect(
		context.Background(),
		c.
			ClientOptions.
			ApplyURI(*c.ConnectionURI).
			SetServerAPIOptions(serverAPI).
			SetMonitor(nrMon))
	if err != nil {
		return nil, err
	}

	return client, nil
}

// createCollections creates a database (mongodb) collection if it doesn't already exist
func (c *Client) createCollections() error {
	ctx := context.Background()
	if c.Conn == nil {
		return fmt.Errorf("invalid input argument. database connection: %v", c.Conn)
	}

	if c.DatabaseName == nil {
		return fmt.Errorf("invalid input argument. databaseName: %v", c.DatabaseName)
	}

	// check if collection already exists
	db := c.Conn.Database(*c.DatabaseName)
	collections, err := db.ListCollectionNames(ctx, bson.D{{}})
	if err != nil {
		return err
	}

	// create the collections
	for _, collection := range c.CollectionNames {
		// ensure the collection is not in the already existing collection set
		if !contains(collections, collection) {
			if err := db.CreateCollection(ctx, collection); err != nil {
				return err
			}

			// update the collection map
			if c.collectionNameToCollectionObjectMap == nil {
				c.collectionNameToCollectionObjectMap = make(map[string]*mongo.Collection, 0)
			}

			// add the collection to the map
			if _, ok := c.collectionNameToCollectionObjectMap[collection]; !ok {
				c.collectionNameToCollectionObjectMap[collection] = db.Collection(collection)
			}
		}
	}

	return nil
}

// The function checks if a given string is present in a collection of strings.
func contains(collectionSet []string, collection string) bool {
	if len(collectionSet) == 0 {
		return false
	}

	if collection == "" {
		return false
	}

	for _, c := range collectionSet {
		if c == collection {
			return true
		}
	}

	return false
}
