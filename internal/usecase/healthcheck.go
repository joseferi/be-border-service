package usecase

import (
	"be-border-service/internal/common"
	"be-border-service/pkg/logger"
	"net/http"
)

type healthCheckService struct {
}

func NewHealthCheck() UseCase {
	return &healthCheckService{}
}

func (u *healthCheckService) Serve(data *common.Data) common.Response {
	logger.Info("Success health check")
	return *common.NewResponse().WithStatusCode(http.StatusOK).WithMessage("Ok")
}
