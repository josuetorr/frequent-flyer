package services

import (
	"bytes"
	"context"

	"github.com/josuetorr/frequent-flyer/internal/utils"
	emailTemplates "github.com/josuetorr/frequent-flyer/web/templates/email"
	"gopkg.in/gomail.v2"
)

type MailService struct{}

func NewMailService() *MailService {
	return &MailService{}
}

func (s *MailService) SendVerificationEmail(ctx context.Context, to string) error {
	appEmail := utils.GetAppEmail()
	appEmailPassword := utils.GetAppEmailPassword()

	var body bytes.Buffer
	emailTemplates.Verification().Render(ctx, &body)

	m := gomail.NewMessage()
	m.SetHeader("From", appEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Account verification")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, appEmail, appEmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
