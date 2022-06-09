package idler_email

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github/eugene-krivtsov/idler-email/internal/config"
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
	services := service.NewServices(sender)

	services.MailService.Send("kia-77@mail.ru", "web/[Idler]Confirm.html")
}
