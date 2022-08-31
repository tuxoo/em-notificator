package service

import (
	"context"
	"github/tuxoo/em-notificator/pkg/mail"
)

type Mails interface {
	Send(ctx context.Context, toEmail, path string) error
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
