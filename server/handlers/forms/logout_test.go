package forms_test

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/handlers/forms"
)

const (
	sessionCookieName  = "test_logout_session_cookie"
	sessionCookieValue = "test_logout"
)

func TestLogoutSuccess(t *testing.T) {
	// setup
	r, req, rw := setupLogout()

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	expectedStatusCode := http.StatusOK
	receivedStatusCode := res.StatusCode
	if receivedStatusCode != expectedStatusCode {
		t.Errorf("Expected status code: %d. Received status code: %d", expectedStatusCode, receivedStatusCode)
	}

	cookies := res.Cookies()
	index := slices.IndexFunc(cookies, func(c *http.Cookie) bool {
		return c.Name == sessionCookieName
	})
	if index == -1 {
		t.Errorf("Could not find test session cookie: %s", sessionCookieName)
	}
	c := cookies[index]
	if c.Value != "" {
		t.Errorf("Did not logout properly. Cookie has value: %+v", c.Value)
	}
}

func setupLogout() (chi.Router, *http.Request, *httptest.ResponseRecorder) {
	r := chi.NewRouter()

	logoutEndpt := handlers.LogoutEndpoint

	req := httptest.NewRequest(http.MethodPost, logoutEndpt, nil)
	rw := httptest.NewRecorder()
	http.SetCookie(rw, &http.Cookie{
		Name:  sessionCookieName,
		Value: sessionCookieValue,
		// HttpOnly: true,
		// Secure:   true,
		Path:     "/",
		MaxAge:   10,
		SameSite: http.SameSiteStrictMode,
	})
	r.Post(logoutEndpt, forms.HandleLogout(sessionCookieName).ServeHTTP)
	return r, req, rw
}
