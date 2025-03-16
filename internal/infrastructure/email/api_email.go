package email

import (
	"net/smtp"
	"os"
)

type SMTPEmailService struct {
	Host     string
	Port     string
	Username string
	ApiKey   string
	Sender   string
	Domain   string
}

func NewSMTPEmailService() *SMTPEmailService {
	return &SMTPEmailService{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		ApiKey:   os.Getenv("SMTP_API_KEY"),
		Sender:   os.Getenv("SMTP_SENDER"),
		Domain:   os.Getenv("SMTP_DOMAIN"),
	}
}

func (s *SMTPEmailService) SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.Username, s)
}
