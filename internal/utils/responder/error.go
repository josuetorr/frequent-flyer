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

func (e *AppError) Respond(w http.ResponseWriter, r *http.Request) {
	respond(w, r, e.Header, e.StatusCode, e.Component)
}

func newAppError(e error, sc int, h http.Header, c templ.Component) *AppError {
	return &AppError{
		Message:    e.Error(),
		StatusCode: sc,
		Header:     h,
		Component:  c,
	}
}

func NewBadRequest(e error, h http.Header, c templ.Component) *AppError {
	return newAppError(e, http.StatusBadRequest, h, c)
}

func NewNotFound(e error, h http.Header, c templ.Component) *AppError {
	return newAppError(e, http.StatusNotFound, h, c)
}

func NewInternalServer(e error, h http.Header, c templ.Component) *AppError {
	return newAppError(e, http.StatusInternalServerError, h, c)
}

func NewUnsupportedMediaType(e error, h http.Header, c templ.Component) *AppError {
	return newAppError(e, http.StatusUnsupportedMediaType, h, c)
}

func NewUnauthorized(e error, h http.Header, c templ.Component) *AppError {
	return newAppError(e, http.StatusUnauthorized, h, c)
}
