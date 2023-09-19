package taskprocessor // import "github.com/SolomonAIEngineering/backend-core-library/task-processor"

// ProcessingInterval represents the various processing intervals for a recurring task
type ProcessingInterval string

const (
	EveryYear          ProcessingInterval = "@yearly"
	EveryMonth         ProcessingInterval = "@monthly"
	EveryWeek          ProcessingInterval = "@weekly"
	EveryDayAtMidnight ProcessingInterval = "@midnight"
	EveryDay           ProcessingInterval = "@daily"
	Every24Hours       ProcessingInterval = "@every 24h"
	Every12Hours       ProcessingInterval = "@every 12h"
	Every6Hours        ProcessingInterval = "@every 6h"
	Every3Hours        ProcessingInterval = "@every 3h"
	EveryHour          ProcessingInterval = "@every 1h"
	Every30Minutes     ProcessingInterval = "@every 30m"
	Every15Minutes     ProcessingInterval = "@every 15m"
	Every10Minutes     ProcessingInterval = "@every 10m"
	Every5Minutes      ProcessingInterval = "@every 5m"
	Every3Minutes      ProcessingInterval = "@every 3m"
	Every1Minutes      ProcessingInterval = "@every 1m"
	Every30Seconds     ProcessingInterval = "@every 30s"
)

// The `String()` method is a method defined on the `ProcessingInterval` type. It is used to convert a
// `ProcessingInterval` value to its corresponding string representation.
func (p ProcessingInterval) String() string {
	return string(p)
}
