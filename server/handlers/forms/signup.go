package forms

import (
	"log/slog"
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/server/handlers"
)

type SignupPostHandler struct {
	authService handlers.AuthService
	mailService handlers.MailService
}

func NewSignupHandler(authService handlers.AuthService, mailService handlers.MailService) *SignupPostHandler {
	return &SignupPostHandler{authService: authService, mailService: mailService}
}

func (h *SignupPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Form error", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("password-confirm")

	if _, err := mail.ParseAddress(email); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	if password != passwordConfirm {
		errMsg := "Passwords do not match"
		slog.Error(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	_, err := h.authService.Signup(ctx, email, password)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := h.mailService.SendVerificationEmail(ctx, email); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusCreated)
}
