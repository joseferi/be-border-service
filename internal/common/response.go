package common

import (
	"encoding/json"
	"sync"
)

type Response struct {
	Status  int    `json:"statusCode"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

var (
	oneResp sync.Once
	rsp     *Response
)

func NewResponse() *Response {
	oneResp.Do(func() {
		rsp = &Response{
			Status:  200,
			Message: "Success",
		}
	})
	x := *rsp
	return &x
}

func (r *Response) WithMessage(message string) *Response {
	r.Message = message
	return r
}
func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) WithError(err any) *Response {
	r.Errors = err
	return r
}

func (r *Response) WithStatusCode(statusCode int) *Response {
	r.Status = statusCode
	return r
}
func (r *Response) Byte() []byte {

	b, _ := json.Marshal(r)
	return b
}
