package handlers

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/utils"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Token string `json:"token"`
}

func Signup(authService AuthService) ApiHandleFn {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse, *utils.ApiError) {
		var req SignupRequest
		if err := utils.ParseJSON(r, &req); err != nil {
			return nil, utils.NewApiError(err, "Invalid json", http.StatusBadRequest)
		}

		token, err := authService.Signup(r.Context(), req.Email, req.Password)
		if err != nil {
			return nil, utils.NewApiError(err, "Internal server error", http.StatusInternalServerError)
		}
		return utils.NewApiResponse(&SignupResponse{Token: token}, int(http.StatusCreated)), nil
	}
}
