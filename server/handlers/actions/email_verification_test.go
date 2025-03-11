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
	req := httptest.NewRequest(http.MethodGet, handlers.VerifyEmailEndpoint+"/"+token, nil)
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
	userId := models.ID("123")
	expiresIn := time.Millisecond
	tokenSecret := "this is an invalid signature"
	r, token := setup(t, userId, expiresIn, tokenSecret, nil)
	req := httptest.NewRequest(http.MethodGet, handlers.VerifyEmailEndpoint+"/"+token, nil)
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
	userId := models.ID("123")
	expiresIn := time.Millisecond
	tokenSecret := "this is an invalid signature"
	r, token := setup(t, userId, expiresIn, tokenSecret, nil)
	req := httptest.NewRequest(http.MethodGet, handlers.VerifyEmailEndpoint+"/"+token, nil)
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
	userId := models.ID("123")
	expiresIn := time.Millisecond
	tokenSecret := "this is an invalid signature"
	r, token := setup(t, userId, expiresIn, tokenSecret, nil)
	req := httptest.NewRequest(http.MethodGet, handlers.VerifyEmailEndpoint+"/"+token, nil)
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
	userId := models.ID("123")
	expiresIn := time.Millisecond
	tokenSecret := "this is an invalid signature"
	r, token := setup(t, userId, expiresIn, tokenSecret, nil)
	req := httptest.NewRequest(http.MethodGet, handlers.VerifyEmailEndpoint+"/"+token, nil)
	rw := httptest.NewRecorder()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected a BadRequest response. Received: %s", res.Status)
	}
}

type MockConfigFunc func(handlers.UserService) *handlers.MockUserService

func setup(t *testing.T, userId models.ID, expiration time.Duration, tokenSecret string, fn MockConfigFunc) (chi.Router, string) {
	ctrl := gomock.NewController(t)
	mockUserService := handlers.NewMockUserService(ctrl)
	if fn != nil {
		mockUserService = fn(mockUserService)
	}

	expiresAt := time.Now().Add(expiration).Unix()
	token := utils.GenerateTokenWithExpiration(userId, expiresAt, tokenSecret)

	r := chi.NewRouter()
	r.Get(handlers.VerifyEmailEndpoint+"/{token}", actions.HandleEmailVerification(mockUserService).ServeHTTP)
	return r, token
}
