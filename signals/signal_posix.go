//go:build !windows
// +build !windows

package signals // import "github.com/SolomonAIEngineering/backend-core-library/signals"

import (
	"os"
	"syscall"
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
