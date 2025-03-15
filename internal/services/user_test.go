package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/services"
	"github.com/josuetorr/frequent-flyer/internal/utils"
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
	u := &models.User{ID: "test_user_qwe"}

	// act
	as := services.NewUserService(ur)
	err := as.VerifyUser(ctx, u.ID)

	// assert
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error: %s. Received error: %s", expectedErr, err)
	}
}

func TestVerifyUser_UpdateFailed_Failure(t *testing.T) {
	// setup
	u := &models.User{ID: "test_user_qwe", Verified: false}
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	ur := services.NewMockUserRepository(ctrl)
	expectedErr := errors.New("Internal server error")
	ur.EXPECT().
		GetById(gomock.Eq(ctx), gomock.Any()).
		Return(u, nil)
	ur.EXPECT().
		Update(gomock.Eq(ctx), gomock.Any(), gomock.Any()).
		Return(expectedErr)

	// act
	as := services.NewUserService(ur)
	err := as.VerifyUser(ctx, u.ID)

	// assert
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error: %s. Received error: %s", expectedErr, err)
	}
}

func TestUpdatePassword_Successful(t *testing.T) {
	// setup
	u := &models.User{ID: "test_user_qwe", Password: "password"}
	newPassword := "new_password"
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	ur := services.NewMockUserRepository(ctrl)
	ur.EXPECT().
		GetById(gomock.Eq(ctx), gomock.Any()).
		Return(u, nil)
	ur.EXPECT().
		Update(gomock.Eq(ctx), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, uID string, u *models.User) error {
			bytes, _ := utils.HashPassword(newPassword)
			u.Password = string(bytes)
			return nil
		})

	// act
	as := services.NewUserService(ur)
	err := as.UpdatePassword(ctx, u.ID, "new_password")
	// assert
	if err != nil {
		t.Errorf("Expected update password to be sucessful. Error: %s", err)
	}
	if err := utils.ComparePassword(u.Password, newPassword); err != nil {
		t.Errorf("Expected updated user to have password: %s. Resulting password: %s", newPassword, u.Password)
	}
}
