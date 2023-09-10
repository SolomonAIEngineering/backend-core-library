# Transaction Manager

One Paragraph of the project description

Initially appeared on
[gist](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2). But the page cannot open anymore so that is why I have moved it here.

## Getting Started

These instructions will give you a copy of the project up and running on
your local machine for development and testing purposes. See deployment
for notes on deploying the project on a live system.

### Prerequisites

Requirements for the software and other tools to build, test and push 
- [Example 1](https://www.example.com)
- [Example 2](https://www.example.com)

### Installing

To use the library, you need to have go installed on your machine. You can install Go by following the instructions on the official website.

Once you have Go installed, you can install the library by running the following command
```go
go get github.com/SimifiniiCTO/simfiny-core-lib/third-party/transaction-manager
```

### Usage 
To use the library, you must first initialize the client

```go 
import (
    txm "github.com/SimifiniiCTO/simfiny-core-lib/third-party/transaction-manager"
)

func main(){
    temporalOpts := &client.Options{
		HostPort:  "9999",
		Namespace: "test-namespace",
	}

    retryPolicy := &txm.Policy{
		RetryInitialInterval:    TemporalRetryInitialInterval,
		RetryBackoffCoefficient: TemporalBackoffCoefficient,
		MaximumInterval:         TemporalMaxRetryInterval,
		MaximumAttempts:         TemporalMaxRetryAttempts,
	}

    opts := []txm.Options{
        txm.WithClientOptions(temporalOpts),
        txm.WithInstrumentationClient(&instrumentation.Client{}),
        txm.WithLogger(log),
        txm.WithPostgres(postgresClient), // Optional 
        txm.WithMongo(mongoClient), // Optional
        txm.WithRetryPolicy(retryPolicy),
        txm.WithRpcTimeout(rpcTimeout),
        txm.WithMetricsEnabled(metricsEnabled)
        txm.WithMessageQueueClient(messageQueueClient) // Optional
    }

    txmClient, err := txm.New(opts...)
    if err != nil {
        ...
    }
}
```

## Running the tests

Explain how to run the automated tests for this system

### Sample Tests

Explain what these tests test and why

    Give an example

### Style test

Checks if the best practices and the right coding style has been used.

    Give an example

## Deployment

Add additional notes to deploy this on a live system

## Built With

  - [Contributor Covenant](https://www.contributor-covenant.org/) - Used
    for the Code of Conduct
  - [Creative Commons](https://creativecommons.org/) - Used to choose
    the license

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code
of conduct, and the process for submitting pull requests to us.

## Versioning

We use [Semantic Versioning](http://semver.org/) for versioning. For the versions
available, see the [tags on this
repository](https://github.com/PurpleBooth/a-good-readme-template/tags).

## Authors

  - **Billie Thompson** - *Provided README Template* -
    [PurpleBooth](https://github.com/PurpleBooth)

See also the list of
[contributors](https://github.com/PurpleBooth/a-good-readme-template/contributors)
who participated in this project.

## License

This project is licensed under the [CC0 1.0 Universal](LICENSE.md)
Creative Commons License - see the [LICENSE.md](LICENSE.md) file for
details

## Acknowledgments

  - Hat tip to anyone whose code is used
  - Inspiration
  - etc
