# Feeds Library

This is a Go wrapper for the GetStream API. The GetStream API is a service for building scalable newsfeeds and activity streams. 
The wrapper is designed to make it easy for Go developers to interact with the GetStream API by providing a simple and easy-to-use interface.

## Installation
To use this wrapper, you need to have Go installed on your machine. You can install Go by following the instructions on the official website.

Once you have Go installed, you can install the GetStream wrapper by running the following command:
```go
go get github.com/SimifiniiCTO/simfiny-core-lib/third-party/getstream
```

## Usage 
To use the GetStream wrapper, you need to first create a new client by providing your API credentials. You can then use this client to interact with the GetStream API by calling various methods.

Here's an example of how to create a new client:
```go
import (
    "github.com/SimifiniiCTO/simfiny-core-lib/third-party/getstream"
)

func main() {
    opts := []getstream.Options{
        getstream.WithLogger(zap.L()),
        getstream.WithKey("key"),
        getstream.WithSecret("secret"),
        getstream.WithInstrumentationClient(&instrumentation.Client{})
    }
    client, err := getstream.New(opts...)
    if err != nil {
        // handle error
    }
}
```

You can then use the client to interact with the GetStream API. Here are some examples of what you can do:

## Creating a new activity
```go
feedId := "feed-id"
ctx := context.Background()
actor := stream.NewUser("user:1")
object := stream.NewObject("product:1")
verb := "purchased"

activity := stream.NewActivity(actor, verb, object)
if _, err := client.CreateActivity(ctx, feedId, activity); err != nil {
    // handle error
}
```

## Fetching Activities
```go
feedId := "feed-id"
ctx := context.Background()
if activities, err := client.GetTimeline(ctx, feedId); err != nil {
    // handle error
} else {
    // do something with activities
}
```

## Follow Activities
```go
source := "feed-id-source"
target := "feed-id-target"
ctx := context.Background()
if err := client.FollowFeed(ctx, source, target); err != nil {
    ...
}
```

## Testing
To test the GetStream wrapper, you can use the built-in testing framework in Go. The tests are located in the *_test.go files in the repository.

To run the tests, simply run the following command:
```go
go test -v ./...
```

This will run all the tests in the repository and print out the results.

## Contributions

Contributions to the GetStream wrapper are welcome! If you find a bug or have an idea for a new feature, feel free to open an issue or submit a pull request.

