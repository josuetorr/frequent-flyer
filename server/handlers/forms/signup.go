package forms

import (
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
		w.Header().Set("HX-FOCUS", "email")
		w.WriteHeader(http.StatusBadRequest)
		errorTempl.Signup("Invalid email").Render(ctx, w)
		return
	}

	const minPasswordLen = 8
	if len(password) < minPasswordLen {
		w.Header().Add("HX-FOCUS", "password")
		w.WriteHeader(http.StatusBadRequest)
		errMsg := fmt.Sprintf("Password must be at least %d characters long", minPasswordLen)
		errorTempl.Signup(errMsg).Render(ctx, w)
		return
	}

	// TODO: clean this. It's nasty
	if password != passwordConfirm {
		w.Header().Set("HX-FOCUS", "password-confirm")
		w.WriteHeader(http.StatusBadRequest)
		errorTempl.Signup("Passwords do not match").Render(ctx, w)
		return
	}

	userID, err := h.authService.Signup(ctx, email, password)
	if err != nil {
		slog.Error("Error signing up" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		errorTempl.Signup("Oops... something went wrong").Render(ctx, w)
		return
	}

	secret := utils.GetEmailVerificationSecret()
	link := h.mailService.GenerateEmailVerificationLink(userID, secret)

	if err := h.mailService.SendVerificationEmail(ctx, link, email); err != nil {
		slog.Error("Error sending verification email" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		errorTempl.Signup("Oops... something went wrong").Render(ctx, w)
		return
	}

	w.Header().Set("HX-REDIRECT", "/login")
	w.WriteHeader(http.StatusCreated)
}
