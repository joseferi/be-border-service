package workerx

import (
	"time"

	"be-border-service/internal/constants"

	"github.com/hibiken/asynq"
)

type EnqueueOptions func(*[]asynq.Option)

func WithQueueName(name string) EnqueueOptions {
	return func(opts *[]asynq.Option) {
		*opts = append(*opts, asynq.Queue(name))
	}
}

func WithTimeout(timeout time.Duration) EnqueueOptions {
	return func(opts *[]asynq.Option) {
		*opts = append(*opts, asynq.Timeout(timeout))
	}
}

func WithProcessIn(d time.Duration) EnqueueOptions {
	return func(opts *[]asynq.Option) {
		*opts = append(*opts, asynq.ProcessIn(d))
	}
}

func WithProcessAt(t time.Time) EnqueueOptions {
	return func(opts *[]asynq.Option) {
		*opts = append(*opts, asynq.ProcessAt(t))
	}
}

func WithRetention(d time.Duration) EnqueueOptions {
	return func(opts *[]asynq.Option) {
		*opts = append(*opts, asynq.Retention(d))
	}
}

func WithMaxRetry(n int) EnqueueOptions {
	return func(opts *[]asynq.Option) {
		*opts = append(*opts, asynq.MaxRetry(n))
	}
}

func WithQueuePriority(priority constants.QueuePriority) EnqueueOptions {
	// Bisa di-mapping ke nama queue
	return func(opts *[]asynq.Option) {
		switch priority {
		case constants.Critical:
			*opts = append(*opts, asynq.Queue("critical"))
		case constants.Default:
			*opts = append(*opts, asynq.Queue("default"))
		case constants.Low:
			*opts = append(*opts, asynq.Queue("low"))
		default:
			*opts = append(*opts, asynq.Queue("default"))
		}
	}
}
