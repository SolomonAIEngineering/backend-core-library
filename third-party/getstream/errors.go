package getstream // import "github.com/SimifiniiCTO/simfiny-core-lib/third-party/getstream"

import "errors"

var (
	ErrInvalidKeyOrSecret           = errors.New("invalid key or secret")
	ErrInvalidInstrumentationClient = errors.New("invalid instrumentation client")
	ErrInvalidLogger                = errors.New("invalid logger")
)
