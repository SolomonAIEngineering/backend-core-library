package signals // import "github.com/SimifiniiCTO/simfiny-core-lib/signals"

import (
	"os"
)

var shutdownSignals = []os.Signal{os.Interrupt}
