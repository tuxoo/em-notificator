package service

import "github/eugene-krivtsov/idler-email/pkg/mail"

type Mails interface {
	Send(toEmail, path string)
}

type Services struct {
	MailService Mails
}

func NewServices(sender mail.Sender) *Services {
	mailService := NewMailService(sender)

	return &Services{
		MailService: mailService,
	}
}
