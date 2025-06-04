package middleware

import (
	"be-border-service/internal/common"
	"be-border-service/internal/constants"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	rsp := common.NewResponse().WithStatusCode(http.StatusNotFound)
	w.Header().Set(constants.HeaderContentTypeKey, constants.HeaderContentTypeJSON)
	w.WriteHeader(rsp.Status)
	w.Write(rsp.Byte())
	return
}
