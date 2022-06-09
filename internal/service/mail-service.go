package service

import (
	"context"
	"errors"
	"github/eugene-krivtsov/idler-email/internal/model/dto"
	"github/eugene-krivtsov/idler-email/internal/model/entity"
	"github/eugene-krivtsov/idler-email/internal/repository/mongo"
	postgres_repository "github/eugene-krivtsov/idler-email/internal/repository/postgres"
	"github/eugene-krivtsov/idler-email/pkg/mail"
	"time"
)

type MailService struct {
	sender             mail.Sender
	mongoRepository    mongo_repository.Mails
	postgresRepository postgres_repository.Users
}

func NewMailService(sender mail.Sender, mongoRepository mongo_repository.Mails, postgresRepository postgres_repository.Users) *MailService {
	return &MailService{
		sender:             sender,
		mongoRepository:    mongoRepository,
		postgresRepository: postgresRepository,
	}
}

func (s *MailService) Send(ctx context.Context, toEmail, path string) error {
	user, err := s.postgresRepository.FindByEmail(toEmail)
	if err != nil {
		return err
	} else if user == nil {
		return errors.New("unknown user to confirmation")
	}

	subject := s.sender.GetSubjectByPath(path)
	mailFields := dto.RegConfirmDTO{
		User:         user.Name,
		RegisteredAt: user.RegisteredAt.Format(time.Layout),
		Code:         user.Id.String(),
	}

	messageTemplate := s.sender.FillEmailTemplate(path, mailFields)
	if err := s.sender.Send(toEmail, subject, messageTemplate); err != nil {
		return err
	}

	newMail := entity.Mail{
		Address: toEmail,
		Subject: subject,
		SentAt:  time.Now(),
	}

	return s.mongoRepository.Save(ctx, newMail)
}
