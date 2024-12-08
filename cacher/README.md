<div align="center">
    <h1 align="center">Solomon AI Cache Library</h1>
    <h3 align="center">Efficient and Type-safe Redis-based Caching for Go Applications</h3>
</div>

<div align="center">
  
[![Go Report Card](https://goreportcard.com/badge/github.com/SolomonAIEngineering/backend-core-library)](https://goreportcard.com/report/github.com/SolomonAIEngineering/backend-core-library)
[![GoDoc](https://godoc.org/github.com/SolomonAIEngineering/backend-core-library?status.svg)](https://godoc.org/github.com/SolomonAIEngineering/backend-core-library)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

- Connection pooling for optimal performance
- Automatic key prefixing to prevent collisions
- Distributed tracing integration
- Configurable cache TTLs
- Bulk operations support
- JSON serialization helpers
- Comprehensive error handling

## Installation

```go
package main

import (
    "context"
    "log"
    
    "github.com/SolomonAIEngineering/backend-core-library/cacher"
    "github.com/gomodule/redigo/redis"
    "go.uber.org/zap"
)

func main() {
    // Initialize Redis pool with recommended settings
    pool := &redis.Pool{
        MaxIdle:     80,
        MaxActive:   12000,
        IdleTimeout: 240 * time.Second,
        Wait:        true,
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
        Dial: func() (redis.Conn, error) {
            return redis.Dial("tcp", "localhost:6379")
        },
    }

    // Create cache client with all required options
    client, err := cacher.New(
        cacher.WithRedisConn(pool),
        cacher.WithLogger(zap.L()),
        cacher.WithServiceName("my-service"),
        cacher.WithCacheTTLInSeconds(3600),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()
    
    // Basic write operation
    if err := client.WriteToCache(ctx, "user:123", []byte("data")); err != nil {
        log.Printf("Failed to write to cache: %v", err)
    }
}
```

## Advanced Usage

### Type-Safe Cache Keys
The library provides a type-safe way to manage cache keys:

```go
// Define type-safe cache keys
key := cacher.NewCacheKey("user")
userKey := key.Enrich("123") // Results in "user:123"
```

### Batch Operations with Pipeline
Efficiently handle multiple cache operations:

```go
// Write multiple values atomically
data := map[string][]byte{
    "key1": []byte("value1"),
    "key2": []byte("value2"),
}
err := client.WriteManyToCache(ctx, data)

// Read multiple values
values, err := client.GetManyFromCache(ctx, []string{"key1", "key2"})
```

### Custom TTL Support
Set custom TTL for specific cache entries:

```go
err := client.WriteToCacheWithTTL(ctx, "temp-key", []byte("data"), 300) // 5 minutes TTL
```

## Configuration Options
| Option | Description | Required | Default |
|--------|-------------|----------|---------|
| WithRedisConn | Redis connection pool | Yes | - |
| WithLogger | Zap logger instance | Yes | - |
| WithServiceName | Service name for key prefixing | Yes | - |
| WithCacheTTLInSeconds | Default TTL for cache entries | Yes | - |
| WithIntrumentationClient | Instrumentation client | No | nil |

### Best Practices
Connection Pool Management

```go
pool := &redis.Pool{
    MaxIdle:     80,
    MaxActive:   12000,
    IdleTimeout: 240 * time.Second,
    Wait:        true,
    TestOnBorrow: func(c redis.Conn, t time.Time) error {
        _, err := c.Do("PING")
        return err
    },
}
defer pool.Close()
```

### Error Handling
```go
// Always check for errors and implement retries for critical operations
value, err := client.GetFromCache(ctx, key)
if err != nil {
    if errors.Is(err, redis.ErrNil) {
        // Handle cache miss
    } else {
        // Handle other errors
    }
}
```

### Key Management
```go
// Use service name for automatic key prefixing
client, _ := cacher.New(
    cacher.WithServiceName("user-service"),
    // ... other options
)
```

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📚 [Documentation](https://github.com/SolomonAIEngineering/backend-core-library)
- 💬 [GitHub Discussions](https://github.com/SolomonAIEngineering/backend-core-library/discussions)
- 🐛 [Issue Tracker](https://github.com/SolomonAIEngineering/backend-core-library/issues)

---

<div align="center">
    <sub>Built with ❤️ by Solomon AI Engineering</sub>
</div>