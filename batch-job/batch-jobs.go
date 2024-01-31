package batch_job // import "github.com/SolomonAIEngineering/backend-core-library/batch-job"

type BatchJobs struct {
	jobs []*BatchJob
}

type BatchJobsOption func(*BatchJobs)

// WithBatchJobs sets the batch jobs for the batcher.
func WithBatchJob(batchJob *BatchJob) BatchJobsOption {
	return func(b *BatchJobs) {
		if len(b.jobs) == 0 {
			b.jobs = make([]*BatchJob, 0)
		}

		b.jobs = append(b.jobs, batchJob)
	}
}

func WithBatchJobs(batchJobs *BatchJobs) BatchJobsOption {
	return func(b *BatchJobs) {
		b.jobs = batchJobs.jobs
	}
}

func NewBatchJobs(options ...BatchJobsOption) *BatchJobs {
	batchJobs := &BatchJobs{}

	for _, option := range options {
		option(batchJobs)
	}

	return batchJobs
}

// AddJob appends a new BatchJob to the jobs slice.
func (b *BatchJobs) AddJob(job *BatchJob) {
	b.jobs = append(b.jobs, job)
}
