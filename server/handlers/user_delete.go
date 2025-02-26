package handlers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/utils"
)

func DeleteUser(userService UserService) ApiHandleFn {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse, *utils.ApiError) {
		id := chi.URLParam(r, "id")
		hard := strings.ToLower(r.URL.Query().Get("hard"))

		if err := userService.Delete(r.Context(), id, hard == "true"); err != nil {
			return nil, utils.NewApiError(err, "Internal server error", http.StatusInternalServerError)
		}

		return nil, nil
	}
}
