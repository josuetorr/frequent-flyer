package services

import (
	"log/slog"

	"github.com/josuetorr/frequent-flyer/internal/utils"
	"gopkg.in/gomail.v2"
)

type MailService struct{}

func NewMailService() *MailService {
	return &MailService{}
}

func SendVerificationEmail(to string) {
	appEmail := utils.GetAppEmail()
	appEmailPassword := utils.GetAppEmailPassword()

	m := gomail.NewMessage()
	m.SetHeader("From", appEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Account verification")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("smtp.gmail.com", 587, appEmail, appEmailPassword)

	if err := d.DialAndSend(m); err != nil {
		slog.Error(err.Error())
		return
	}
}
