package idler_email

import (
	"fmt"
	"github/eugene-krivtsov/idler-email/internal/model/dto"
	"github/eugene-krivtsov/idler-email/pkg/mail"
	"time"
)

func Run() {
	fmt.Println(`
 ================================================
 \\\   ######~~#####~~~##~~~~~~#####~~~#####   \\\
  \\\  ~~##~~~~##~~##~~##~~~~~~##~~~~~~##~~##   \\\
   ))) ~~##~~~~##~~##~~##~~~~~~####~~~~#####     )))
  ///  ~~##~~~~##~~##~~##~~~~~~##~~~~~~##~~##   ///
 ///   ######~~#####~~~######~~#####~~~##~~##  ///
 ================================================
	`)

	config := mail.SenderConfig{
		ServerName:    "smtp.yandex.ru:465",
		Username:      "idler.email",
		Password:      "iwnzboafqyhevgua",
		SenderName:    "Idler",
		SenderAddress: "idler.email@yandex.ru",
	}
	sender := mail.NewSmtpSender(config)

	mes := dto.RegConfirmDTO{
		User:         "Evgeny",
		RegisteredAt: time.Now().Format(time.RFC822),
		Code:         "811a54b0-e7f2-11ec-8fea-0242ac120002",
	}

	sender.Send("kia-77@mail.ru", "web/[Idler]Confirm.html", mes)
}
