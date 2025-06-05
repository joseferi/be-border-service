package workerserver

import (
	"be-border-service/internal/config"
	"be-border-service/internal/constants"
	"be-border-service/pkg/logger"
	"be-border-service/pkg/workerx"
	"fmt"

	"github.com/hibiken/asynq"
)

func Start() {
	config := config.RegisterConfiguration()

	srv := workerx.NewAsynqserver(&config)

	if err := srv.Run(asynq.NewServeMux()); err != nil {
		logger.Warn(fmt.Sprintf("worker asynq error, err: %s", err.Error()), logger.EventName(constants.WorkerStarting))
	}
}
