package forms_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/handlers/forms"
	"go.uber.org/mock/gomock"
)

func TestPasswordReset_Successful(t *testing.T) {
	// setup
	u := &models.User{ID: "123"}
	data := &url.Values{}
	data.Add("password", "new_password")
	data.Add("password-confirm", "new_password")
	setupMus := func(mus *handlers.MockUserService) {
		mus.EXPECT().
			UpdatePassword(gomock.Any(), u.ID, data.Get("password")).
			Return(nil)
	}
	r, req, rw := setupPasswordReset(t, data, setupMus)

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	expectedStatusCode := http.StatusAccepted
	receivedStatusCode := res.StatusCode
	if expectedStatusCode != receivedStatusCode {
		t.Errorf("Expected status code: %d. Received status code: %d", expectedStatusCode, receivedStatusCode)
	}
}

func setupPasswordReset(
	t *testing.T,
	data *url.Values,
	setupMus func(*handlers.MockUserService),
) (chi.Router, *http.Request, *httptest.ResponseRecorder) {
	ctrl := gomock.NewController(t)
	mus := handlers.NewMockUserService(ctrl)
	if setupMus != nil {
		setupMus(mus)
	}

	secret := "secret"
	token := utils.GenerateToken("token", secret)
	endpoint := handlers.PasswordResetEndpoint + "/" + token
	req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	rw := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Post(endpoint, forms.HandlePasswordResetSubmission(mus, secret).ServeHTTP)

	return r, req, rw
}
