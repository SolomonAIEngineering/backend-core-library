package cacher // import "github.com/SolomonAIEngineering/backend-core-library/cacher"

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

// Client implements a Redis cache client with connection pooling and instrumentation support.
// It automatically handles connection management, key prefixing, and tracing of Redis operations.
type Client struct {
	logger                *zap.Logger             // Logger for operational logging
	pool                  *redis.Pool             // Connection pool for Redis
	serviceName           string                  // Service name used for key prefixing
	instrumentationClient *instrumentation.Client // Optional instrumentation for tracing
	cacheTTLInSeconds     int                     // Default TTL for cache entries
}

// New creates a new Redis cache client with the provided options.
// Example usage:
//
//	client, err := cacher.New(
//	    cacher.WithLogger(logger),
//	    cacher.WithRedisConn(pool),
//	    cacher.WithServiceName("myservice"),
//	    cacher.WithCacheTTLInSeconds(3600),
//	)
func New(opts ...Option) (*Client, error) {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

// WriteToCache writes a value to the cache with the configured TTL.
// The key will be automatically prefixed with the service name.
// Example usage:
//
//	err := client.WriteToCache(ctx, "user:123", []byte(`{"name":"John"}`))
func (s *Client) WriteToCache(ctx context.Context, key string, value []byte) error {
	if s.instrumentationClient != nil {
		txn := s.instrumentationClient.GetTraceFromContext(ctx)
		span := s.instrumentationClient.StartRedisDatastoreSegment(txn, "redis-write-to-cache")
		defer span.End()
	}

	// validate the key
	if key == "" {
		return fmt.Errorf("empty key")
	}

	prefixedCacheKey := fmt.Sprintf("%s:%s", s.serviceName, key)

	conn := s.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", prefixedCacheKey, string(value), "EX", s.cacheTTLInSeconds); err != nil {
		return err
	}

	return nil
}

// WriteManyToCache writes multiple key-value pairs to the cache atomically.
// This is more efficient than multiple individual writes for bulk operations.
// Example usage:
//
//	pairs := map[string][]byte{
//	    "user:123": []byte(`{"name":"John"}`),
//	    "user:456": []byte(`{"name":"Jane"}`),
//	}
//	err := client.WriteManyToCache(ctx, pairs)
func (s *Client) WriteManyToCache(ctx context.Context, pairs map[string][]byte) error {
	if s.instrumentationClient != nil {
		txn := s.instrumentationClient.GetTraceFromContext(ctx)
		span := s.instrumentationClient.StartRedisDatastoreSegment(txn, "redis-write-many-to-cache")
		defer span.End()
	}

	// validate the key
	if len(pairs) == 0 {
		return fmt.Errorf("empty cache reference set")
	}

	conn := s.pool.Get()
	defer conn.Close()

	// Use the redis.Args helper to create the arguments for the MSET command
	args := redis.Args{}.AddFlat(pairs)
	if _, err := conn.Do("MSET", args...); err != nil {
		return err
	}

	return nil
}

// GetFromCache retrieves a value from the cache by key.
// Returns nil and an error if the key doesn't exist.
// Example usage:
//
//	data, err := client.GetFromCache(ctx, "user:123")
//	if err != nil {
//	    if err == redis.ErrNil {
//	        // Handle cache miss
//	    }
//	    return err
//	}
func (s *Client) GetFromCache(ctx context.Context, key string) ([]byte, error) {
	if s.instrumentationClient != nil {
		txn := s.instrumentationClient.GetTraceFromContext(ctx)
		span := s.instrumentationClient.StartRedisDatastoreSegment(txn, "redis-read-from-cache")
		defer span.End()
	}

	// validate the key
	if key == "" {
		return nil, fmt.Errorf("empty key")
	}

	conn := s.pool.Get()
	defer conn.Close()

	value, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return value, nil
}

// GetManyFromCache retrieves multiple values from the cache in a single operation.
// If any key is missing, it returns an error. For partial success/failure handling,
// use individual GetFromCache calls.
// Example usage:
//
//	keys := []string{"user:123", "user:456"}
//	values, err := client.GetManyFromCache(ctx, keys)
func (s *Client) GetManyFromCache(ctx context.Context, keys []string) ([][]byte, error) {
	isNilByteSlice := func(b []byte) bool {
		return b == nil
	}

	if s.instrumentationClient != nil {
		txn := s.instrumentationClient.GetTraceFromContext(ctx)
		span := s.instrumentationClient.StartRedisDatastoreSegment(txn, "redis-read-many-from-cache")
		defer span.End()
	}

	// validate the key
	if len(keys) == 0 {
		return nil, fmt.Errorf("empty key")
	}

	args := redis.Args{}.AddFlat(keys)
	conn := s.pool.Get()
	defer conn.Close()

	values, err := redis.ByteSlices(conn.Do("MGET", args...))
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		// this occurs if there was no matching value for the query
		if isNilByteSlice(value) {
			return nil, fmt.Errorf("nil value")
		}
	}

	return values, nil
}

// DeleteFromCache removes a value from the cache.
// It's safe to delete non-existent keys.
// Example usage:
//
//	err := client.DeleteFromCache(ctx, "user:123")
func (s *Client) DeleteFromCache(ctx context.Context, key string) error {
	if s.instrumentationClient != nil {
		txn := s.instrumentationClient.GetTraceFromContext(ctx)
		span := s.instrumentationClient.StartRedisDatastoreSegment(txn, "redis-delete-from-cache")
		defer span.End()
	}

	// validate the key
	if key == "" {
		return fmt.Errorf("empty key")
	}

	conn := s.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", key); err != nil {
		return err
	}

	return nil
}

// WriteAnyToCache marshals any JSON-serializable value and writes it to the cache.
// This is a convenience wrapper around WriteToCache with JSON marshaling.
// Example usage:
//
//	type User struct {
//	    Name string `json:"name"`
//	    Age  int    `json:"age"`
//	}
//	user := User{Name: "John", Age: 30}
//	err := client.WriteAnyToCache(ctx, "user:123", user)
func (s *Client) WriteAnyToCache(ctx context.Context, key string, value interface{}) error {
	if s.instrumentationClient != nil {
		txn := s.instrumentationClient.GetTraceFromContext(ctx)
		span := s.instrumentationClient.StartRedisDatastoreSegment(txn, "redis-write-any-to-cache")
		defer span.End()
	}

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := s.WriteToCache(ctx, key, data); err != nil {
		return err
	}

	return nil
}

// WriteToCacheWithTTL writes a value to the cache with a custom TTL.
// If timeToLiveInSeconds is <= 0, defaults to 60 seconds.
// Example usage:
//
//	// Cache for 5 minutes
//	err := client.WriteToCacheWithTTL(ctx, "temp:123", data, 300)
func (s *Client) WriteToCacheWithTTL(ctx context.Context, key string, value []byte, timeToLiveInSeconds int) error {
	if timeToLiveInSeconds <= 0 {
		timeToLiveInSeconds = 60
	}

	if s.instrumentationClient != nil {
		txn := s.instrumentationClient.GetTraceFromContext(ctx)
		span := s.instrumentationClient.StartRedisDatastoreSegment(txn, "redis-write-to-cache-with-ttl")
		defer span.End()
	}

	// validate the key
	if key == "" {
		return fmt.Errorf("empty key")
	}

	prefixedCacheKey := fmt.Sprintf("%s:%s", s.serviceName, key)

	conn := s.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", prefixedCacheKey, string(value), "EX", timeToLiveInSeconds); err != nil {
		return err
	}

	return nil
}
