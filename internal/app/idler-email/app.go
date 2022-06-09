package idler_email

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/pkg/db/mongo"
	"github.com/sirupsen/logrus"
	"github/eugene-krivtsov/idler-email/internal/config"
	mongo_repository "github/eugene-krivtsov/idler-email/internal/repository/mongo"
	"github/eugene-krivtsov/idler-email/internal/service"
	"github/eugene-krivtsov/idler-email/pkg/mail"
)

func Run(configPath string) {
	fmt.Println(`
 ================================================
 \\\   ######~~#####~~~##~~~~~~#####~~~#####   \\\
  \\\  ~~##~~~~##~~##~~##~~~~~~##~~~~~~##~~##   \\\
   ))) ~~##~~~~##~~##~~##~~~~~~####~~~~#####     )))
  ///  ~~##~~~~##~~##~~##~~~~~~##~~~~~~##~~##   ///
 ///   ######~~#####~~~######~~#####~~~##~~##  ///
 ================================================
	`)

	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	sender := mail.NewSmtpSender(cfg.Mail)

	mongoClient, err := mongo.NewMongoDb(cfg.Mongo)
	if err != nil {
		logrus.Fatalf("error initializing postgres: %s", err.Error())
	}
	mongoDB := mongoClient.Database(cfg.Mongo.DB)
	repositories := mongo_repository.NewRepositories(mongoDB)

	services := service.NewServices(sender, repositories)
	services.MailService.Send(context.Background(), "kia-77@mail.ru", "web/[Idler]Confirm.html")
}
