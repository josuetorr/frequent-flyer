package responder

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type Responder interface {
	Respond(http.ResponseWriter, *http.Request)
}

type AppError struct {
	Message    string
	StatusCode int
	Header     http.Header
	Component  templ.Component
}

func (err AppError) Error() string {
	return err.Message
}

func NewAppError(err error, statusCode int, header http.Header, component templ.Component) *AppError {
	return &AppError{
		Message:    err.Error(),
		StatusCode: statusCode,
		Header:     header,
		Component:  component,
	}
}

func (e *AppError) Respond(w http.ResponseWriter, r *http.Request) {
	respond(w, r, e.Header, e.StatusCode, e.Component)
}

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

type AppHandler func(http.ResponseWriter, *http.Request) (*AppResponse, *AppError)

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := h(w, r)
	if err != nil {
		slog.Error(err.Error())
		err.Respond(w, r)
		return
	}

	res.Respond(w, r)
}

func respond(w http.ResponseWriter, r *http.Request, header http.Header, statusCode int, component templ.Component) {
	header.Write(w)
	w.WriteHeader(statusCode)
	if component != nil {
		component.Render(r.Context(), w)
	}
}
