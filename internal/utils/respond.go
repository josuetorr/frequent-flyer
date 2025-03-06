package utils

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

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

type AppHandler func(http.ResponseWriter, *http.Request) (*AppResponse, *AppError)

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := h(w, r)
	if err != nil {
		slog.Error(err.Error())
		respond(w, r, err.Header, err.StatusCode, err.Component)
		return
	}

	respond(w, r, res.Header, res.StatusCode, res.Component)
}

func respond(w http.ResponseWriter, r *http.Request, header http.Header, statusCode int, component templ.Component) {
	header.Write(w)
	w.WriteHeader(statusCode)
	if component != nil {
		component.Render(r.Context(), w)
	}
}
