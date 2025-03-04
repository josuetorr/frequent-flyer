package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/utils"
)

const sep = "."

func GenerateEmailToken(userID models.ID, secret string) string {
	expiresAt := time.Now().Add(time.Minute * 15).Unix()
	payload := fmt.Appendf([]byte{}, "%s:%d", userID, expiresAt)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	signature := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(payload) + sep + base64.StdEncoding.EncodeToString(signature)
}

func VerifyToken(token string, secret string) (models.ID, error) {
	parts := strings.Split(token, sep)
	if len(parts) != 2 {
		return "", fmt.Errorf("Invalid token")
	}

	payloadBytes, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}

	signature, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payloadBytes)
	expectedSignature := mac.Sum(nil)
	if !hmac.Equal(signature, expectedSignature) {
		return "", fmt.Errorf("Invalid signature")
	}

	var userId models.ID
	var expiredAt int64
	if _, err = fmt.Sscanf(strings.Replace(string(payloadBytes), ":", " ", 1), "%s %d", &userId, &expiredAt); err != nil {
		return "", err
	}

	if expiredAt <= time.Now().Unix() {
		return "", fmt.Errorf("Token expired")
	}

	return userId, nil
}

func GenerateEmailVerificationLink(token string) string {
	return fmt.Sprintf("%s/verify-email/%s", utils.GetAppHostURL(), token)
}
