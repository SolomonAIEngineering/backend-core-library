package algoliasearch

import (
	"testing"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	"github.com/stretchr/testify/assert"
)

func TestAlgoliaSearchConfig_Validate(t *testing.T) {
	// Test validation of a valid config
	config := &Config{
		ApplicationID: "test-app-id",
		APIKey:        "test-api-key",
		IndexName:     "test-index-name",
		TelemetrySDK:  &instrumentation.Client{},
	}
	err := config.Validate()
	assert.Nil(t, err)

	// Test validation of a config with a missing application ID
	config = &Config{
		APIKey:    "test-api-key",
		IndexName: "test-index-name",
	}
	err = config.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrMissingApplicationID)

	// Test validation of a config with a missing API key
	config = &Config{
		ApplicationID: "test-app-id",
		IndexName:     "test-index-name",
	}
	err = config.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrMissingAPIKey)

	// Test validation of a config with a missing index name
	config = &Config{
		ApplicationID: "test-app-id",
		APIKey:        "test-api-key",
	}
	err = config.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrMissingIndexName)
}

func TestNewClient(t *testing.T) {
	// Test initialization of the client
	client, err := New(
		WithAlgoliaSearchApplicationID("test-app-id"),
		WithAlgoliaSearchAPIKey("test-api-key"),
		WithAlgoliaSearchIndexName("test-index-name"),
		WithAlgoliaSearchTelemetrySDK(&instrumentation.Client{}),
	)
	assert.NotNil(t, client)
	assert.Nil(t, err)

	// Test failure to initialize client due to missing application ID
	_, err = New(
		WithAlgoliaSearchAPIKey("test-api-key"),
		WithAlgoliaSearchIndexName("test-index-name"),
	)
	assert.NotNil(t, err)

	// Test failure to initialize client due to missing API key
	_, err = New(
		WithAlgoliaSearchApplicationID("test-app-id"),
		WithAlgoliaSearchIndexName("test-index-name"),
	)
	assert.NotNil(t, err)

	// Test failure to initialize client due to missing index name
	_, err = New(
		WithAlgoliaSearchApplicationID("test-app-id"),
		WithAlgoliaSearchAPIKey("test-api-key"),
	)
	assert.NotNil(t, err)
}

func TestWithAlgoliaSearchApplicationID(t *testing.T) {
	config := &Config{}

	WithAlgoliaSearchApplicationID("app_id")(config)
	assert.Equal(t, "app_id", config.ApplicationID)
}

func TestWithAlgoliaSearchAPIKey(t *testing.T) {
	config := &Config{}

	WithAlgoliaSearchAPIKey("api_key")(config)
	assert.Equal(t, "api_key", config.APIKey)
}

func TestWithAlgoliaSearchIndexName(t *testing.T) {
	config := &Config{}

	WithAlgoliaSearchIndexName("index_name")(config)
	assert.Equal(t, "index_name", config.IndexName)
}
