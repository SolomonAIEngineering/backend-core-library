package mongo // import "github.com/SolomonAIEngineering/backend-core-library/database/mongo"

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// StartDbSegment starts a new
// `newrelic.DatastoreSegment` for database operations. It takes in a context, a name for the
// operation, and a collection name as parameters. If telemetry is enabled, it retrieves the
// transaction from the context and starts a new datastore segment for the operation using the
// `StartNosqlDatastoreSegment` method from the `Telemetry` object. It returns the newly created
// segment. If telemetry is not enabled, it returns `nil`.
func (c *Client) StartDbSegment(ctx context.Context, name, collectionName string) *newrelic.DatastoreSegment {
	if c.Telemetry != nil {
		txn := c.Telemetry.GetTraceFromContext(ctx)
		span := c.Telemetry.StartNosqlDatastoreSegment(txn, name, collectionName)
		return span
	}

	return nil
}
