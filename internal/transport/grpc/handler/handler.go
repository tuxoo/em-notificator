package handler

import "github/tuxoo/em-notificator/internal/service"

type Handler struct {
	MailSenderHandler *MailSenderHandler
}

func NewHandler(mailService service.Mails) *Handler {
	mailHandler := NewMailSenderHandler(mailService)
	return &Handler{
		MailSenderHandler: mailHandler,
	}
}
