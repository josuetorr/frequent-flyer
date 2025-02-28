package api

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/internal/utils"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(authService AuthService) ApiHandleFn {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse, *utils.ApiError) {
		var req LoginRequest
		if err := utils.ParseJSON(r, &req); err != nil {
			return nil, utils.NewApiError(err, "Invalid json", http.StatusBadRequest)
		}

		accessToken, refreshToken, err := authService.Login(r.Context(), req.Email, req.Password)
		if err != nil {
			return nil, utils.NewApiError(err, "Internal server error", http.StatusInternalServerError)
		}

		res := LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}
		return utils.NewApiResponse(res, http.StatusOK), nil
	}
}
