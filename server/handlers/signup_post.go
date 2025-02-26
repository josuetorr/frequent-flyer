package handlers

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/utils"
)

func Signup(authService AuthService) ApiHandleFn[any] {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse[any], *utils.ApiError) {
		return nil, nil
	}
}
