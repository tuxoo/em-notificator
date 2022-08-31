package service

import (
	"context"
	"github/tuxoo/idler-email/internal/model/dto"
	"github/tuxoo/idler-email/pkg/mail"
)

type MailService struct {
	sender mail.Sender
}

func NewMailService(sender mail.Sender) *MailService {
	return &MailService{
		sender: sender,
	}
}

func (s *MailService) Send(ctx context.Context, toEmail, path string) error {
	sender, subject := s.sender.ParsePath(path)
	mailFields := dto.RegConfirmDTO{
		User:         "Eugen", // user.Name,
		RegisteredAt: "",      // user.RegisteredAt.Format(time.Layout),
		Code:         "",      //  user.Id.String(),
	}

	messageTemplate := s.sender.FillEmailTemplate(path, mailFields)
	content := s.sender.CreateContent(toEmail, sender, subject, messageTemplate)

	return s.sender.Send(toEmail, content)
}
