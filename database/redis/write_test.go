package redis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Write(t *testing.T) {
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
	err := client.Write(ctx, key, value)
	assert.NoError(t, err)

	// Read value from Redis
	conn := pool.Get()
	defer conn.Close()
	result, err := redis.String(conn.Do("GET", key))
	assert.NoError(t, err)
	assert.Equal(t, "test-value", result)
}

func TestClient_WriteMany(t *testing.T) {
	// create a mock context
	ctx := context.Background()

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

	// create a mock pair of key-value pairs to write to the cache
	pairs := map[string][]byte{
		"key1": []byte("value1"),
		"key2": []byte("value2"),
	}

	// call the WriteMany function with the mock context and pairs
	err := client.WriteMany(ctx, pairs)

	// check that no error occurred
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	// check that the values were written to the cache
	for key, value := range pairs {
		cachedValue, err := client.Get(ctx, key)
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}
		if !bytes.Equal(cachedValue, value) {
			t.Errorf("expected cached value to be %v, but got %v", value, cachedValue)
		}
	}
}

func TestClient_WriteAny(t *testing.T) {
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

	// Set up test data
	testKey := "testkey"
	testValue := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}

	// Call WriteAny to write the test value to Redis
	err := client.WriteAny(context.Background(), testKey, testValue)
	require.NoError(t, err)

	// Verify that the value was written correctly
	conn := client.pool.Get()
	defer conn.Close()
	rawValue, err := redis.Bytes(conn.Do("GET", testKey))
	require.NoError(t, err)

	// Unmarshal the raw value to a map to compare with the test value
	var retrievedValue map[string]interface{}
	err = json.Unmarshal(rawValue, &retrievedValue)
	require.NoError(t, err)
	require.Equal(t, len(getMapKeys(testValue)), len(getMapKeys(retrievedValue)))
}

func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
