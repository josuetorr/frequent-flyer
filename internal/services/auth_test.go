package services_test

import (
	"context"
	"testing"

	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/services"
	"go.uber.org/mock/gomock"
)

func TestSignup_Successful(t *testing.T) {
	// setup
	u := &models.User{ID: "test_123"}
	ctrl := gomock.NewController(t)
	mur := services.NewMockUserRepository(ctrl)
	mur.EXPECT().
		GetByEmail(gomock.Any(), gomock.Any()).
		Return(nil, nil)
	mur.EXPECT().
		Insert(gomock.Any(), gomock.Any()).
		Return(u, nil)
	msr := services.NewMockSessionRepository(ctrl)

	// act
	as := services.NewAuthService(mur, msr)
	res, err := as.Signup(context.Background(), "test@test.com", "password")
	// assert
	if err != nil {
		t.Errorf("Failed to signup. Error: %s", err)
	}
	if res != u.ID {
		t.Errorf("Expected userID: %s. Received userID: %s", res, u.ID)
	}
}
