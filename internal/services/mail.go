package services

import (
	"bytes"
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	emailTemplates "github.com/josuetorr/frequent-flyer/web/templates/email"
	"gopkg.in/gomail.v2"
)

type MailService struct{}

func NewMailService() *MailService {
	return &MailService{}
}

func (s *MailService) GenerateEmailLink(userID models.ID, endpoint string, secret string) string {
	token := utils.GenerateToken(userID, utils.GetTokenSecret())
	return utils.GenerateEmailLink(endpoint, token)
}

func (s *MailService) SendVerificationEmail(ctx context.Context, link string, to string) error {
	appEmail := utils.GetAppEmail()
	appEmailPassword := utils.GetAppEmailPassword()

	var body bytes.Buffer
	emailTemplates.Verification(link).Render(ctx, &body)

	m := gomail.NewMessage()
	m.SetHeader("From", appEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Account verification")
	m.SetBody("text/html", body.String())

	// NOTE: change later when we get our own smtp server
	emailHost := "smtp.gmail.com"
	d := gomail.NewDialer(emailHost, 587, appEmail, appEmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (s *MailService) SendPasswordResetEmail(ctx context.Context, link string, to string) error {
	appEmail := utils.GetAppEmail()
	appEmailPassword := utils.GetAppEmailPassword()

	var body bytes.Buffer
	emailTemplates.PasswordReset(link).Render(ctx, &body)

	m := gomail.NewMessage()
	m.SetHeader("From", appEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Password reset")
	m.SetBody("text/html", body.String())

	// NOTE: change later when we get our own smtp server
	emailHost := "smtp.gmail.com"
	d := gomail.NewDialer(emailHost, 587, appEmail, appEmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
