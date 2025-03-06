package responder

import (
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

func (e *AppError) Respond(w http.ResponseWriter, r *http.Request) {
	respond(w, r, e.Header, e.StatusCode, e.Component)
}
