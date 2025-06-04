package routes

import (
	"be-border-service/internal/common"
	"be-border-service/internal/config"
	"be-border-service/internal/usecase"
	"net/http"
)

// httpHandlerFunc is a contract http handler for router
type httpHandlerFunc func(request *http.Request, svc usecase.UseCase, conf *config.Config) common.Response
type Router interface {
	Route() *router
}
