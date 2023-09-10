package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"
)

func TestClient_Delete(t *testing.T) {
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

	// write to redis
	err := client.Write(ctx, "test-key", []byte("test-value"))
	require.NoError(t, err)

	// delete from redis
	err = client.Delete(ctx, "test-key")
	require.NoError(t, err)

	// query redis again
	_, err = client.Get(ctx, "test-key")
	require.Error(t, err)
}
