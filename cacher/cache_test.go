package cacher

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/alicebob/miniredis"
	redigoredis "github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var (
	redisServer    *miniredis.Miniredis
	mockTestClient = &Client{
		logger:                zap.L(),
		pool:                  setupRedis(),
		serviceName:           "test-service",
		instrumentationClient: &instrumentation.Client{},
		cacheTTLInSeconds:     60,
	}
)

const (
	EMPTY         = ""
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func TestServer_WriteToCache(t *testing.T) {
	type args struct {
		ctx   context.Context
		key   string
		value []byte
	}
	tests := []struct {
		name    string
		s       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - write to cache",
			s:    mockTestClient,
			args: args{
				ctx:   context.Background(),
				key:   "key",
				value: []byte("value"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.WriteToCache(tt.args.ctx, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Server.WriteToCache() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				// attempt to get from cache
				key := fmt.Sprintf("%s:%s", mockTestClient.serviceName, tt.args.key)
				bytes, err := tt.s.GetFromCache(tt.args.ctx, key)
				assert.NoError(t, err)
				assert.True(t, reflect.DeepEqual(tt.args.value, bytes))
			}
		})
	}
}

func TestServer_WriteManyToCache(t *testing.T) {
	type args struct {
		ctx   context.Context
		pairs map[string][]byte
	}
	tests := []struct {
		name    string
		s       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - write many to cache",
			s:    mockTestClient,
			args: args{
				ctx: context.Background(),
				pairs: map[string][]byte{
					"key1": []byte("value1"),
					"key2": []byte("value2"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.WriteManyToCache(tt.args.ctx, tt.args.pairs); (err != nil) != tt.wantErr {
				t.Errorf("Server.WriteManyToCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_GetFromCache(t *testing.T) {
	type args struct {
		ctx          context.Context
		key          string
		precondition func(ctx context.Context, t *testing.T, keys string)
	}
	tests := []struct {
		name    string
		s       *Client
		args    args
		wantErr bool
	}{
		{
			name: "get from cache",
			s:    mockTestClient,
			args: args{
				ctx: context.Background(),
				key: "key1",
				precondition: func(ctx context.Context, t *testing.T, key string) {
					conn := mockTestClient.pool.Get()
					defer conn.Close()

					randomStringValue := generateRandomString(10)
					_, err := conn.Do("SET", key, randomStringValue)
					assert.NoError(t, err)
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.precondition(tt.args.ctx, t, tt.args.key)
			got, err := tt.s.GetFromCache(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetFromCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestServer_GetManyFromCache(t *testing.T) {
	type args struct {
		ctx          context.Context
		keys         []string
		precondition func(ctx context.Context, t *testing.T, keys []string)
	}
	tests := []struct {
		name    string
		s       *Client
		args    args
		want    [][]byte
		wantErr bool
	}{
		{
			name: "pass - get a valid key",
			s:    mockTestClient,
			args: args{
				ctx:  context.Background(),
				keys: []string{"key1", "key2", "key3"},
				precondition: func(ctx context.Context, t *testing.T, keys []string) {
					conn := mockTestClient.pool.Get()
					defer conn.Close()

					for _, key := range keys {
						randomStringValue := generateRandomString(10)
						_, err := conn.Do("SET", key, randomStringValue)
						assert.NoError(t, err)
					}
				},
			},
			want:    [][]byte{[]byte("value1")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.precondition(tt.args.ctx, t, tt.args.keys)

			got, err := tt.s.GetManyFromCache(tt.args.ctx, tt.args.keys)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetManyFromCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, len(tt.args.keys), len(got))
			}
		})
	}
}

func TestServer_DeleteFromCache(t *testing.T) {
	type args struct {
		ctx          context.Context
		key          string
		value        string
		precondition func(ctx context.Context, t *testing.T, key string, value string)
	}
	tests := []struct {
		name    string
		s       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - delete a valid key",
			s:    mockTestClient,
			args: args{
				ctx: context.Background(),
				key: "test",
				precondition: func(ctx context.Context, t *testing.T, key string, value string) {
					err := mockTestClient.WriteToCache(ctx, key, []byte(value))
					assert.NoError(t, err)
				},
			},
			wantErr: false,
		},
		{
			name: "fail - delete a non-existent key",
			s:    mockTestClient,
			args: args{
				ctx:   context.Background(),
				key:   "test-deletion",
				value: "test-deletion",
				precondition: func(ctx context.Context, t *testing.T, key string, value string) {
				},
			},
			// no error should be present if the key does not exist
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.precondition(tt.args.ctx, t, tt.args.key, tt.args.value)

			if err := tt.s.DeleteFromCache(tt.args.ctx, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteFromCache() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				_, err := tt.s.GetFromCache(tt.args.ctx, tt.args.key)
				assert.Error(t, err)
			}
		})
	}
}

func TestServer_WriteAnyToCache(t *testing.T) {
	type args struct {
		ctx   context.Context
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		s       *Client
		args    args
		wantErr bool
	}{
		{
			name: "pass - write a sample string to cache",
			s:    mockTestClient,
			args: args{
				ctx:   context.Background(),
				key:   "test1",
				value: "test1",
			},
			wantErr: false,
		},
		{
			name: "pass - write a sample struct to cache",
			s:    mockTestClient,
			args: args{
				ctx: context.Background(),
				key: "test2",
				value: struct {
					Name string
					Age  int
				}{
					Name: "test2",
					Age:  10,
				},
			},
			wantErr: false,
		},
		{
			name: "pass - write a sample map to cache",
			s:    mockTestClient,
			args: args{
				ctx: context.Background(),
				key: "test3",
				value: map[string]interface{}{
					"name": "test3",
					"age":  10,
				},
			},
			wantErr: false,
		},
		{
			name: "pass - write a sample slice to cache",
			s:    mockTestClient,
			args: args{
				ctx: context.Background(),
				key: "test4",
				value: []interface{}{
					"name", "test4",
					"age", 10,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.WriteAnyToCache(tt.args.ctx, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Server.WriteAnyToCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// mockRedis returns a mock redis server.
func setupRedis() *redigoredis.Pool {
	redisServer = mockRedis()

	return &redigoredis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        getRedisCacheConn,
		TestOnBorrow: func(c redigoredis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// getRedisCacheConn returns a redis connection.
func getRedisCacheConn() (redigoredis.Conn, error) {
	redisUrl, err := url.Parse(getRedisConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %v", err)
	}

	var opts []redigoredis.DialOption
	if user := redisUrl.User; user != nil {
		opts = append(opts, redigoredis.DialUsername(user.Username()))
		if password, ok := user.Password(); ok {
			opts = append(opts, redigoredis.DialPassword(password))
		}
	}

	return redigoredis.Dial("tcp", redisUrl.Host, opts...)
}

// mockRedis returns a mock redis server.
func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	return s
}

func getRedisConnectionString() string {
	return fmt.Sprintf("redis://:@%s", redisServer.Addr())
}

// GenerateRandomString generates a random string based on the size specified by the client
func generateRandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
