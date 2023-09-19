// Package rediswrapper provides a simple Redis client wrapper to easily interface with Redis server for caching and storage purposes.
//
// The Redis client implementation is based on the popular "github.com/go-redis/redis/v8" package, and provides simple functions for key-value set, get, delete, and hash operations.
//
// To use the Redis client, simply create a new client using the NewRedisClient function, passing in the Redis server host and port, as well as any authentication details if necessary.
//
// The Redis client is thread-safe and supports connection pooling to minimize overhead when making frequent connections to the Redis server.
//
// Example usage:
//
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/example/rediswrapper"
//
// )
//
//	func main() {
//	   // Create a new Redis client
//	   client := rediswrapper.NewRedisClient("localhost:6379", "", 0)
//
//	   // Set a value in Redis
//	   err := client.SetValue("key", "value", 0)
//	   if err != nil {
//	       fmt.Println("Error setting value:", err)
//	       return
//	   }
//
//	   // Get a value from Redis
//	   value, err := client.GetValue("key")
//	   if err != nil {
//	       fmt.Println("Error getting value:", err)
//	       return
//	   }
//	   fmt.Println("Value:", value)
//	}
package redis // import "github.com/SimifiniiCTO/simfiny-core-lib/database/redis"
