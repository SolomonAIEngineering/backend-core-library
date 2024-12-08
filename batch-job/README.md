<div align="center">
    <h1 align="center">Solomon AI Batch Job Library</h1>
    <h3 align="center">Efficient and Type-safe Batch Job Processing for Go Applications</h3>
</div>

<div align="center">
  
[![Go Report Card](https://goreportcard.com/badge/github.com/SolomonAIEngineering/backend-core-library)](https://goreportcard.com/report/github.com/SolomonAIEngineering/backend-core-library)
[![GoDoc](https://godoc.org/github.com/SolomonAIEngineering/backend-core-library?status.svg)](https://godoc.org/github.com/SolomonAIEngineering/backend-core-library)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[Documentation](https://github.com/SolomonAIEngineering/backend-core-library) | 
[API Reference](https://github.com/SolomonAIEngineering/backend-core-library) | 
[Examples](https://github.com/SolomonAIEngineering/backend-core-library/examples)

</div>

## Overview

The Solomon AI Batch Job Library is a production-ready Go package that simplifies batch job processing with a focus on type safety, reliability, and developer experience. Built on top of [Asynq](https://github.com/hibiken/asynq), it provides a robust foundation for scheduling and executing recurring tasks in your Go applications.

## Key Features

🚀 **High Performance**
- Efficient job scheduling with minimal overhead
- Optimized for high-throughput scenarios
- Built-in connection pooling and resource management

⚡ **Developer Experience**
- Intuitive builder pattern API
- Strongly typed configurations
- Comprehensive validation out of the box
- Clear error messages and debugging support

🛡️ **Reliability**
- Sophisticated retry mechanism with exponential backoff
- Job persistence and recovery
- Graceful shutdown handling
- Detailed job execution metrics

🔧 **Flexibility**
- Multiple scheduling interval options
- Customizable job parameters
- Easy integration with existing applications
- Extensible architecture

## Installation

```bash
go get github.com/SolomonAIEngineering/backend-core-library/batch
```

## Quick Start

Here's a simple example to get you started:

```go
package main

import (
    "context"
    "log"
    
    batchjob "github.com/SolomonAIEngineering/backend-core-library/batch-job"
    taskprocessor "github.com/SolomonAIEngineering/backend-core-library/task-processor"
    "github.com/hibiken/asynq"
)

func main() {
    // 1. Create a batch job
    task := asynq.NewTask("cleanup", nil)
    job := batchjob.NewBatchJob(
        batchjob.WithBaseConfig(batchjob.BaseConfig{
            Enabled:    true,
            Interval:   "@daily",
            MaxRetries: 3,
        }),
        batchjob.WithTask(task),
        batchjob.WithTaskId(stringPtr("daily-cleanup")),
    )

    // 2. Create batch jobs collection
    jobs := batchjob.NewBatchJobs(
        batchjob.WithBatchJob(job),
    )

    // 3. Initialize the task processor
    processorConfig := taskprocessor.Config{
        RedisAddr: "localhost:6379",
        // Add other configuration options as needed
    }
    processor := taskprocessor.NewTaskProcessor(processorConfig)

    // 4. Initialize and start the batcher
    batcher := batchjob.NewBatcher(
        batchjob.WithBatchJobsRef(jobs),
        batchjob.WithTaskProcessor(processor),
    )

    // 5. Register and start the jobs
    if err := batcher.RegisterRecurringBatchJobs(context.Background()); err != nil {
        log.Fatal(err)
    }
}

func stringPtr(s string) *string {
    return &s
}
```

## Scheduling Options

The library supports various interval patterns:

### Standard Intervals
- `@yearly` or `@annually` - Run once a year at midnight of January 1
- `@monthly` - Run once a month at midnight of the first day
- `@weekly` - Run once a week at midnight of Sunday
- `@daily` or `@midnight` - Run once a day at midnight
- `@hourly` - Run once an hour at the beginning of the hour

### Custom Intervals
- Minutes: `@every 1m`, `@every 5m`, `@every 15m`
- Hours: `@every 1h`, `@every 2h`, `@every 6h`
- Custom: `@every 30s`, `@every 2h30m`

### Cron Expressions
```go
job := batchjob.NewBatchJob(
    batchjob.WithBaseConfig(batchjob.BaseConfig{
        Interval: "0 */2 * * *", // Every 2 hours
    }),
    // ... other configurations
)
```

## Advanced Usage

### Configuring Retries

```go
job := batchjob.NewBatchJob(
    batchjob.WithBaseConfig(batchjob.BaseConfig{
        MaxRetries: 5,
        RetryDelays: []time.Duration{
            time.Second * 30,
            time.Minute * 1,
            time.Minute * 5,
            time.Minute * 15,
            time.Hour * 1,
        },
    }),
    // ... other configurations
)
```

### Custom Task Parameters

```go
payload := map[string]interface{}{
    "batchSize": 1000,
    "priority":  "high",
}
task := asynq.NewTask("data-processing", payload)

job := batchjob.NewBatchJob(
    batchjob.WithTask(task),
    // ... other configurations
)
```

### Job Monitoring and Metrics

The library provides built-in monitoring capabilities:

```go
type JobMetrics struct {
    SuccessCount   int64
    FailureCount   int64
    LastExecuted   time.Time
    ExecutionTime  time.Duration
    RetryCount     int
}

// Access metrics for a specific job
metrics := batcher.GetJobMetrics("daily-cleanup")
```

## Best Practices

1. **Error Handling**: Always implement proper error handling and logging
   ```go
   if err := batcher.RegisterRecurringBatchJobs(ctx); err != nil {
       log.Printf("Failed to register batch jobs: %v", err)
       // Implement your error handling strategy
   }
   ```

2. **Resource Management**: Properly close resources when shutting down
   ```go
   defer processor.Shutdown()
   defer batcher.Shutdown(ctx)
   ```

3. **Configuration**: Use environment variables for configuration
   ```go
   config := batchjob.BaseConfig{
       Enabled:    os.Getenv("BATCH_JOB_ENABLED") == "true",
       Interval:   os.Getenv("BATCH_JOB_INTERVAL"),
       MaxRetries: 3,
   }
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