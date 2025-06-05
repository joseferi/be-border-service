package workerx

import (
	"be-border-service/internal/config"
	"time"

	"github.com/hibiken/asynq"
)

type Asynqx interface {
	Register(pattern string, handler asynq.Handler)
	Run() error
}

type asynqServerConfig struct {
	RedisAddr, username, password string
	DB                            int
	WriteTimeout, ReadTimeout     time.Duration
	Concurrency                   int
	asynqServer                   *asynq.Server
	mux                           *asynq.ServeMux
}

func NewAsynqserver(config *config.Config) Asynqx {
	srv := asynq.NewServer(asynq.RedisClientOpt{
		Addr:         config.Redis.Addr,
		Username:     config.Redis.Username,
		Password:     config.Redis.Password,
		DB:           config.Redis.DB,
		ReadTimeout:  config.Redis.ReadTimeout,
		WriteTimeout: config.Redis.WriteTimeout,
	}, asynq.Config{
		Concurrency: config.Asynq.Concurrency,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})
	mux := asynq.NewServeMux()

	return &asynqServerConfig{
		mux:         mux,
		asynqServer: srv,
	}

}
func (s *asynqServerConfig) Run() error {
	return s.asynqServer.Run(s.mux)
}

func (s *asynqServerConfig) Register(pattern string, handler asynq.Handler) {
	s.mux.Handle(pattern, handler)
}
