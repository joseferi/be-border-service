package usecase

import (
	"be-border-service/internal/common"
	"be-border-service/internal/constants"
	"be-border-service/pkg/logger"
	"be-border-service/pkg/workerx"
	"fmt"
	"net/http"
)

type healthCheckService struct {
	workerx workerx.AsynqClient
}

func NewHealthCheck(w workerx.AsynqClient) UseCase {
	return &healthCheckService{
		workerx: w,
	}
}

func (u *healthCheckService) Serve(_ *common.Data) common.Response {
	var (
		lf = logger.NewFields(logger.EventName("HealthCheck"))
	)

	task, err := workerx.NewTask(constants.TaskHealthCheck, nil)
	if err != nil {
		logger.Warn(fmt.Sprintf("error assign task health check got :%v", err), lf...)
		return *common.NewResponse().WithStatusCode(http.StatusBadRequest).WithMessage("bad request")
	}

	// u.workerx.Enqueue(task, workerx.WithMaxRetry(1))
	go func() {
		result, err := u.workerx.Enqueue(task)
		if err != nil {
			logger.Error(fmt.Sprintf("error enqueue task, got :%v", err), lf...)
			return
		}
		logger.Info(fmt.Sprintf("Task enqueue with ID: %s", result.ID), lf...)
	}()
	logger.Info("Success health check", lf...)
	return *common.NewResponse().WithStatusCode(http.StatusOK).WithMessage("Ok")
}
