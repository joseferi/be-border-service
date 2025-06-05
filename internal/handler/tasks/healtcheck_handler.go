package tasks

import (
	"be-border-service/pkg/logger"
	"context"

	"github.com/hibiken/asynq"
)

func HealthCheckHandler() asynq.HandlerFunc {
	return func(context.Context, *asynq.Task) error {
		logger.Info("healtcheck handler success processing")
		return nil
	}
}
