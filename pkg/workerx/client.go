package workerx

import (
	"be-border-service/internal/config"

	"github.com/hibiken/asynq"
)

type AsynqClient interface {
	Enqueue(task *asynq.Task, opts ...EnqueueOptions) (*asynq.TaskInfo, error)
}

func NewAsynqClient(cfg *config.Config) AsynqClient {
	return &asynqServerConfig{
		RedisAddr: "",
		password:  "user",
	}
}

func (a *asynqServerConfig) Enqueue(task *asynq.Task, opts ...EnqueueOptions) (*asynq.TaskInfo, error) {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:         a.RedisAddr,
		Username:     a.username,
		Password:     a.password,
		DB:           a.DB,
		ReadTimeout:  a.ReadTimeout,
		WriteTimeout: a.WriteTimeout,
	})

	defer client.Close()

	var asynqOpts []asynq.Option
	for _, opt := range opts {
		opt(&asynqOpts)
	}

	return client.Enqueue(task, asynqOpts...)
}
