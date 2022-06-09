package idler_email

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/pkg/db/mongo"
	"github.com/eugene-krivtsov/idler/pkg/db/postgres"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github/eugene-krivtsov/idler-email/internal/config"
	mongorepository "github/eugene-krivtsov/idler-email/internal/repository/mongo"
	postgresrepository "github/eugene-krivtsov/idler-email/internal/repository/postgres"
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

	postgresDB, err := postgres.NewPostgresDB(postgres.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		DB:       cfg.Postgres.DB,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		logrus.Fatalf("error initializing postgres: %s", err.Error())
	}

	mongoClient, err := mongo.NewMongoDb(cfg.Mongo)
	if err != nil {
		logrus.Fatalf("error initializing postgres: %s", err.Error())
	}
	mongoDB := mongoClient.Database(cfg.Mongo.DB)
	postgresRepositories := postgresrepository.NewRepositories(postgresDB)
	mongoRepositories := mongorepository.NewRepositories(mongoDB)

	services := service.NewServices(service.ServicesDepends{
		PostgresRepositories: postgresRepositories,
		MongoRepositories:    mongoRepositories,
		Sender:               sender,
	})

	services.MailService.Send(context.Background(), "kia-77@mail.ru", "web/[Idler]Confirm.html")
}
