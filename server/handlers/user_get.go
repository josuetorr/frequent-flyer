package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/utils"
)

func GetUser(userService UserService) ApiHandleFn[User] {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse[User], *utils.ApiError) {
		id := chi.URLParam(r, "id")

		u, err := userService.Get(r.Context(), id)
		if err != nil {
			return nil, utils.NewApiError(err, "Resource not found", http.StatusNotFound)
		}

		return utils.NewApiResponse(u, int(http.StatusOK)), nil
	}
}
