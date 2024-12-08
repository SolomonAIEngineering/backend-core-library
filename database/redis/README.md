# Redis Client 

This is a Redis client wrapper for Golang that provides a simplified interface to interact with Redis.

## Installation
To use this library install using the `go get commant`
```go
go get -u github.com/SimifiniiCTO/backend-core-lib/database/redis
```

## Usage 
To use this library first import 
```go
import github.com/SimifiniiCTO/backend-core-lib/database/redis
import github.com/SimifiniiCTO/backend-core-lib/signals
```

## Create The Client 
To create a new Redis client, use the New function:
```go
    stopCh := signals.SetupSignalHandler()
    opts := []redis.Options{
        redis.WithLogger(zap.L()),
        redis.WithTelemetrySdk(&instrumentation.Client{}),
        redis.WithURI(host),
        redis.WithServiceName(serviceName), 
        redis.WithCacheTTLInSeconds(cacheTTLInSeconds),
    }

    c, err := redis.New(stopCh, opts...)
    if err != nil {
        ...
    }
```

Where:
* __host__: the Redis server host (string).
* __serviceName__: the service initialization the redis connection 
* __logger__: zap logger instance
* __cacheTTLInSeconds__: the TTL (in seconds) for cache entries (int).
* __telemetrySdk__: the telemetry SDK to use (an object implementing the ITelemetrySdk interface).

## Reading data from Redis
To read data from Redis, use the Read function:
```go
value, err := client.Read(ctx, key)
if err != nil {
    // handle error
}
```

Where:
* __ctx__: the context (context.Context) for the Redis read operation.
* __key__: the key (string) for the data to read.

## Deleting data from Redis
To delete data from Redis, use the Delete function:
```go
err := client.Delete(ctx, key)
if err != nil {
    // handle error
}
```

Where:
* __ctx__: the context (context.Context) for the Redis delete operation.
* __key__: the key (string) for the data to read.

## Testing
To run the unit tests, use the go test command:
```go
go test -v
```

## Contributing
Contributions to this Redis client  are welcome. To contribute, follow these steps:

* Fork this repository.
* Create a new branch: __`git checkout -b my-new-feature`__.
* Make changes and add tests for the new feature.
* Run tests and ensure they pass: __`go test -v`__.
* Commit your changes: __`git commit -am 'Add some feature'`__.
* Push to the branch: __`git push origin my-new-feature`__.
*  Create a new Pull Request.

## License
This Redis client is released under the MIT License. See LICENSE for more information.