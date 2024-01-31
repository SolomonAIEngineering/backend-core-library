package batch_job // import "github.com/SolomonAIEngineering/backend-core-library/batch-job"

import (
	"fmt"

	taskprocessor "github.com/SolomonAIEngineering/backend-core-library/task-processor"
)

type BaseConfig struct {
	Enabled    bool   `json:"enabled"`
	Interval   string `json:"interval"` // e.g. "10s", "1m", "1h"
	MaxRetries int64  `json:"maxRetries"`
}

func (c *BaseConfig) ProcessingInterval() taskprocessor.ProcessingInterval {
	return toProcessingInterval(&c.Interval)
}

// Validate checks the correctness of the base config values.
func (c *BaseConfig) Validate() error {
	if c.MaxRetries < 0 {
		return fmt.Errorf("maxRetries should be a non-negative integer")
	}

	// Check interval format validity (this doesn't validate every possible string but is a start)
	if c.Interval == "" {
		return fmt.Errorf("interval is empty")
	}

	return nil
}

// toProcessingInterval converts a string to a processing interval
func toProcessingInterval(str *string) taskprocessor.ProcessingInterval {
	if str == nil {
		// handle nil input in some way, I'll return EveryDay as a default for this example
		return taskprocessor.EveryDay
	}
	switch *str {
	case "@yearly":
		return taskprocessor.EveryYear
	case "@monthly":
		return taskprocessor.EveryMonth
	case "@weekly":
		return taskprocessor.EveryWeek
	case "@midnight":
		return taskprocessor.EveryDayAtMidnight
	case "@daily":
		return taskprocessor.EveryDayAtMidnight
	case "@every 24h":
		return taskprocessor.Every24Hours
	case "@every 12h":
		return taskprocessor.Every12Hours
	case "@every 6h":
		return taskprocessor.Every6Hours
	case "@every 3h":
		return taskprocessor.Every3Hours
	case "@every 1h":
		return taskprocessor.EveryHour
	case "@every 30m":
		return taskprocessor.Every30Minutes
	case "@every 15m":
		return taskprocessor.Every15Minutes
	case "@every 10m":
		return taskprocessor.Every10Minutes
	case "@every 5m":
		return taskprocessor.Every5Minutes
	case "@every 3m":
		return taskprocessor.Every3Minutes
	case "@every 1m":
		return taskprocessor.Every1Minutes
	case "@every 30s":
		return taskprocessor.Every30Seconds
	default:
		// If the string doesn't match any of the predefined intervals, handle it accordingly.
		// I'll return EveryDay as a default for this example.
		return taskprocessor.EveryDay
	}
}
