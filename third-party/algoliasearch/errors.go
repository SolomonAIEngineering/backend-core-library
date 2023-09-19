package algoliasearch // import "github.com/SimifiniiCTO/simfiny-core-lib/third-party/algoliasearch"

import "errors"

var (
	// ErrMissingApplicationID is returned when the application id is missing
	ErrMissingApplicationID = errors.New("missing application id")
	// ErrMissingAPIKey is returned when the api key is missing
	ErrMissingAPIKey = errors.New("missing api key")
	// ErrMissingIndexName is returned when the index name is missing
	ErrMissingIndexName = errors.New("missing index name")
	// ErrMissingObjectID is returned when the object id is missing
	ErrMissingObjectID = errors.New("missing object id")
	// ErrMissingTelemetrySDK is returned when the telemetry sdk is missing
	ErrMissingTelemetrySDK = errors.New("missing telemetry sdk")
)
