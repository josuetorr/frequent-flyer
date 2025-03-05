package forms

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	errorTempl "github.com/josuetorr/frequent-flyer/web/templates/errors"
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

	ctx := r.Context()
	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("password-confirm")

	if _, err := mail.ParseAddress(email); err != nil {
		var errBody bytes.Buffer
		errorTempl.Signup("Invalid email").Render(ctx, &errBody)
		w.Header().Set("HX-FOCUS", "email")
		http.Error(w, errBody.String(), http.StatusBadRequest)
		return
	}

	const minPasswordLen = 8
	if len(password) < minPasswordLen {
		errMsg := fmt.Sprintf("Password must be at least %d characters long", minPasswordLen)
		var errBody bytes.Buffer
		errorTempl.Signup(errMsg).Render(ctx, &errBody)
		w.Header().Add("HX-FOCUS", "password")
		http.Error(w, errBody.String(), http.StatusBadRequest)
		return
	}

	if password != passwordConfirm {
		errMsg := "Passwords do not match"
		var errBody bytes.Buffer
		errorTempl.Signup(errMsg).Render(ctx, &errBody)
		w.Header().Set("HX-FOCUS", "password-confirm")
		http.Error(w, errBody.String(), http.StatusBadRequest)
		return
	}

	userID, err := h.authService.Signup(ctx, email, password)
	if err != nil {
		slog.Error("Error signing up" + err.Error())
		var errBody bytes.Buffer
		errorTempl.Signup("Oops... something went wrong").Render(ctx, &errBody)
		http.Error(w, errBody.String(), http.StatusInternalServerError)
		return
	}

	secret := utils.GetEmailVerificationSecret()
	link := h.mailService.GenerateEmailVerificationLink(userID, secret)

	if err := h.mailService.SendVerificationEmail(ctx, link, email); err != nil {
		slog.Error("Error sending verification email" + err.Error())
		var errBody bytes.Buffer
		errorTempl.Signup("Oops... something went wrong").Render(ctx, &errBody)
		http.Error(w, errBody.String(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-REDIRECT", "/login")
	w.WriteHeader(http.StatusCreated)
}
