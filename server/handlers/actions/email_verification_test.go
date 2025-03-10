package actions

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"go.uber.org/mock/gomock"
)

// TODO: Test failures
// TODO: Use a setup func
// TODO: Add a common flow for endpoints (not sure what flow to use yet)
func TestHandleEmailVerification_Successful(t *testing.T) {
	// setup
	testUser := &models.User{ID: "123"}
	ctrl := gomock.NewController(t)
	mockUserService := handlers.NewMockUserService(ctrl)
	mockUserService.EXPECT().
		GetById(gomock.Any(), testUser.ID).
		Return(testUser, nil)
	mockUserService.EXPECT().
		VerifyUser(gomock.Any(), testUser.ID).
		Return(nil)

	r := chi.NewRouter()
	r.Get("/verify-email/{token}", HandleEmailVerification(mockUserService).ServeHTTP)
	token := utils.GenerateToken("123", utils.GetEmailSecret())
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
