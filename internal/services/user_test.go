package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/services"
	"go.uber.org/mock/gomock"
)

func TestVerifyUser_Successful(t *testing.T) {
	// setup
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	ur := services.NewMockUserRepository(ctrl)
	u := &models.User{ID: "test_user_qwe", Verified: false}
	ur.EXPECT().
		GetById(gomock.Eq(ctx), gomock.Eq(u.ID)).
		Return(u, nil)
	ur.EXPECT().
		Update(gomock.Eq(ctx), gomock.Eq(u.ID), gomock.Eq(u)).
		Do(func(_ context.Context, _ string, u *models.User) {
			u.Verified = true
		}).
		Return(nil)

	// act
	us := services.NewUserService(ur)
	err := us.VerifyUser(ctx, u.ID)
	// assert
	if err != nil {
		t.Logf("Error verifying user: %s", err)
	}
}

func TestVerifyUser_UserNotFound_Failure(t *testing.T) {
	// setup
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	ur := services.NewMockUserRepository(ctrl)
	expectedErr := errors.New("User not found")
	ur.EXPECT().
		GetById(gomock.Eq(ctx), gomock.Any()).
		Return(nil, expectedErr)
	// TODO: write behaviour
	u := &models.User{ID: "test_user_qwe", Verified: false}

	// act
	as := services.NewUserService(ur)
	err := as.VerifyUser(ctx, u.ID)

	// assert
	if expectedErr != err {
		t.Errorf("Expected error: %s. Received error: %s", expectedErr, err)
	}
}
