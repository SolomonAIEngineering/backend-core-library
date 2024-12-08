package batch_job // import "github.com/SolomonAIEngineering/backend-core-library/batch-job"

// BatchJobs manages a collection of batch jobs that can be scheduled and executed.
// It provides methods for adding and managing individual BatchJob instances.
type BatchJobs struct {
	jobs []*BatchJob
}

// BatchJobsOption defines a function type for configuring BatchJobs instances.
// This follows the functional options pattern for flexible configuration.
type BatchJobsOption func(*BatchJobs)

// WithBatchJob adds a single batch job to the collection.
// If the jobs slice hasn't been initialized, it will create a new slice.
//
// Example:
//
//	job := &BatchJob{...}
//	jobs := NewBatchJobs(WithBatchJob(job))
func WithBatchJob(batchJob *BatchJob) BatchJobsOption {
	return func(b *BatchJobs) {
		if len(b.jobs) == 0 {
			b.jobs = make([]*BatchJob, 0)
		}

		b.jobs = append(b.jobs, batchJob)
	}
}

// WithBatchJobs replaces the entire collection of batch jobs with the provided jobs.
// This is useful when you want to configure multiple jobs at once.
//
// Example:
//
//	existingJobs := &BatchJobs{...}
//	newJobs := NewBatchJobs(WithBatchJobs(existingJobs))
func WithBatchJobs(batchJobs *BatchJobs) BatchJobsOption {
	return func(b *BatchJobs) {
		b.jobs = batchJobs.jobs
	}
}

// NewBatchJobs creates a new BatchJobs instance with the provided options.
// It uses the functional options pattern to allow flexible configuration.
//
// Example:
//
//	jobs := NewBatchJobs(
//	    WithBatchJob(job1),
//	    WithBatchJob(job2),
//	)
func NewBatchJobs(options ...BatchJobsOption) *BatchJobs {
	batchJobs := &BatchJobs{}

	for _, option := range options {
		option(batchJobs)
	}

	return batchJobs
}

// AddJob appends a new BatchJob to the jobs collection.
// This method can be used to add jobs after initialization.
//
// Example:
//
//	jobs := NewBatchJobs()
//	jobs.AddJob(newJob)
func (b *BatchJobs) AddJob(job *BatchJob) {
	b.jobs = append(b.jobs, job)
}
