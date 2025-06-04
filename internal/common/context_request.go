package common

import (
	"be-border-service/internal/config"
	"net/http"
)

type Data struct {
	Request *http.Request
	Config  *config.Config
}
