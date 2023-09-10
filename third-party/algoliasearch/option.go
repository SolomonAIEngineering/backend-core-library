package algoliasearch

import "github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"

type (
	// Config is the configuration for the algolia search handler
	Config struct {
		// ApplicationID is the algolia application id
		ApplicationID string
		// APIKey is the algolia api key
		APIKey string
		// IndexName is the name of the index to use
		IndexName string
		// TelemetrySDK is the telemetry sdk
		TelemetrySDK *instrumentation.Client
	}
)

// Option is a function that configures the algolia search handler
type Option func(*Config)

func (c *Config) Validate() error {
	if c.ApplicationID == "" {
		return ErrMissingApplicationID
	}

	if c.APIKey == "" {
		return ErrMissingAPIKey
	}

	if c.IndexName == "" {
		return ErrMissingIndexName
	}

	if c.TelemetrySDK == nil {
		return ErrMissingTelemetrySDK
	}

	return nil
}

// WithAlgoliaSearchApplicationID sets the application id for the algolia search handler
func WithAlgoliaSearchApplicationID(id string) Option {
	return func(c *Config) {
		c.ApplicationID = id
	}
}

// WithAlgoliaSearchAPIKey sets the api key for the algolia search handler
func WithAlgoliaSearchAPIKey(key string) Option {
	return func(c *Config) {
		c.APIKey = key
	}
}

// WithAlgoliaSearchIndexName sets the index name for the algolia search handler
func WithAlgoliaSearchIndexName(name string) Option {
	return func(c *Config) {
		c.IndexName = name
	}
}

// WithAlgoliaSearchTelemetrySDK sets the telemetry sdk for the algolia search handler
func WithAlgoliaSearchTelemetrySDK(sdk *instrumentation.Client) Option {
	return func(c *Config) {
		c.TelemetrySDK = sdk
	}
}
