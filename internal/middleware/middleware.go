package middleware

import (
	"be-border-service/internal/config"
	"net/http"
)

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request, conf *config.Config) bool

func FilterFunc(w http.ResponseWriter, r *http.Request, conf *config.Config, mfs []MiddlewareFunc) bool {
	for _, mf := range mfs {
		return mf(w, r, conf)
	}

	return true
}
