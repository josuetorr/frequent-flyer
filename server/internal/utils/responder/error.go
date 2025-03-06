package responder

import (
	"net/http"

	"github.com/a-h/templ"
)

type AppError struct {
	Message    string
	StatusCode int
	Component  templ.Component
}

func (err AppError) Error() string {
	return err.Message
}

func (e *AppError) Respond(w http.ResponseWriter, r *http.Request) {
	respond(w, r, e.StatusCode, e.Component)
}

func newAppError(e error, sc int, c templ.Component) *AppError {
	return &AppError{
		Message:    e.Error(),
		StatusCode: sc,
		Component:  c,
	}
}

func NewBadRequest(e error, c templ.Component) *AppError {
	return newAppError(e, http.StatusBadRequest, c)
}

func NewUnauthorized(e error, c templ.Component) *AppError {
	return newAppError(e, http.StatusUnauthorized, c)
}

func NewNotFound(e error, c templ.Component) *AppError {
	return newAppError(e, http.StatusNotFound, c)
}

func NewUnsupportedMediaType(e error, c templ.Component) *AppError {
	return newAppError(e, http.StatusUnsupportedMediaType, c)
}

func NewInternalServer(e error, c templ.Component) *AppError {
	return newAppError(e, http.StatusInternalServerError, c)
}
