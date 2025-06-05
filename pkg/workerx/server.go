package workerx

import (
	"be-border-service/internal/config"
	"time"

	"github.com/hibiken/asynq"
)

type Asynqx interface {
	Run(asynq.Handler) error
}

type asynqServerConfig struct {
	RedisAddr, username, password string
	DB                            int
	WriteTimeout, ReadTimeout     time.Duration
	Concurrency                   int
}

func NewAsynqserver(config *config.Config) Asynqx {
	return &asynqServerConfig{
		RedisAddr:    config.Redis.Addr,
		username:     config.Redis.Username,
		password:     config.Redis.Password,
		ReadTimeout:  config.Redis.ReadTimeout,
		WriteTimeout: config.Redis.WriteTimeout,
		Concurrency:  config.Asynq.Concurrency,
	}
}

func (a *asynqServerConfig) Run(h asynq.Handler) error {
	srv := asynq.NewServer(asynq.RedisClientOpt{
		Addr:         a.RedisAddr,
		Username:     a.username,
		Password:     a.password,
		DB:           a.DB,
		ReadTimeout:  a.ReadTimeout,
		WriteTimeout: a.WriteTimeout,
	}, asynq.Config{
		Concurrency: a.Concurrency,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})
	return srv.Run(h)
}
