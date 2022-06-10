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
	"github/eugene-krivtsov/idler-email/internal/server"
	"github/eugene-krivtsov/idler-email/internal/service"
	"github/eugene-krivtsov/idler-email/internal/transport/grpc/handler"
	"github/eugene-krivtsov/idler-email/pkg/mail"
	"os"
	"os/signal"
	"syscall"
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

	grpcHandlers := handler.NewHandler(services.MailService)
	grpcServer := server.NewGrpcServer(grpcHandlers.MailSenderHandler)

	go func() {
		if err := grpcServer.Run(cfg.Grpc.Port); err != nil {
			logrus.Errorf("error occurred while running gRPC server: %s\n", err.Error())
		}
	}()

	logrus.Print("IDLER mail service application has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	grpcServer.Shutdown()

	if err := postgresDB.Close(); err != nil {
		logrus.Errorf("error occured on postgres connection close: %s", err.Error())
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logrus.Errorf("error occured on mongo connection close: %s", err.Error())
	}
}
