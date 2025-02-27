package handlers

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/utils"
)

func CreateUser(userService UserService) ApiHandleFn {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse, *utils.ApiError) {
		var u User
		if err := utils.ParseJSON(r, &u); err != nil {
			return nil, utils.NewApiError(err, "Invalid json", http.StatusBadRequest)
		}

		err := userService.Insert(r.Context(), &u)
		if err != nil {
			return nil, utils.NewApiError(err, "Internal server error", http.StatusInternalServerError)
		}

		return nil, nil
	}
}
