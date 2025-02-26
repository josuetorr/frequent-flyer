package handlers

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/utils"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Login(authService AuthService) ApiHandleFn {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse, *utils.ApiError) {
		return nil, nil
	}
}
