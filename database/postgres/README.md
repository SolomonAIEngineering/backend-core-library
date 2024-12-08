# Backend Core Library - PostgreSQL Client

![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14%2B-blue.svg)
![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue.svg)
[![GoDoc](https://godoc.org/github.com/SolomonAIEngineering/backend-core-library?status.svg)](https://godoc.org/github.com/SolomonAIEngineering/backend-core-library)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

A production-ready PostgreSQL client for Go applications with built-in support for connection pooling, transactions, GORM integration, and comprehensive testing utilities.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Connection Management](#connection-management)
- [Transaction Handling](#transaction-handling)
- [GORM Integration](#gorm-integration)
- [Error Handling](#error-handling)
- [Testing](#testing)
- [Monitoring & Instrumentation](#monitoring--instrumentation)
- [Best Practices](#best-practices)
- [Migration Guide](#migration-guide)
- [API Reference](#api-reference)
- [Contributing](#contributing)

## Features

### Core Features
- 🔄 Smart connection pooling with automatic management
- 📊 Comprehensive transaction support with rollback
- 🔌 Seamless GORM integration
- 📈 Built-in instrumentation and monitoring
- 🔁 Configurable retry mechanisms
- ⏱️ Query timeout handling
- 🧪 In-memory testing capabilities
- 🔍 Detailed logging and debugging
- 🛡️ Connection security options

### Performance Features
- Connection pool optimization
- Automatic retry with exponential backoff
- Query timeout management
- Efficient resource cleanup
- Statement caching
- Connection lifecycle management

## Installation

```bash
go get -u github.com/SolomonAIEngineering/backend-core-library/database/postgres
```

## Quick Start

```go
package main

import (
    "context"
    "time"
    "log"

    "github.com/SolomonAIEngineering/backend-core-library/database/postgres"
    "github.com/SolomonAIEngineering/backend-core-library/instrumentation"
    "go.uber.org/zap"
)

func main() {
    // Initialize client with comprehensive configuration
    client, err := postgres.New(
        postgres.WithQueryTimeout(30 * time.Second),
        postgres.WithMaxConnectionRetries(3),
        postgres.WithConnectionString("postgresql://user:pass@localhost:5432/mydb?sslmode=verify-full"),
        postgres.WithMaxIdleConnections(10),
        postgres.WithMaxOpenConnections(100),
        postgres.WithMaxConnectionLifetime(1 * time.Hour),
        postgres.WithLogger(zap.L()),
        postgres.WithInstrumentationClient(&instrumentation.Client{}),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Your application code here
}
```

## Connection Management

### Basic Configuration

```go
type ConnectionConfig struct {
    // Connection parameters
    Host            string
    Port            int
    Database        string
    User            string
    Password        string
    SSLMode         string
    
    // Pool configuration
    MaxIdleConns    int
    MaxOpenConns    int
    ConnMaxLifetime time.Duration
    
    // Retry configuration
    MaxRetries      int
    RetryTimeout    time.Duration
    RetrySleep      time.Duration
}

// Example configuration
config := ConnectionConfig{
    Host:            "localhost",
    Port:            5432,
    Database:        "myapp",
    MaxIdleConns:    10,
    MaxOpenConns:    100,
    ConnMaxLifetime: time.Hour,
    MaxRetries:      3,
    RetryTimeout:    time.Minute,
    RetrySleep:      time.Second,
}
```

### SSL/TLS Configuration

```go
import "crypto/tls"

// Configure TLS
tlsConfig := &tls.Config{
    MinVersion: tls.VersionTLS12,
    ServerName: "your-db-host",
}

connStr := fmt.Sprintf(
    "host=%s port=%d user=%s password=%s dbname=%s sslmode=verify-full",
    config.Host, config.Port, config.User, config.Password, config.Database,
)

client, err := postgres.New(
    postgres.WithConnectionString(&connStr),
    postgres.WithTLSConfig(tlsConfig),
)
```

## Transaction Handling

### Simple Transactions

```go
// Basic transaction
err := client.PerformTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
    // Create a user
    user := User{Name: "John Doe", Email: "john@example.com"}
    if err := tx.Create(&user).Error; err != nil {
        return err
    }
    
    // Create an order for the user
    order := Order{UserID: user.ID, Amount: 100.00}
    if err := tx.Create(&order).Error; err != nil {
        return err
    }
    
    return nil
})
```

### Complex Transactions

```go
type TransactionResult struct {
    UserID  uint
    OrderID uint
}

// Transaction with return value
result, err := client.PerformComplexTransaction(ctx, func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
    user := User{Name: "Jane Doe"}
    if err := tx.Create(&user).Error; err != nil {
        return nil, err
    }
    
    order := Order{UserID: user.ID, Amount: 200.00}
    if err := tx.Create(&order).Error; err != nil {
        return nil, err
    }
    
    return &TransactionResult{
        UserID:  user.ID,
        OrderID: order.ID,
    }, nil
})
```

### Nested Transactions

```go
err := client.PerformTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
    // Outer transaction
    if err := tx.Create(&User{Name: "Outer"}).Error; err != nil {
        return err
    }
    
    // Nested transaction
    return tx.Transaction(func(tx2 *gorm.DB) error {
        return tx2.Create(&User{Name: "Inner"}).Error
    })
})
```

## GORM Integration

### Model Definition

```go
type User struct {
    gorm.Model
    Name     string `gorm:"size:255;not null"`
    Email    string `gorm:"size:255;uniqueIndex"`
    Orders   []Order
}

type Order struct {
    gorm.Model
    UserID   uint
    Amount   float64
    Status   string
}

// Auto-migrate models
if err := client.DB().AutoMigrate(&User{}, &Order{}); err != nil {
    log.Fatal(err)
}
```

### Advanced Queries

```go
// Complex query with joins and conditions
var users []User
err := client.DB().
    Preload("Orders", "status = ?", "completed").
    Joins("LEFT JOIN orders ON orders.user_id = users.id").
    Where("orders.amount > ?", 1000).
    Group("users.id").
    Having("COUNT(orders.id) > ?", 5).
    Find(&users).Error
```

## Error Handling

```go
// Custom error types
var (
    ErrConnectionFailed = errors.New("failed to connect to postgresql")
    ErrQueryTimeout     = errors.New("query timeout")
    ErrDuplicateKey    = errors.New("duplicate key violation")
)

// Error handling example
if err := client.DB().Create(&user).Error; err != nil {
    switch {
    case errors.Is(err, gorm.ErrRecordNotFound):
        // Handle not found
    case errors.Is(err, gorm.ErrDuplicatedKey):
        // Handle duplicate key
    case errors.Is(err, context.DeadlineExceeded):
        // Handle timeout
    default:
        // Handle other errors
    }
    return err
}
```

## Testing

### In-Memory Database

```go
func TestUserService(t *testing.T) {
    // Create test client
    client, err := postgres.NewInMemoryTestDbClient(
        &User{},
        &Order{},
    )
    if err != nil {
        t.Fatal(err)
    }
    defer client.Close()

    // Run tests
    t.Run("CreateUser", func(t *testing.T) {
        user := &User{Name: "Test User"}
        err := client.DB().Create(user).Error
        assert.NoError(t, err)
        assert.NotZero(t, user.ID)
    })
}
```

### Transaction Testing

```go
func TestOrderCreation(t *testing.T) {
    client, err := postgres.NewInMemoryTestDbClient(&Order{})
    if err != nil {
        t.Fatal(err)
    }
    
    handler := client.ConfigureNewTxCleanupHandlerForUnitTests()
    defer handler.savePointRollbackHandler(handler.Tx)
    
    t.Run("CreateOrder", func(t *testing.T) {
        // Your test code here
        order := &Order{Amount: 100}
        err := handler.Tx.Create(order).Error
        assert.NoError(t, err)
    })
}
```

## Monitoring & Instrumentation

```go
// Initialize with instrumentation
client, err := postgres.New(
    postgres.WithInstrumentationClient(&instrumentation.Client{
        ServiceName: "my-service",
        Environment: "production",
    }),
    postgres.WithMetricsEnabled(true),
)

// Access metrics
metrics := client.GetMetrics()
fmt.Printf("Total Queries: %d\n", metrics.TotalQueries)
fmt.Printf("Failed Queries: %d\n", metrics.FailedQueries)
fmt.Printf("Average Query Time: %v\n", metrics.AverageQueryTime)
```

## Best Practices

### Connection Management
```go
// Initialize once at application startup
client, err := postgres.New(
    postgres.WithMaxIdleConnections(10),
    postgres.WithMaxOpenConnections(100),
    postgres.WithMaxConnectionLifetime(time.Hour),
)
if err != nil {
    log.Fatal(err)
}
defer client.Close()
```

### Query Optimization
```go
// Use indexes effectively
db.Exec(`CREATE INDEX idx_users_email ON users(email)`)

// Use appropriate batch sizes
const batchSize = 1000
for i := 0; i < len(records); i += batchSize {
    batch := records[i:min(i+batchSize, len(records))]
    client.DB().CreateInBatches(batch, batchSize)
}
```

## API Reference

See our [GoDoc](https://godoc.org/github.com/SolomonAIEngineering/backend-core-library) for complete API documentation.

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/SolomonAIEngineering/backend-core-library.git

# Install dependencies
go mod download

# Run tests
make test

# Run linting
make lint

# Generate coverage report
make coverage
```

## License

This PostgreSQL client is released under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.

---

## Support

- 📚 [Documentation](https://github.com/SolomonAIEngineering/backend-core-library)
- 💬 [GitHub Discussions](https://github.com/SolomonAIEngineering/backend-core-library/discussions)
- 🐛 [Issue Tracker](https://github.com/SolomonAIEngineering/backend-core-library/issues)