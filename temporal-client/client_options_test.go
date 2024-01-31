package temporalclient

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestWithCertPath(t *testing.T) {
	type testCase struct {
		Name                   string
		CertPath               string
		ExpectedTemporalOption TemporalOption
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualTemporalOption := WithCertPath(tc.CertPath)
			config := &TemporalConfig{}

			// Apply the TemporalOption to our config
			actualTemporalOption(config)

			// Check if the CertPath was set correctly
			if config.ClientCertPath != tc.CertPath {
				t.Errorf("WithCertPath() did not set the correct value, got: %s, want: %s", config.ClientCertPath, tc.CertPath)
			}
		})
	}

	// Define your test cases
	testCases := []testCase{
		{
			Name:                   "Test WithCertPath with non-empty cert path",
			CertPath:               "example-cert.pem",
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.ClientCertPath = "example-cert.pem" },
		},
		{
			Name:                   "Test WithCertPath with empty cert path",
			CertPath:               "",
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.ClientCertPath = "" },
		},
		{
			Name:                   "Test WithCertPath with different cert paths",
			CertPath:               "cert1.pem",
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.ClientCertPath = "cert1.pem" },
		},
	}

	// Iterate through the test cases and run the validation function
	for _, tc := range testCases {
		validate(t, &tc)
	}
}

func TestWithKeyPath(t *testing.T) {
	type testCase struct {
		Name string

		KeyPath string

		ExpectedTemporalOption TemporalOption
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualTemporalOption := WithKeyPath(tc.KeyPath)
			config := &TemporalConfig{}
			// Apply the TemporalOption to our config
			actualTemporalOption(config)
			assert.Equal(t, tc.KeyPath, config.ClientKeyPath)
		})
	}

	// Define your test cases
	testCases := []testCase{
		{
			Name:                   "Test WithKeyPath with non-empty key path",
			KeyPath:                "example-key.pem",
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.ClientKeyPath = "example-key.pem" },
		},
		{
			Name:                   "Test WithKeyPath with empty key path",
			KeyPath:                "",
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.ClientKeyPath = "" },
		},
		{
			Name:                   "Test WithKeyPath with different key paths",
			KeyPath:                "key1.pem",
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.ClientKeyPath = "key1.pem" },
		},
	}

	// Iterate through the test cases and run the validation function
	for _, tc := range testCases {
		validate(t, &tc)
	}
}

func TestWithHostPort(t *testing.T) {
	type testCase struct {
		Name string

		HostPort string

		ExpectedTemporalOption TemporalOption
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualTemporalOption := WithHostPort(tc.HostPort)

			assert.Equal(t, tc.ExpectedTemporalOption, actualTemporalOption)
		})
	}

	// Define your test cases
	testCases := []testCase{
		{
			Name:                   "Test WithHostPort with valid host:port",
			HostPort:               "localhost:8080",
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.HostPort = "localhost:8080" },
		},
		{
			Name:                   "Test WithHostPort with empty host:port",
			HostPort:               "",
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.HostPort = "" },
		},
		{
			Name:                   "Test WithHostPort with different host:port",
			HostPort:               "example.com:9000",
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.HostPort = "example.com:9000" },
		},
	}

	// Iterate through the test cases and run the validation function
	for _, tc := range testCases {
		validate(t, &tc)
	}
}

func TestWithMtlsEnabled(t *testing.T) {
	type testCase struct {
		Name string

		MtlsEnabled bool

		ExpectedTemporalOption TemporalOption
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualTemporalOption := WithMtlsEnabled(tc.MtlsEnabled)

			assert.Equal(t, tc.ExpectedTemporalOption, actualTemporalOption)
		})
	}

	// Define your test cases
	testCases := []testCase{
		{
			Name:                   "Test WithMtlsEnabled with enabled MTLS",
			MtlsEnabled:            true,
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.MtlsEnabled = true },
		},
		{
			Name:                   "Test WithMtlsEnabled with disabled MTLS",
			MtlsEnabled:            false,
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.MtlsEnabled = false },
		},
	}

	// Iterate through the test cases and run the validation function
	for _, tc := range testCases {
		validate(t, &tc)
	}
}

func TestWithLogger(t *testing.T) {
	type testCase struct {
		Name string

		Logger *zap.Logger

		ExpectedTemporalOption TemporalOption
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualTemporalOption := WithLogger(tc.Logger)

			assert.Equal(t, tc.ExpectedTemporalOption, actualTemporalOption)
		})
	}

	// Define your test cases
	testCases := []testCase{
		{
			Name:                   "Test WithLogger with non-nil logger",
			Logger:                 zap.NewExample(),
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.Logger = zap.NewExample() },
		},
		{
			Name:                   "Test WithLogger with nil logger",
			Logger:                 nil,
			ExpectedTemporalOption: func(tc *TemporalConfig) { tc.Logger = nil },
		},
	}

	// Iterate through the test cases and run the validation function
	for _, tc := range testCases {
		validate(t, &tc)
	}
}

func TestTemporalConfigValidate(t *testing.T) {
	type testCase struct {
		Name string

		TemporalConfig *TemporalConfig

		ExpectedError error
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualError := tc.TemporalConfig.validate()

			assert.Equal(t, tc.ExpectedError, actualError)
		})
	}

	// Define your test cases
	testCases := []testCase{
		{
			Name:           "Test TemporalConfigValidate with valid configuration",
			TemporalConfig: &TemporalConfig{HostPort: "localhost:8080", MtlsEnabled: true},
			ExpectedError:  nil,
		},
		{
			Name:           "Test TemporalConfigValidate with invalid configuration",
			TemporalConfig: &TemporalConfig{HostPort: "", MtlsEnabled: false},
			ExpectedError:  errors.New("ServerHostPort cannot be empty"),
		},
	}

	// Iterate through the test cases and run the validation function
	for _, tc := range testCases {
		validate(t, &tc)
	}
}

func TestNewTemporalConfig(t *testing.T) {
	type testCase struct {
		Name string

		Opts []TemporalOption

		ExpectedTemporalConfig *TemporalConfig
		ExpectedError          error
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualTemporalConfig, actualError := NewTemporalConfig(tc.Opts...)

			assert.Equal(t, tc.ExpectedTemporalConfig, actualTemporalConfig)
			assert.Equal(t, tc.ExpectedError, actualError)
		})
	}

	// Define your test cases
	testCases := []testCase{
		{
			Name:                   "Test NewTemporalConfig with valid options",
			Opts:                   []TemporalOption{WithHostPort("localhost:8080"), WithMtlsEnabled(true)},
			ExpectedTemporalConfig: &TemporalConfig{HostPort: "localhost:8080", MtlsEnabled: true},
			ExpectedError:          nil,
		},
		{
			Name:                   "Test NewTemporalConfig with invalid options",
			Opts:                   []TemporalOption{WithHostPort(""), WithMtlsEnabled(false)},
			ExpectedTemporalConfig: nil,
			ExpectedError:          errors.New("ServerHostPort cannot be empty"),
		},
	}

	// Iterate through the test cases and run the validation function
	for _, tc := range testCases {
		validate(t, &tc)
	}
}
