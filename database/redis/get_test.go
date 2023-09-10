package redis

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestClient_Get(t *testing.T) {
	redisTestServer := miniredis.RunT(t)
	defer redisTestServer.Close()

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf(":%s", redisTestServer.Port()))
		},
	}

	client := &Client{
		pool:              pool,
		cacheTTLInSeconds: 60,
	}

	ctx := context.Background()
	key := "test-key"
	value := []byte("test-value")

	// Write value to Redis
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	assert.NoError(t, err)

	// Read value from Redis
	result, err := client.Get(ctx, key)
	assert.NoError(t, err)
	assert.Equal(t, value, result)
}

func TestClient_GetMany(t *testing.T) {
	redisTestServer := miniredis.RunT(t)
	defer redisTestServer.Close()

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf(":%s", redisTestServer.Port()))
		},
	}

	client := &Client{
		pool:              pool,
		cacheTTLInSeconds: 60,
	}

	// Define a list of keys to retrieve
	keys := []string{"key1", "key2", "key3"}

	// Add test data to Redis
	err := client.Write(context.Background(), "key1", []byte("value1"))
	if err != nil {
		t.Fatalf("Failed to write test data to Redis: %v", err)
	}
	err = client.Write(context.Background(), "key2", []byte("value2"))
	if err != nil {
		t.Fatalf("Failed to write test data to Redis: %v", err)
	}
	err = client.Write(context.Background(), "key3", []byte("value3"))
	if err != nil {
		t.Fatalf("Failed to write test data to Redis: %v", err)
	}

	// Retrieve values from Redis
	values, err := client.GetMany(context.Background(), keys)
	if err != nil {
		t.Fatalf("Failed to retrieve values from Redis: %v", err)
	}

	// Verify that the values were retrieved correctly
	if len(values) != len(keys) {
		t.Errorf("Expected %d values, but got %d", len(keys), len(values))
	}

	for i, value := range values {
		expectedValue := []byte(fmt.Sprintf("value%d", i+1))
		if !bytes.Equal(value, expectedValue) {
			t.Errorf("Expected value %v for key %v, but got %v", expectedValue, keys[i], value)
		}
	}

	// Clean up test data
	for _, key := range keys {
		err = client.Delete(context.Background(), key)
		if err != nil {
			t.Fatalf("Failed to clean up test data from Redis: %v", err)
		}
	}

}
