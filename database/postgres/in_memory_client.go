package postgres // import "github.com/SolomonAIEngineering/backend-core-library/database/postgres"

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewInMemoryTestDbClient creates a new in memory test db client
// This is useful only for unit tests
func NewInMemoryTestDbClient(models ...any) (*Client, error) {
	var (
		mockQueryTimeout              = 10 * time.Minute
		mockMaxAttempts               = 3
		mockMaxConnectionRetryTimeout = 10 * time.Minute
		mockRetrySleep                = 10 * time.Second
		once                          sync.Once
		testdb                        *gorm.DB
		testDatabaseName              = "gen_test.db"
	)

	// initialize db with sqlite
	once.Do(func() {
		var err error
		testdb, err = gorm.Open(sqlite.Open(testDatabaseName), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("open sqlite %q fail: %w", testDatabaseName, err))
		}
	})

	if err := testdb.AutoMigrate(models...); err != nil {
		return nil, err
	}

	return &Client{
		Engine:                    testdb,
		QueryTimeout:              &mockQueryTimeout,
		MaxConnectionRetries:      &mockMaxAttempts,
		MaxConnectionRetryTimeout: &mockMaxConnectionRetryTimeout,
		RetrySleep:                &mockRetrySleep,
		InstrumentationClient:     &instrumentation.Client{},
	}, nil
}

// TestTxCleanupHandlerForUnitTests is a handler that can be used to rollback a transaction to a save point.
// This is useful for unit tests that need to rollback a transaction to a save point.
type TestTxCleanupHandlerForUnitTests struct {
	cancelFunc               context.CancelFunc
	Tx                       *gorm.DB
	savePointRollbackHandler func(tx *gorm.DB)
}

// ConfigureNewTxCleanupHandlerForUnitTests creates a new transaction with a save point and returns a handler that can be used to rollback to the save point.
// This is useful for unit tests that need to rollback a transaction to a save point.
func (c *Client) ConfigureNewTxCleanupHandlerForUnitTests() *TestTxCleanupHandlerForUnitTests {
	const SAVE_POINT = "test_save_point"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	tx := c.Engine.WithContext(ctx).Begin()
	tx.SavePoint(SAVE_POINT)

	return &TestTxCleanupHandlerForUnitTests{
		cancelFunc: cancel,
		Tx:         tx,
		savePointRollbackHandler: func(tx *gorm.DB) {
			tx.RollbackTo(SAVE_POINT)
			tx.Commit()
		},
	}
}
