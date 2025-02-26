package handlers

import (
	"errors"
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/utils"
)

func CreateUser(userService UserService) ApiHandleFn[any] {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse[any], *utils.ApiError) {
		ct := r.Header.Get("Content-Type")
		if ct != "application/json" {
			http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
			return nil, utils.NewApiError(
				errors.New("Unsupported Media Type"), "Unsupported Media Type", http.StatusUnsupportedMediaType)
		}
		defer r.Body.Close()

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
