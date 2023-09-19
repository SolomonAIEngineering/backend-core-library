package redis // import "github.com/SolomonAIEngineering/backend-core-library/database/redis"

import (
	"context"
	"fmt"
)

// Delete deletes a value from the cache
func (c *Client) Delete(ctx context.Context, key string) error {
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartRedisDatastoreSegment(txn, RedisDeleteFromCacheTxn.String())
	defer span.End()

	// validate the key
	if key == "" {
		return fmt.Errorf("empty key")
	}

	conn := c.pool.Get()
	defer conn.Close()
	if _, err := conn.Do("DEL", key); err != nil {
		return err
	}

	return nil
}
