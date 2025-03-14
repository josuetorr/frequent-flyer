package forms_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/handlers/forms"
	"go.uber.org/mock/gomock"
)

func TestSignup_Successful(t *testing.T) {
	// setup
	data := &url.Values{}
	data.Add("email", "test@test.com")
	data.Add("password", "password")
	data.Add("password-confirm", "password")
	userId := "123_id"
	tokenSecret := "test_secret"
	setupMas := func(mas *handlers.MockAuthService) {
		mas.EXPECT().
			Signup(gomock.Any(), gomock.Eq(data.Get("email")), gomock.Eq(data.Get("password"))).
			Return(userId, nil)
	}
	setupMms := func(mms *handlers.MockMailService) {
		mms.EXPECT().
			GenerateEmailLink(gomock.Eq(userId), gomock.Eq(handlers.VerifyEmailEndpoint), gomock.Eq(tokenSecret)).
			Return("some-link")
		mms.EXPECT().
			SendVerificationEmail(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil)
	}

	r, req, rw := setupSignup(t, data, tokenSecret, setupMas, setupMms)

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	expectedStatusCode := http.StatusCreated
	receivedStatusCode := res.StatusCode
	if expectedStatusCode != receivedStatusCode {
		t.Errorf("Expected status code: %d. Received status code: %d", expectedStatusCode, receivedStatusCode)
	}
	expectedHxRedirectHeader := handlers.LoginEndpoint
	receivedHxRedirectHeader := res.Header.Get("HX-REDIRECT")
	if expectedHxRedirectHeader != receivedHxRedirectHeader {
		t.Errorf("Expected HX-REDIRECT: %s. Received HX-REDIRECT: %s", expectedHxRedirectHeader, receivedHxRedirectHeader)
	}
}

func TestSignup_InvalidEmail_Failure(t *testing.T) {
	// setup
	data := &url.Values{}
	data.Add("email", "invalid_email")
	tokenSecret := "test_secret"

	r, req, rw := setupSignup(t, data, tokenSecret, nil, nil)

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	expectedStatusCode := http.StatusBadRequest
	receivedStatusCode := res.StatusCode
	if expectedStatusCode != receivedStatusCode {
		t.Errorf("Expected status code: %d. Received status code: %d", expectedStatusCode, receivedStatusCode)
	}
	expectedHxFocus := "#email"
	receivedHxFocus := res.Header.Get("HX-FOCUS")
	if expectedHxFocus != receivedHxFocus {
		t.Errorf("Expected HX-FOCUS: %s. Received HX-FOCUS: %s", expectedHxFocus, receivedHxFocus)
	}
}

func setupSignup(
	t *testing.T,
	data *url.Values,
	tokenSecret string,
	setupMas func(*handlers.MockAuthService),
	setupMms func(*handlers.MockMailService),
) (chi.Router, *http.Request, *httptest.ResponseRecorder) {
	r := chi.NewRouter()

	endpoint := handlers.SignupEndpoint

	req := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rw := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	mas := handlers.NewMockAuthService(ctrl)
	if setupMas != nil {
		setupMas(mas)
	}
	mms := handlers.NewMockMailService(ctrl)
	if setupMms != nil {
		setupMms(mms)
	}
	r.Post(endpoint, forms.HandleSignupForm(mas, mms, tokenSecret).ServeHTTP)

	return r, req, rw
}
