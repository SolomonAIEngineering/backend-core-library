package redis // import "github.com/SimifiniiCTO/simfiny-core-lib/database/redis"

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// Get reads a value from the cache
func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartRedisDatastoreSegment(txn, RedisReadFromCacheTxn.String())
	defer span.End()

	// validate the key
	if key == "" {
		return nil, fmt.Errorf("empty key")
	}

	conn := c.pool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

// GetManyFromCache reads a value from the cache
func (c *Client) GetMany(ctx context.Context, keys []string) ([][]byte, error) {
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartRedisDatastoreSegment(txn, RedisReadManyFromCacheTxn.String())
	defer span.End()

	// validate the key
	if len(keys) == 0 {
		return nil, fmt.Errorf("empty key")
	}

	conn := c.pool.Get()
	defer conn.Close()

	var ifaceSlice []interface{}
	for _, s := range keys {
		ifaceSlice = append(ifaceSlice, s)
	}

	values, err := redis.ByteSlices(conn.Do("MGET", ifaceSlice...))
	if err != nil {
		return nil, err
	}

	return values, nil
}

// Get reads a value from the cache
func (c *Client) Exists(ctx context.Context, key string) (bool, error) {
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartRedisDatastoreSegment(txn, RedisReadFromCacheTxn.String())
	defer span.End()

	// validate the key
	if key == "" {
		return false, fmt.Errorf("empty key")
	}

	conn := c.pool.Get()
	defer conn.Close()
	value, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	return value, nil
}
