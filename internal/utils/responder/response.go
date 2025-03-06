package responder

import (
	"net/http"

	"github.com/a-h/templ"
)

type AppResponse struct {
	StatusCode int
	Header     http.Header
	Component  templ.Component
}

func NewAppResponse(statusCode int, header http.Header, component templ.Component) *AppResponse {
	return &AppResponse{
		StatusCode: statusCode,
		Header:     header,
		Component:  component,
	}
}

func (res *AppResponse) Respond(w http.ResponseWriter, r *http.Request) {
	respond(w, r, res.Header, res.StatusCode, res.Component)
}
