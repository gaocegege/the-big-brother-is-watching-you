package api

import (	
	"github.com/emicklei/go-restful"
)

// HealthzResponse is the response type for healthz request.
type HealthzResponse struct {
	Message string `json:"message,omitempty"`
}

// Healthz is the handler for /healthz
func Healthz(request *restful.Request, response *restful.Response) {
	var getResponse HealthzResponse
	getResponse.Message = "Hello, World"
	
	response.WriteEntity(getResponse)
}