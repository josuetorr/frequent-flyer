package handlers

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/josuetorr/frequent-flyer/server/utils"
)

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func RefreshAccessToken() ApiHandleFn {
	return func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse, *utils.ApiError) {
		var req RefreshAccessTokenRequest
		if err := utils.ParseJSON(r, &req); err != nil {
			return nil, utils.NewApiError(err, "Invalid json", http.StatusBadRequest)
		}

		token, err := jwt.Parse(req.RefreshToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(utils.GetJwtRefreshSecret()), nil
		})
		if err != nil || !token.Valid {
			return nil, utils.NewApiError(err, "Unauthorized", http.StatusUnauthorized)
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			errMsg := "Invalid claims"
			return nil, utils.NewApiError(errors.New(errMsg), errMsg, http.StatusUnauthorized)
		}

		userId, ok := claim["sub"]
		if !ok {
			errMsg := "Invalid subject claim"
			return nil, utils.NewApiError(errors.New(errMsg), errMsg, http.StatusUnauthorized)
		}

		accessToken, err := utils.NewAccessToken(userId.(ID))
		if err != nil {
			return nil, utils.NewApiError(err, "Internal server error", http.StatusInternalServerError)
		}

		return utils.NewApiResponse(RefreshAccessTokenResponse{AccessToken: accessToken}, http.StatusOK), nil
	}
}
