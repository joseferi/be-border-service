package handler

import (
	"be-border-service/internal/common"
	"be-border-service/internal/config"
	"be-border-service/internal/usecase"
	"context"
	"net/http"
)

// HttpRequest handler func wrapper
func HttpRequest(request *http.Request, svc usecase.UseCase, conf *config.Config) common.Response {
	ctx := context.Background()

	req := request.WithContext(ctx)

	data := &common.Data{
		Request:     req,
		ServiceType: "http",
		Config:      conf,
	}

	return svc.Serve(data)
}
