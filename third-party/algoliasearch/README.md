# Algolia Search Library for Go

This is a wrapper library for the Algolia Search API, designed to make it easier to integrate Algolia into Go applications. It provides a simple and intuitive interface for indexing and searching data in Algolia.

## Installation
You can install the Algolia Search wrapper for Go using go get:

```bash
go get -u github.com/SimifiniiCTO/simfiny-core-lib/third-party/algoliasearch
```

## Usage
To use the Algolia Search wrapper, you first need to create a client instance:

```go
import "github.com/SimifiniiCTO/simfiny-core-lib/third-party/algoliasearch"

opts := []algoliasearch.Options{
    algoliasearch.WithAlgoliaSearchApplicationID("salfhgjdahgjlsdhg"),
    algoliasearch.WithAlgoliaSearchAPIKey("sdlhgjhzdgjsdhgsd"),
    algoliasearch.WithAlgoliaSearchIndexName("test")
}

client, err := algoliasearch.New(opts...)
if err != nil {
    ....
}

```