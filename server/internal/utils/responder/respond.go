package responder

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type Responder interface {
	Respond(http.ResponseWriter, *http.Request)
}

type AppHandler func(http.ResponseWriter, *http.Request) *AppError

func respond(w http.ResponseWriter, r *http.Request, header http.Header, statusCode int, component templ.Component) {
	header.Write(w)
	w.WriteHeader(statusCode)
	if component != nil {
		// NOTE: what do we do if we get an error here?
		if err := component.Render(r.Context(), w); err != nil {
			slog.Error("Error rendering component: " + err.Error())
		}
	}
}

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		slog.Error(err.Error())
		err.Respond(w, r)
		return
	}
}
