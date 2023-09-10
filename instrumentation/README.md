# Instrumentation Library

This is a simple wrapper for the New Relic API that makes it easier to interact with the New Relic service from your golang code.

## Installation
You can install the New Relic wrapper using `go get`:

```bash
go get -u github.com/SimifiniiCTO/simfiny-core-lib/instrumentation
```

## Usage
You need to to import the library
```go
import github.com/SimifiniiCTO/simfiny-core-lib/instrumentation
```

Then initialize the Instrumentation Client with the following configuration
```go
    opts := []instrumentation.Option{
        instrumentation.WithServiceName("test-service"),
        instrumentation.WithServiceVersion("v1.0.0"),
        instrumentation.WithServiceEnvironment("prod"),
        instrumentation.WithEnabled(true), // enables instrumentation
        instrumentation.WithNewrelicKey("newrelic-license-key"),
        instrumentation.WithLogger(zap.L()),
    }
    client, err := instrumentation.New(opts...)
    if err != nil {
        ... 
    }

```

Now you can user the `client` object to instrument your applicatuon. 
for example, to initiate a trace
```go
    ctx := context.Background()
    tx := client.StartTransaction("test-transaction") // starts a new transaction
    span := client.NewChildSpan("test-span", tx)
    defer span.End()

```

## Development

If you want to contribute to this project, you can clone the repository and install the development dependencies:
```bash
git clone github.com/SimifiniiCTO/simfiny-core-lib/instrumentation
cd simfiny-core-lib/instrumentation
go mod tidy
```

Please see the __contributing guide__ for more information on how to contribute.

## License
This project is licensed under the MIT License - see the LICENSE file for details.