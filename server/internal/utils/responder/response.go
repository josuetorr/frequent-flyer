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

func (res *AppResponse) Respond(w http.ResponseWriter, r *http.Request) {
	respond(w, r, res.Header, res.StatusCode, res.Component)
}

func newAppResponse(statusCode int, header http.Header, component templ.Component) *AppResponse {
	return &AppResponse{
		StatusCode: statusCode,
		Header:     header,
		Component:  component,
	}
}

func NewOk(h http.Header, c templ.Component) *AppResponse {
	return newAppResponse(http.StatusOK, h, c)
}

func NewCreated(h http.Header, c templ.Component) *AppResponse {
	return newAppResponse(http.StatusCreated, h, c)
}

func NewAccepted(h http.Header, c templ.Component) *AppResponse {
	return newAppResponse(http.StatusAccepted, h, c)
}

func NewNoContent(h http.Header, c templ.Component) *AppResponse {
	return newAppResponse(http.StatusNoContent, h, c)
}

func NewSeeOther(h http.Header, c templ.Component) *AppResponse {
	return newAppResponse(http.StatusSeeOther, h, c)
}
