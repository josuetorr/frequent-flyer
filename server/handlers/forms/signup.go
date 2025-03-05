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

	const minPasswordLen = 8
	if len(password) < minPasswordLen {
		errMsg := fmt.Sprintf("Password must be at least %d characters long", minPasswordLen)
		errorTempl.Signup(errMsg).Render(ctx, w)
		return
	}

	if _, err := mail.ParseAddress(email); err != nil {
		slog.Error("Invalid email")
		errorTempl.Signup("Invalid email").Render(ctx, w)
		return
	}

	if password != passwordConfirm {
		slog.Error("Passwords do not match")
		errorTempl.Signup("Passwords do not match").Render(ctx, w)
		return
	}

	userID, err := h.authService.Signup(ctx, email, password)
	if err != nil {
		slog.Error("Error signing up" + err.Error())
		errorTempl.Signup("Oops... something went wrong").Render(ctx, w)
		return
	}

	secret := utils.GetEmailVerificationSecret()
	link := h.mailService.GenerateEmailVerificationLink(userID, secret)

	if err := h.mailService.SendVerificationEmail(ctx, link, email); err != nil {
		slog.Error("Error sending verification email" + err.Error())
		errorTempl.Signup("Oops... something went wrong").Render(ctx, w)
		return
	}

	w.Header().Set("HX-REDIRECT", "/login")
	w.WriteHeader(http.StatusCreated)
}
