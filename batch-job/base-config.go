package batch_job // import "github.com/SolomonAIEngineering/backend-core-library/batch-job"

import (
	"fmt"

	taskprocessor "github.com/SolomonAIEngineering/backend-core-library/task-processor"
)

// BaseConfig provides fundamental configuration options for batch jobs.
// It defines common settings such as job scheduling intervals and retry policies.
type BaseConfig struct {
	// Enabled determines if the batch job is active and should be processed
	Enabled bool `json:"enabled"`

	// Interval specifies the frequency at which the job should run.
	// Supports cron-like syntax and duration strings (e.g., "@daily", "10s", "1m", "1h")
	Interval string `json:"interval"`

	// MaxRetries defines the maximum number of retry attempts for failed jobs
	MaxRetries int64 `json:"maxRetries"`
}

// ProcessingInterval converts the string interval configuration into a structured
// taskprocessor.ProcessingInterval value.
//
// Example:
//
//	cfg := BaseConfig{Interval: "@daily"}
//	interval := cfg.ProcessingInterval() // Returns taskprocessor.EveryDayAtMidnight
func (c *BaseConfig) ProcessingInterval() taskprocessor.ProcessingInterval {
	return toProcessingInterval(&c.Interval)
}

// Validate checks the correctness of the base configuration values.
// Returns an error if any values are invalid.
//
// Validation rules:
// - MaxRetries must be non-negative
// - Interval must not be empty
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

// toProcessingInterval converts a string interval specification to a ProcessingInterval value.
// Supports various time formats including cron-style specifications and duration strings.
//
// Supported formats include:
// - Cron-style:
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
