# MongoDB Client Library

![MongoDB](https://img.shields.io/badge/MongoDB-4.4+-green.svg)
![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)
[![GoDoc](https://godoc.org/github.com/SolomonAIEngineering/backend-core-library?status.svg)](https://godoc.org/github.com/SolomonAIEngineering/backend-core-library)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A production-ready MongoDB client wrapper for Go applications with built-in support for connection pooling, transactions, telemetry, and testing.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Working with Collections](#working-with-collections)
- [Transactions](#transactions)
- [Error Handling](#error-handling)
- [Telemetry & Monitoring](#telemetry--monitoring)
- [Testing](#testing)
- [Best Practices](#best-practices)
- [Migration Guide](#migration-guide)
- [Contributing](#contributing)

## Features

### Core Features
- 🔄 Automatic connection management and pooling
- 📊 Comprehensive transaction support
- 📈 Built-in telemetry and monitoring
- 🔁 Configurable retry mechanisms
- ⏱️ Query timeout handling
- 🧪 In-memory testing capabilities
- 🔍 Detailed logging and debugging
- 🛡️ Connection security options

### Performance Features
- Connection pooling optimization
- Automatic retry with exponential backoff
- Query timeout management
- Efficient resource cleanup

## Installation

```bash
go get -u github.com/SolomonAIEngineering/backend-core-library/database/mongo
```

## Quick Start

```go
package main

import (
    "context"
    "time"
    
    "github.com/SolomonAIEngineering/backend-core-library/database/mongo"
    "github.com/SolomonAIEngineering/backend-core-library/instrumentation"
    "go.uber.org/zap"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // Initialize client
    client, err := mongo.New(
        mongo.WithDatabaseName("myapp"),
        mongo.WithConnectionURI("mongodb://localhost:27017"),
        mongo.WithQueryTimeout(30 * time.Second),
        mongo.WithLogger(zap.L()),
    )
    if err != nil {
        panic(err)
    }
    defer client.Close()

    // Basic CRUD operations
    collection, err := client.GetCollection("users")
    if err != nil {
        panic(err)
    }

    // Insert a document
    ctx := context.Background()
    user := map[string]interface{}{
        "name": "John Doe",
        "email": "john@example.com",
        "created_at": time.Now(),
    }
    
    result, err := collection.InsertOne(ctx, user)
    if err != nil {
        panic(err)
    }
}
```

## Configuration

### Available Options

```go
type Config struct {
    // Client configuration
    opts := []mongo.Option{
        // Database Configuration
        mongo.WithDatabaseName("mydb"),
        mongo.WithConnectionURI("mongodb://localhost:27017"),
        
        // Timeout Settings
        mongo.WithQueryTimeout(60 * time.Second),
        mongo.WithRetryTimeOut(30 * time.Second),
        
        // Retry Configuration
        mongo.WithMaxConnectionAttempts(3),
        mongo.WithMaxRetriesPerOperation(3),
        mongo.WithOperationSleepInterval(5 * time.Second),
        
        // Monitoring & Logging
        mongo.WithTelemetry(&instrumentation.Client{}),
        mongo.WithLogger(zap.L()),
        
        // Collection Management
        mongo.WithCollectionNames([]string{"users", "products", "orders"}),
        
        // Advanced Options
        mongo.WithClientOptions(options.Client().
            ApplyURI("mongodb://localhost:27017").
            SetMaxPoolSize(100).
            SetMinPoolSize(10)),
    }
}
```

### Security Configuration

```go
// TLS Configuration
tlsConfig := &tls.Config{
    // ... your TLS configuration
}

opts := []mongo.Option{
    mongo.WithClientOptions(options.Client().
        ApplyURI("mongodb://localhost:27017").
        SetTLSConfig(tlsConfig)),
    mongo.WithCredentials(options.Credential{
        Username: "user",
        Password: "pass",
        AuthSource: "admin",
    }),
}
```

## Working with Collections

### Basic CRUD Operations

```go
// Get collection
collection, err := client.GetCollection("users")
if err != nil {
    return err
}

// Insert
doc := map[string]interface{}{
    "name": "Alice",
    "age":  30,
}
result, err := collection.InsertOne(ctx, doc)

// Find
var user map[string]interface{}
err = collection.FindOne(ctx, bson.M{"name": "Alice"}).Decode(&user)

// Update
update := bson.M{"$set": bson.M{"age": 31}}
_, err = collection.UpdateOne(
    ctx,
    bson.M{"name": "Alice"},
    update,
)

// Delete
_, err = collection.DeleteOne(ctx, bson.M{"name": "Alice"})
```

### Bulk Operations

```go
// Bulk Insert
documents := []interface{}{
    bson.D{{"name", "User1"}, {"age", 25}},
    bson.D{{"name", "User2"}, {"age", 30}},
}
result, err := collection.InsertMany(ctx, documents)

// Bulk Update
updates := []mongo.WriteModel{
    mongo.NewUpdateOneModel().
        SetFilter(bson.M{"name": "User1"}).
        SetUpdate(bson.M{"$set": bson.M{"age": 26}}),
    mongo.NewUpdateOneModel().
        SetFilter(bson.M{"name": "User2"}).
        SetUpdate(bson.M{"$set": bson.M{"age": 31}}),
}
_, err = collection.BulkWrite(ctx, updates)
```

## Transactions

### Standard Transaction Example

```go
err := client.StandardTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
    // Get collections
    users, err := client.GetCollection("users")
    if err != nil {
        return nil, err
    }
    orders, err := client.GetCollection("orders")
    if err != nil {
        return nil, err
    }

    // Perform operations
    _, err = users.InsertOne(sessCtx, userDoc)
    if err != nil {
        return nil, err
    }
    
    _, err = orders.InsertOne(sessCtx, orderDoc)
    if err != nil {
        return nil, err
    }

    return nil, nil
})
```

### Complex Transaction with Return Value

```go
type TransactionResult struct {
    UserID  string
    OrderID string
}

result, err := client.ComplexTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
    // Your transaction logic here
    userResult, err := users.InsertOne(sessCtx, userDoc)
    if err != nil {
        return nil, err
    }
    
    orderResult, err := orders.InsertOne(sessCtx, orderDoc)
    if err != nil {
        return nil, err
    }

    return &TransactionResult{
        UserID:  userResult.InsertedID.(string),
        OrderID: orderResult.InsertedID.(string),
    }, nil
})
```

## Error Handling

```go
// Custom error types
var (
    ErrConnectionFailed = errors.New("failed to connect to mongodb")
    ErrQueryTimeout     = errors.New("query timeout")
)

// Error handling example
collection, err := client.GetCollection("users")
if err != nil {
    switch {
    case errors.Is(err, mongo.ErrCollectionNotFound):
        // Handle collection not found
    case errors.Is(err, mongo.ErrDatabaseNotFound):
        // Handle database not found
    default:
        // Handle other errors
    }
    return err
}
```

## Telemetry & Monitoring

```go
// Initialize with telemetry
telemetryClient := &instrumentation.Client{
    // Configure your telemetry client
}

client, err := mongo.New(
    mongo.WithTelemetry(telemetryClient),
    mongo.WithMetricsEnabled(true),
)

// Access metrics
metrics := client.GetMetrics()
fmt.Printf("Total Queries: %d\n", metrics.TotalQueries)
fmt.Printf("Failed Queries: %d\n", metrics.FailedQueries)
fmt.Printf("Average Query Time: %v\n", metrics.AverageQueryTime)
```

## Testing

### In-Memory Testing

```go
func TestUserRepository(t *testing.T) {
    // Create test client
    testClient, err := mongo.NewInMemoryTestDbClient(
        []string{"users", "orders"},
    )
    if err != nil {
        t.Fatal(err)
    }
    defer testClient.Teardown()

    // Run tests
    t.Run("InsertUser", func(t *testing.T) {
        collection, err := testClient.GetCollection("users")
        if err != nil {
            t.Fatal(err)
        }

        user := bson.M{"name": "Test User"}
        _, err = collection.InsertOne(context.Background(), user)
        assert.NoError(t, err)
    })
}
```

### Mocking

```go
type MockMongoClient struct {
    mongo.Client
    MockInsertOne func(context.Context, interface{}) (*mongo.InsertOneResult, error)
}

func (m *MockMongoClient) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
    return m.MockInsertOne(ctx, document)
}
```

## Best Practices

### Connection Management
```go
// Initialize once at application startup
client, err := mongo.New(opts...)
if err != nil {
    log.Fatal(err)
}
defer client.Close()

// Reuse the client throughout your application
```

### Query Optimization
```go
// Use appropriate indexes
collection.CreateIndex(ctx, mongo.IndexModel{
    Keys: bson.D{{"field", 1}},
    Options: options.Index().SetUnique(true),
})

// Use projection to limit returned fields
collection.FindOne(ctx, filter, options.FindOne().SetProjection(bson.M{
    "needed_field": 1,
    "_id": 0,
}))
```

### Resource Management
```go
// Use context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// Use cursor properly
cursor, err := collection.Find(ctx, filter)
if err != nil {
    return err
}
defer cursor.Close(ctx)
```

## Migration Guide

### Upgrading from v1.x to v2.x

```go
// Old v1.x configuration
oldClient := mongo.NewClient(
    "mongodb://localhost:27017",
    "mydb",
)

// New v2.x configuration
newClient, err := mongo.New(
    mongo.WithConnectionURI("mongodb://localhost:27017"),
    mongo.WithDatabaseName("mydb"),
)
```

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/SolomonAIEngineering/backend-core-library.git

# Install dependencies
go mod download

# Run tests
go test -v ./...

# Run linting
golangci-lint run
```

## License

This MongoDB client is released under the MIT License. See [LICENSE](LICENSE) for details.

---

## Support

- 📚 [Documentation](https://github.com/SolomonAIEngineering/backend-core-library)
- 💬 [GitHub Discussions](https://github.com/SolomonAIEngineering/backend-core-library/discussions)
- 🐛 [Issue Tracker](https://github.com/SolomonAIEngineering/backend-core-library/issues)