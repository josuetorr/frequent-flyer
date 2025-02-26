package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/utils"
)

func UpdateUser(userService UserService) ApiHandleFn[any] {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse[any], *utils.ApiError) {
		id := chi.URLParam(r, "id")
		var updatedUser User
		if err := utils.ParseJSON(r, &updatedUser); err != nil {
			return nil, utils.NewApiError(err, "Invalid json", http.StatusBadRequest)
		}

		if err := userService.Update(r.Context(), id, &updatedUser); err != nil {
			return nil, utils.NewApiError(err, "Internal server error", http.StatusInternalServerError)
		}

		return nil, nil
	}
}
