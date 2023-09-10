package getstream

import "errors"

var (
	ErrInvalidKeyOrSecret           = errors.New("invalid key or secret")
	ErrInvalidInstrumentationClient = errors.New("invalid instrumentation client")
	ErrInvalidLogger                = errors.New("invalid logger")
)
