package forms_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/handlers/forms"
	"go.uber.org/mock/gomock"
)

func TestPasswordResetEmailSubmission_Successful(t *testing.T) {
	// setup
	u := &models.User{
		ID:    "123",
		Email: "test@test.com",
	}
	setupMus := func(mus *handlers.MockUserService) {
		mus.EXPECT().
			GetByEmail(gomock.Any(), gomock.Eq(u.Email)).
			Return(u, nil)
	}
	setupMms := func(mms *handlers.MockMailService) {
		mms.EXPECT().
			GenerateEmailLink(gomock.Eq(u.ID), gomock.Any(), gomock.Any()).
			Return("some-link")
		mms.EXPECT().
			SendPasswordResetEmail(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil)
	}
	r, req, rw := setupPasswordResetEmailSubmission(t, u.Email, setupMus, setupMms)

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	expectedStatusCode := http.StatusOK
	receivedStatusCode := res.StatusCode
	if expectedStatusCode != receivedStatusCode {
		t.Errorf("Expected status code :%d. Received status code: %d", expectedStatusCode, receivedStatusCode)
	}
}

func TestPasswordResetEmailSubmission_InvalidEmail_Failure(t *testing.T) {
	// setup
	r, req, rw := setupPasswordResetEmailSubmission(t, "invalid_email", nil, nil)

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	// assert
	expectedStatusCode := http.StatusBadRequest
	receivedStatusCode := res.StatusCode
	if expectedStatusCode != receivedStatusCode {
		t.Errorf("Expected status code :%d. Received status code: %d", expectedStatusCode, receivedStatusCode)
	}
}

func TestPasswordResetEmailSubmission_UserNotFound_Failure(t *testing.T) {
	// setup
	setupMus := func(mms *handlers.MockUserService) {
		mms.EXPECT().
			GetByEmail(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("User not found"))
	}
	r, req, rw := setupPasswordResetEmailSubmission(t, "test@test.com", setupMus, nil)

	// act
	r.ServeHTTP(rw, req)
	res := rw.Result()

	expectedStatusCode := http.StatusNotFound
	receivedStatusCode := res.StatusCode
	if expectedStatusCode != receivedStatusCode {
		t.Errorf("Expected status code :%d. Received status code: %d", expectedStatusCode, receivedStatusCode)
	}
}

func setupPasswordResetEmailSubmission(
	t *testing.T,
	email string,
	setupMus func(*handlers.MockUserService),
	SetupMms func(*handlers.MockMailService),
) (chi.Router, *http.Request, *httptest.ResponseRecorder) {
	data := url.Values{}
	data.Add("email", email)
	pwRsetEndpt := handlers.PasswordResetEmailSubmissionEndpoint
	req := httptest.NewRequest(
		http.MethodPost,
		pwRsetEndpt,
		strings.NewReader(data.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rw := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	mus := handlers.NewMockUserService(ctrl)
	if setupMus != nil {
		setupMus(mus)
	}
	mms := handlers.NewMockMailService(ctrl)
	if SetupMms != nil {
		SetupMms(mms)
	}

	r := chi.NewRouter()
	r.Post(pwRsetEndpt, forms.HandlePasswordResetEmailSubmission(mus, mms).ServeHTTP)

	return r, req, rw
}
