package signals // import "github.com/SolomonAIEngineering/backend-core-library/signals"

import (
	"os"
)

var shutdownSignals = []os.Signal{os.Interrupt}
