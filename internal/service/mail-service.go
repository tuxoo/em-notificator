package service

import (
	"github/eugene-krivtsov/idler-email/internal/model/dto"
	"github/eugene-krivtsov/idler-email/pkg/mail"
	"time"
)

type MailService struct {
	Sender mail.Sender
}

func NewMailService(sender mail.Sender) *MailService {
	return &MailService{
		Sender: sender,
	}
}

func (s *MailService) Send(toEmail, path string) {
	mes := dto.RegConfirmDTO{
		User:         "Evgeny",
		RegisteredAt: time.Now().Format(time.RFC822),
		Code:         "811a54b0-e7f2-11ec-8fea-0242ac120002",
	}
	s.Sender.Send(toEmail, path, mes)
}
