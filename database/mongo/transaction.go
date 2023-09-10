package mongo // import "github.com/SimifiniiCTO/simfiny-core-lib/database/mongo"

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoTx is a type alias for the `WithTransaction` method of the `mongo.Session` object.
type MongoTx func(sessCtx mongo.SessionContext) (any, error)

// ComplexTransaction is a wrapper around the `WithTransaction` method of the `mongo.Session` object.
func (c *Client) ComplexTransaction(ctx context.Context, callback MongoTx) (any, error) {
	session, err := c.Conn.StartSession()
	if err != nil {
		return nil, fmt.Errorf("failed creating session | %s", err.Error())
	}

	defer session.EndSession(ctx)

	res, err := session.WithTransaction(ctx, callback)
	if err != nil {
		return nil, fmt.Errorf("failed executing transaction | %s", err.Error())
	}

	return res, nil
}

// StandardTransaction is a wrapper around the `WithTransaction` method of the `mongo.Session` object.
func (c *Client) StandardTransaction(ctx context.Context, callback MongoTx) error {
	session, err := c.Conn.StartSession()
	if err != nil {
		return fmt.Errorf("failed creating session | %s", err.Error())
	}

	defer session.EndSession(ctx)

	if _, err = session.WithTransaction(ctx, callback); err != nil {
		return fmt.Errorf("failed executing transaction | %s", err.Error())
	}

	return nil
}
