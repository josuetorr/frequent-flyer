package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/utils"
)

func GenerateEmailToken(userID models.ID) string {
	expiresAt := time.Now().Add(time.Minute * 15).Unix()
	payload := fmt.Appendf([]byte{}, "%s:%d", userID, expiresAt)

	mac := hmac.New(sha256.New, []byte(utils.GetEmailVerificationSecret()))
	mac.Write(payload)
	signature := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(payload) + "." + base64.StdEncoding.EncodeToString(signature)
}
