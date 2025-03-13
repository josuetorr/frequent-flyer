package forms_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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

var (
	shk = "12ac3if08dea4829ea292917ead4221b898abc2c091f72294f8289d8d2d2ef79"
	sbk = "08aeea1e83291ea298bbb01a2c8a6892"
)

func TestLoginForm_Successful(t *testing.T) {
	// setup
	email := "test@test.com"
	password := "secretpassword"
	s := &models.Session{ID: "123", UserID: "456"}
	r := setupLogin(t, func(mas *handlers.MockAuthService) {
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
	expectedStatusCode := http.StatusOK
	receivedStatusCode := res.StatusCode
	if expectedStatusCode != receivedStatusCode {
		t.Errorf("Expected a status of: %d. Received: %d", expectedStatusCode, receivedStatusCode)
	}
	cookies := res.Cookies()
	cookieName := "test_session_cookie"
	i := slices.IndexFunc(cookies, func(c *http.Cookie) bool {
		return c.Name == cookieName
	})
	c := cookies[i]
	decoded, _ := utils.DecodeCookie(cookieName, c.Value, shk, sbk)
	expected := fmt.Sprintf("%s:%s", s.ID, s.UserID)
	if decoded != expected {
		t.Errorf("Decoded cookie expected: %s. Received: %s", expected, decoded)
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
	sessionCookieName := "test_session_cookie"
	ctrl := gomock.NewController(t)
	mockAuthHandlers := handlers.NewMockAuthService(ctrl)
	if fn != nil {
		fn(mockAuthHandlers)
	}
	r := chi.NewRouter()
	r.Post(handlers.LoginEndpoint, forms.HandleLoginForm(sessionCookieName, mockAuthHandlers, shk, sbk).ServeHTTP)
	return r
}

type mockConfigFunc func(*handlers.MockAuthService)
