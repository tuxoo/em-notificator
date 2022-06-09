package idler_email

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github/eugene-krivtsov/idler-email/internal/config"
	"github/eugene-krivtsov/idler-email/internal/model/dto"
	"github/eugene-krivtsov/idler-email/pkg/mail"
	"time"
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

	mes := dto.RegConfirmDTO{
		User:         "Evgeny",
		RegisteredAt: time.Now().Format(time.RFC822),
		Code:         "811a54b0-e7f2-11ec-8fea-0242ac120002",
	}

	sender.Send("kia-77@mail.ru", "web/[Idler]Confirm.html", mes)
}
