package service

import (
	"context"
	"github/eugene-krivtsov/idler-email/internal/repository/mongo"
	"github/eugene-krivtsov/idler-email/pkg/mail"
)

type Mails interface {
	Send(ctx context.Context, toEmail, path string) error
}

type Services struct {
	MailService Mails
}

func NewServices(sender mail.Sender, mongoRepositories *mongo.Repositories) *Services {
	mailService := NewMailService(sender, mongoRepositories.Mails)

	return &Services{
		MailService: mailService,
	}
}
