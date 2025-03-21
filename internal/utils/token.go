package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

const (
	tokenSep   = "."
	payloadSep = ":"
)

var (
	InvalidTokenErr     = errors.New("Invalid token")
	InvalidSignatureErr = errors.New("Invalid signature")
	ExpiredTokenErr     = errors.New("Expired token")
)

func GenerateTokenWithExpiration(userID models.ID, expiresAt int64, secret string) string {
	payload := fmt.Appendf([]byte{}, "%s%s%d", userID, payloadSep, expiresAt)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	signature := mac.Sum(nil)

	return base64.RawURLEncoding.EncodeToString(fmt.Appendf(payload, "%s%s", tokenSep, signature))
}

func GenerateToken(userID models.ID, secret string) string {
	expiresAt := time.Now().Add(time.Minute * 15).Unix()
	return GenerateTokenWithExpiration(userID, expiresAt, secret)
}

func VerifyToken(token string, secret string) (models.ID, error) {
	tokenBytes, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return "", InvalidTokenErr
	}
	token = string(tokenBytes)
	parts := strings.Split(token, tokenSep)
	if len(parts) != 2 {
		return "", InvalidTokenErr
	}

	payload := parts[0]
	signature := parts[1]

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	expectedSignature := mac.Sum(nil)
	if !hmac.Equal([]byte(signature), expectedSignature) {
		return "", InvalidSignatureErr
	}

	var userId models.ID
	var expiredAt int64
	if _, err = fmt.Sscanf(strings.Replace(payload, payloadSep, " ", 1), "%s %d", &userId, &expiredAt); err != nil {
		return "", err
	}

	if expiredAt <= time.Now().Unix() {
		return "", ExpiredTokenErr
	}

	return userId, nil
}

// NOTE: "endpoint" needs to provide '/' as it's first character
func GenerateEmailLink(endpoint string, token string) string {
	return fmt.Sprintf("%s%s/%s", GetAppHostURL(), endpoint, token)
}
