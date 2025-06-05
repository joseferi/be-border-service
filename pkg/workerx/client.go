package workerx

import (
	"be-border-service/internal/config"
	"be-border-service/internal/constants"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

type AsynqClient interface {
	Enqueue(task *asynq.Task, opts ...EnqueueOptions) (*asynq.TaskInfo, error)
}

func NewAsynqClient(cfg *config.Config) AsynqClient {
	return &asynqServerConfig{
		RedisAddr:    cfg.Redis.Addr,
		username:     cfg.Redis.Username,
		password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		WriteTimeout: cfg.Redis.WriteTimeout,
		ReadTimeout:  cfg.Redis.ReadTimeout,
	}
}
func NewTask(taskType string, payload any) (*asynq.Task, error) {
	if _, ok := constants.AllowedTaskTypes[taskType]; !ok {
		return nil, fmt.Errorf("invalid task type: %s", taskType)
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(taskType, b), nil
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
