package email

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type EmailService struct {
	Domain string
	ApiKey string
}

func NewEmailInfrastructure() *EmailService {
	return &EmailService{
		Domain: os.Getenv("MAILGUN_DOMAIN"),
		ApiKey: os.Getenv("MAILGUN_API_KEY"),
	}
}

func (m *EmailService) SendEmail(to, subject, body string) error {
	mg := mailgun.NewMailgun(m.Domain, m.ApiKey)
	message := mg.NewMessage(
		"Mailgun Sandbox <postmaster@"+m.Domain+">",
		subject,
		body,
		to,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Printf("Message sent with ID: %s\n", id)
	return nil
}
