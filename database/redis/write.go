package redis // import "github.com/SolomonAIEngineering/backend-core-library/database/redis"

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// WriteWithTTL writes a value to the cache with time to live
func (c *Client) WriteWithTTL(ctx context.Context, key string, value []byte, cacheTTLInSeconds int) error {
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartRedisDatastoreSegment(txn, RedisWriteToCacheWithTTLTxn.String())
	defer span.End()

	// validate the key
	if key == "" {
		return fmt.Errorf("empty key")
	}

	conn := c.pool.Get()
	defer conn.Close()
	if _, err := conn.Do("SET", key, string(value), "EX", cacheTTLInSeconds); err != nil {
		return err
	}

	return nil
}

// WriteToCache writes a value to the cache
func (c *Client) Write(ctx context.Context, key string, value []byte) error {
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartRedisDatastoreSegment(txn, RedisWriteToCacheTxn.String())
	defer span.End()

	// validate the key
	if key == "" {
		return fmt.Errorf("empty key")
	}

	conn := c.pool.Get()
	defer conn.Close()
	if _, err := conn.Do("SET", key, string(value), "EX", c.cacheTTLInSeconds); err != nil {
		return err
	}

	return nil
}

// WriteMany writes a many values to the cache
func (c *Client) WriteMany(ctx context.Context, pairs map[string][]byte) error {
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartRedisDatastoreSegment(txn, RedisWriteToCacheTxn.String())
	defer span.End()

	// validate the key
	if len(pairs) == 0 {
		return fmt.Errorf("empty cache reference set")
	}

	conn := c.pool.Get()
	defer conn.Close()

	// Use the redis.Args helper to create the arguments for the MSET command
	args := redis.Args{}.AddFlat(pairs)
	if _, err := conn.Do("MSET", args...); err != nil {
		return err
	}

	return nil
}

// WriteAny writes a value to the cache
func (c *Client) WriteAny(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := c.Write(ctx, key, data); err != nil {
		return err
	}

	return nil
}
