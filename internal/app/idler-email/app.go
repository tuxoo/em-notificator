package idler_email

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github/tuxoo/idler-email/internal/config"
	"github/tuxoo/idler-email/internal/server"
	"github/tuxoo/idler-email/internal/service"
	"github/tuxoo/idler-email/internal/transport/grpc/handler"
	"github/tuxoo/idler-email/pkg/mail"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	sender := mail.NewSmtpSender(cfg.Mail)

	services := service.NewServices(sender)

	grpcHandlers := handler.NewHandler(services.MailService)
	grpcServer := server.NewGrpcServer(grpcHandlers.MailSenderHandler)

	go func() {
		if err := grpcServer.Run(cfg.Grpc.Port); err != nil {
			logrus.Errorf("error occurred while running gRPC server: %s\n", err.Error())
		}
	}()

	logrus.Print("Email notificator service has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	grpcServer.Shutdown()
}
