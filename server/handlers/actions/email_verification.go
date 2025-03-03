package actions

import "net/http"

type EmailVerificationHandler struct{}

func NewEmailVerificationHandler() *EmailVerificationHandler {
	return &EmailVerificationHandler{}
}

func (h *EmailVerificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
