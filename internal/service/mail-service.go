package service

import (
	"context"
	"github/eugene-krivtsov/idler-email/internal/model/dto"
	"github/eugene-krivtsov/idler-email/internal/model/entity"
	"github/eugene-krivtsov/idler-email/internal/repository/mongo"
	"github/eugene-krivtsov/idler-email/pkg/mail"
	"time"
)

type MailService struct {
	sender     mail.Sender
	repository mongo.Mails
}

func NewMailService(sender mail.Sender, repository mongo.Mails) *MailService {
	return &MailService{
		sender:     sender,
		repository: repository,
	}
}

func (s *MailService) Send(ctx context.Context, toEmail, path string) error {
	// From Postgres by email
	mes := dto.RegConfirmDTO{
		User:         "Evgeny",
		RegisteredAt: time.Now().Format(time.RFC822),
		Code:         "811a54b0-e7f2-11ec-8fea-0242ac120002",
	}

	if err := s.sender.Send(toEmail, path, mes); err != nil {
		return err
	}

	newMail := entity.Mail{
		Address: toEmail,
		Subject: "Confirm",
		SentAt:  time.Now(),
	}

	return s.repository.Save(ctx, newMail)
}
