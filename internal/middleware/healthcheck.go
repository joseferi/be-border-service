package middleware

import (
	"be-border-service/internal/common"
	"be-border-service/internal/config"
	"be-border-service/internal/constants"
	"net/http"
)

func HealthCheckMiddleware() MiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request, conf *config.Config) bool {
		resp := common.NewResponse().WithStatusCode(http.StatusUnauthorized).WithMessage("unauthorize")
		w.Header().Set(constants.HeaderContentTypeKey, constants.HeaderContentTypeJSON)
		w.WriteHeader(resp.Status)
		w.Write(resp.Byte())
		return false
	}
}
