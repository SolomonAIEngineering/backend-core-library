package algoliasearch

import (
	"context"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type (
	// IClient is the interface for the algolia search handler
	IClient interface {
		// Sends a piece of data to algolia search
		Send(ctx context.Context, data Record) (*string, error)
		// Deletes a piece of data from algolia search
		Delete(ctx context.Context, objectId string) error
		// Update updates a piece of data in algolia search
		Update(ctx context.Context, record Record) error
		// configureSearchSettings configures the search settings for the index
		configureSearchSettings() error
	}

	// Client is the handler for algolia search
	Client struct {
		// client is the algolia search client
		client *search.Client
		// index is the algolia search index
		index *search.Index
		// telemetrySdk is the telemetry sdk
		telemetrySdk *instrumentation.Client
	}
)

var _ IClient = (*Client)(nil)

// NewClient creates a new algolia search handler
func New(opts ...Option) (*Client, error) {
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	client := search.NewClient(config.ApplicationID, config.APIKey)
	index := client.InitIndex(config.IndexName)

	h := &Client{
		client:       client,
		index:        index,
		telemetrySdk: config.TelemetrySDK,
	}

	// configure index search settings
	h.configureSearchSettings()

	return h, nil
}

// configureSearchSettings configures the search settings for the index
func (h *Client) configureSearchSettings() error {
	_, err := h.index.SetSettings(search.Settings{
		// Select the attributes you want to search in
		SearchableAttributes: opt.SearchableAttributes(
			"unordered(name)",
			"unordered(tags)",
			"unordered(extra.rules)",
			"unordered(extra.description)",
		),
		// Define business metrics for ranking and sorting
		CustomRanking: opt.CustomRanking(
			"desc(name)",
		),
		// Set up some attributes to filter results on
		AttributesForFaceting: opt.AttributesForFaceting(
			"name", "tags",
		),
	})
	return err
}
