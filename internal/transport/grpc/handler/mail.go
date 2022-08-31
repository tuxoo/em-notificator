package handler

import (
	"context"
	"github/tuxoo/em-notificator/internal/model/entity"
	"github/tuxoo/em-notificator/internal/service"
	"github/tuxoo/em-notificator/internal/transport/grpc/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MailSenderService interface {
	SendMail(ctx context.Context, mail *entity.Mail) error
}

type MailSenderHandler struct {
	api.UnimplementedMailSenderServiceServer

	service service.Mails
}

func NewMailSenderHandler(service service.Mails) *MailSenderHandler {
	return &MailSenderHandler{
		service: service,
	}
}

func (h *MailSenderHandler) SendMail(ctx context.Context, mail *api.Mail) (*emptypb.Empty, error) {
	if err := h.service.Send(ctx, mail.Address, "web/[Idler]Confirm.html"); err != nil {
		errMessage := err.Error()
		return nil, status.Errorf(codes.Unimplemented, errMessage)
	}
	return &emptypb.Empty{}, nil
}
