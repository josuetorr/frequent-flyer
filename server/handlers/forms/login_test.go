package forms_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/services"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/handlers/forms"
	"go.uber.org/mock/gomock"
)

func TestLoginForm_Successful(t *testing.T) {
	// setup
	email := "test@test.com"
	password := "secretpassword"
	r := setupLogin(t, func(mas *handlers.MockAuthService) {
		s := &models.Session{ID: "123", UserID: "456"}
		mas.EXPECT().
			Login(gomock.Any(), gomock.Eq(email), gomock.Eq(password)).
			Return(s, nil)
	})
	data := url.Values{}
	data.Add("email", email)
	data.Add("password", password)
	req := httptest.NewRequest(http.MethodPost, handlers.LoginEndpoint, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rw := httptest.NewRecorder()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected a status of: %d. Received: %d", http.StatusOK, res.StatusCode)
	}
	cookies := res.Cookies()
	i := slices.IndexFunc(cookies, func(c *http.Cookie) bool {
		return c.Name == "test_session_cookie"
	})
	c := cookies[i]
	decoded, _ := utils.DecodeCookie("test_session_cookie", c.Value)
	if decoded != "123:456" {
		t.Errorf("Decoded cookie expected: 123:456. Received: %s", decoded)
	}
}

func TestLoginForm_InvalidContentType_Failure(t *testing.T) {
	// setup
	r := setupLogin(t, nil)
	req := httptest.NewRequest(http.MethodPost, handlers.LoginEndpoint, nil)
	rw := httptest.NewRecorder()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	expectedStatusCode := http.StatusUnsupportedMediaType
	if res.StatusCode != expectedStatusCode {
		t.Errorf("Expected a status of: %d. Received: %d", expectedStatusCode, res.StatusCode)
	}
}

func TestLoginForm_InvalidEmail_Failure(t *testing.T) {
	// setup
	data := url.Values{}
	data.Add("email", "malformed_email")
	r := setupLogin(t, nil)
	req := httptest.NewRequest(http.MethodPost, handlers.LoginEndpoint, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rw := httptest.NewRecorder()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	expectedStatusCode := http.StatusBadRequest
	if res.StatusCode != expectedStatusCode {
		t.Errorf("Expected a status of: %d. Received: %d", expectedStatusCode, res.StatusCode)
	}
}

func TestLoginForm_InvalidCredentials_Failure(t *testing.T) {
	// setup
	data := url.Values{}
	data.Add("email", "test@test.com")
	r := setupLogin(t, func(mas *handlers.MockAuthService) {
		mas.EXPECT().
			Login(gomock.Any(), gomock.Eq(data.Get("email")), gomock.Any()).
			Return(nil, services.InvalidCredentialError)
	})
	req := httptest.NewRequest(http.MethodPost, handlers.LoginEndpoint, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rw := httptest.NewRecorder()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	expectedStatusCode := http.StatusBadRequest
	if res.StatusCode != expectedStatusCode {
		t.Errorf("Expected a status of: %d. Received: %d", expectedStatusCode, res.StatusCode)
	}
}

func setupLogin(t *testing.T, fn mockConfigFunc) chi.Router {
	os.Setenv("SESSION_HASH_KEY", "772soie213s2k2lw3op42e14f20c44de8b23b75c091f72294f8289d8d2d2ef79")
	os.Setenv("SESSION_BLOCK_KEY", "82mw8201aspw2l8x8p09801a2c8a6892")
	sessionCookieName := "test_session_cookie"
	ctrl := gomock.NewController(t)
	mockAuthHandlers := handlers.NewMockAuthService(ctrl)
	if fn != nil {
		fn(mockAuthHandlers)
	}
	r := chi.NewRouter()
	r.Post(handlers.LoginEndpoint, forms.HandleLoginForm(sessionCookieName, mockAuthHandlers).ServeHTTP)
	return r
}

type mockConfigFunc func(*handlers.MockAuthService)
