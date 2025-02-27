package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/internal/utils"
)

func GetUser(userService UserService) ApiHandleFn {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse, *utils.ApiError) {
		id := chi.URLParam(r, "id")

		u, err := userService.GetById(r.Context(), id)
		if err != nil {
			return nil, utils.NewApiError(err, "Resource not found", http.StatusNotFound)
		}

		return utils.NewApiResponse(u, int(http.StatusOK)), nil
	}
}
