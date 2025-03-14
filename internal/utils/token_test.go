package utils_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/josuetorr/frequent-flyer/internal/utils"
)

func TestValidToken(t *testing.T) {
	secret := "bob"
	expectedUserId := "123"
	token := utils.GenerateToken(expectedUserId, secret)
	result, err := utils.VerifyToken(token, secret)
	if err != nil {
		t.Error(err)
	}

	if result != expectedUserId {
		t.Errorf("expected userID: %s. Received: %s", expectedUserId, result)
	}
}

func TestInvalidSignedToken(t *testing.T) {
	secret := "bob"
	invalidSecret := "obo"
	expectedUserId := "123"

	token := utils.GenerateToken(expectedUserId, secret)
	result, err := utils.VerifyToken(token, invalidSecret)
	if err == nil {
		t.Error("Verify should have returning an error since signatures should not match")
	}

	if result == expectedUserId {
		t.Errorf("expected userID: %s. Received: %s. Results should be different", expectedUserId, result)
	}
}

func TestLinkGeneration(t *testing.T) {
	const hostURLEnvValue = "APP_HOST_URL"
	os.Setenv(hostURLEnvValue, "localhost:3000")
	secret := "bob"
	userID := "123"
	token := utils.GenerateToken(userID, secret)
	expectedLink := fmt.Sprintf("%s/verify-email/%s", os.Getenv(hostURLEnvValue), token)
	link := utils.GenerateEmailLink("verify-email", token)

	if link != expectedLink {
		t.Errorf("Expected: %s. Received: %s", expectedLink, link)
	}
}
