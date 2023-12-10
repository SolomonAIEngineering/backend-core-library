package cacher

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

type Client struct {
	logger                *zap.Logger
	pool                  *redis.Pool
	serviceName           string
	instrumentationClient *instrumentation.Client
	cacheTTLInSeconds     int
}

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

// WriteToCache writes a value to the cache
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

// WriteManyToCache writes many values to the cache
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

// GetFromCache reads a value from the cache
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

// GetManyFromCache retrieves multiple values from the cache based on the provided keys.
// It returns a slice of byte slices representing the retrieved values, or an error if any occurred.
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

// DeleteFromCache deletes a value from the cache
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

// WriteAnyToCache writes the given value to the cache using the specified key.
// The value is first marshaled into JSON format before being written to the cache.
// If an error occurs during marshaling or writing to the cache, it is returned.
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

// WriteToCacheWithTTL writes a value to the cache with a defined ttl in seconds
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
