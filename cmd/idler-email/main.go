package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

func main() {

	sendmail("kia-77@mail.ru", "Subject: Test email from Go!\n", "<html><body><h2>Hello World!</h2></body></html>")
}

func sendmail(toEmail string, subj string, body string) {

	from := mail.Address{Name: "Idler", Address: "idler.email@yandex.ru"}
	to := mail.Address{Name: "", Address: toEmail}

	headers := make(map[string]string)
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	servername := "smtp.yandex.ru:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", "idler.email", "iwnzboafqyhevgua", host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	a, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = a.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = a.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = a.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := a.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	a.Quit()

}
