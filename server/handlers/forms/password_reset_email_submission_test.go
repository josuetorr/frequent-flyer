package forms_test

import (
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
	r, req, rw := setupPasswordResetEmailSubmission(
		t,
		func(mus *handlers.MockUserService) {
			mus.EXPECT().
				GetByEmail(gomock.Any(), gomock.Eq(u.Email)).
				Return(u, nil)
		},
		func(mms *handlers.MockMailService) {
			mms.EXPECT().
				GenerateEmailLink(gomock.Eq(u.ID), gomock.Any(), gomock.Any()).
				Return("some-link")
			mms.EXPECT().
				SendPasswordResetEmail(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(nil)
		},
	)

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
}

func setupPasswordResetEmailSubmission(
	t *testing.T,
	fnU func(*handlers.MockUserService),
	fnM func(*handlers.MockMailService),
) (chi.Router, *http.Request, *httptest.ResponseRecorder) {
	data := url.Values{}
	data.Add("email", "test@test.com")
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
	fnU(mus)
	mms := handlers.NewMockMailService(ctrl)
	fnM(mms)

	r := chi.NewRouter()
	r.Post(pwRsetEndpt, forms.HandlePasswordResetEmailSubmission(mus, mms).ServeHTTP)

	return r, req, rw
}
