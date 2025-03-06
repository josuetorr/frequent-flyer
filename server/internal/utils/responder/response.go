package responder

import (
	"net/http"

	"github.com/a-h/templ"
)

type AppResponse struct {
	StatusCode int
	Component  templ.Component
}

func (res *AppResponse) Respond(w http.ResponseWriter, r *http.Request) {
	respond(w, r, res.StatusCode, res.Component)
}

func newAppResponse(statusCode int, component templ.Component) *AppResponse {
	return &AppResponse{
		StatusCode: statusCode,
		Component:  component,
	}
}

func NewOk(c templ.Component) *AppResponse {
	return newAppResponse(http.StatusOK, c)
}

func NewCreated(c templ.Component) *AppResponse {
	return newAppResponse(http.StatusCreated, c)
}

func NewAccepted(c templ.Component) *AppResponse {
	return newAppResponse(http.StatusAccepted, c)
}

func NewNoContent(c templ.Component) *AppResponse {
	return newAppResponse(http.StatusNoContent, c)
}

func NewSeeOther(c templ.Component) *AppResponse {
	return newAppResponse(http.StatusSeeOther, c)
}
