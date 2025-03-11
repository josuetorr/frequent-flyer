package actions_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/handlers/actions"
	"go.uber.org/mock/gomock"
)

// TODO: Use a setup func
// TODO: Add a common flow for endpoints (not sure what flow to use yet)
func TestHandleEmailVerification_Successful(t *testing.T) {
	// setup
	testUser := &models.User{ID: "123"}
	ctrl := gomock.NewController(t)
	mockUserService := handlers.NewMockUserService(ctrl)
	mockUserService.EXPECT().
		GetById(gomock.Any(), gomock.Eq(testUser.ID)).
		Return(testUser, nil)
	mockUserService.EXPECT().
		VerifyUser(gomock.Any(), gomock.Eq(testUser.ID)).
		Return(nil)

	r := chi.NewRouter()
	r.Get(handlers.VerifyEmailEndpoint+"/{token}", actions.HandleEmailVerification(mockUserService).ServeHTTP)
	token := utils.GenerateToken("123", utils.GetTokenSecret())
	req := httptest.NewRequest(http.MethodGet, "/verify-email/"+token, nil)
	rw := httptest.NewRecorder()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	if res.StatusCode != http.StatusSeeOther {
		t.Errorf("Expected a redirection to /login. Received: %s", res.Header.Get("Location"))
	}
}

func TestHandleEmailVerification_WhenInvalidToken_Failure(t *testing.T) {
	// setup
	ctrl := gomock.NewController(t)
	mockUserService := handlers.NewMockUserService(ctrl)

	r := chi.NewRouter()
	r.Get(handlers.VerifyEmailEndpoint+"/{token}", actions.HandleEmailVerification(mockUserService).ServeHTTP)
	token := "thisisaninvalidtoken"
	req := httptest.NewRequest(http.MethodGet, "/verify-email/"+token, nil)
	rw := httptest.NewRecorder()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected a BadRequest response. Received: %s", res.Status)
	}
}

func TestHandleEmailVerification_WhenInvalidSignature_Failure(t *testing.T) {
	// setup
	ctrl := gomock.NewController(t)
	mockUserService := handlers.NewMockUserService(ctrl)

	r := chi.NewRouter()
	r.Get(handlers.VerifyEmailEndpoint+"/{token}", actions.HandleEmailVerification(mockUserService).ServeHTTP)
	token := utils.GenerateToken("123", "this is an invalid signature")
	req := httptest.NewRequest(http.MethodGet, "/verify-email/"+token, nil)
	rw := httptest.NewRecorder()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected a BadRequest response. Received: %s", res.Status)
	}
}

func TestHandleEmailVerification_WhenExpiredToken_Failure(t *testing.T) {
	// setup
	ctrl := gomock.NewController(t)
	mockUserService := handlers.NewMockUserService(ctrl)

	r := chi.NewRouter()
	r.Get(handlers.VerifyEmailEndpoint+"/{token}", actions.HandleEmailVerification(mockUserService).ServeHTTP)
	expiresAt := time.Now().Add(time.Microsecond).Unix()
	token := utils.GenerateTokenWithExpiration("123", expiresAt, "this is an invalid signature")
	req := httptest.NewRequest(http.MethodGet, "/verify-email/"+token, nil)
	rw := httptest.NewRecorder()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected a BadRequest response. Received: %s", res.Status)
	}
}

// NOTE: we are here
func TestHandleEmailVerification_WhenUserNotFound_Failure(t *testing.T) {
	// setup
	ctrl := gomock.NewController(t)
	mockUserService := handlers.NewMockUserService(ctrl)

	r := chi.NewRouter()
	r.Get(handlers.VerifyEmailEndpoint+"/{token}", actions.HandleEmailVerification(mockUserService).ServeHTTP)
	expiresAt := time.Now().Add(time.Microsecond).Unix()
	token := utils.GenerateTokenWithExpiration("123", expiresAt, "this is an invalid signature")
	req := httptest.NewRequest(http.MethodGet, "/verify-email/"+token, nil)
	rw := httptest.NewRecorder()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected a BadRequest response. Received: %s", res.Status)
	}
}
