package email

type EmailService interface {
	SendEmail(to, subject, body string) error
}
