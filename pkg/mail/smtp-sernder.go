package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"path/filepath"
	"strings"
)

const (
	protocol = "tcp"
)

type SmtpSender struct {
	SenderConfig SenderConfig
}

func NewSmtpSender(senderConfig SenderConfig) *SmtpSender {
	return &SmtpSender{
		SenderConfig: senderConfig,
	}
}

func (s *SmtpSender) fillEmailTemplate(path string, fields any) string {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Panic(err)
	}

	buffer := new(bytes.Buffer)
	if err := t.Execute(buffer, fields); err != nil {
		log.Panic(err)
	}

	return buffer.String()
}

func (s *SmtpSender) createContent(toEmail string, subject string, text string) []byte {
	senderMail := mail.Address{
		Name:    s.SenderConfig.SenderName,
		Address: s.SenderConfig.SenderAddress,
	}

	receiverMail := mail.Address{
		Name:    "",
		Address: toEmail,
	}

	headers := make(map[string]string)
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""
	headers["From"] = senderMail.String()
	headers["To"] = receiverMail.String()
	headers["Subject"] = subject

	header := ""
	for k, v := range headers {
		header += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	return []byte(header + "\r\n" + text)
}

func getSubjectByPath(path string) string {
	fileName := filepath.Base(path)
	separateName := strings.Split(fileName, ".")
	return separateName[0]
}

func (s *SmtpSender) Send(toEmail, path string, fields any) error {
	subject := getSubjectByPath(path)

	messageTemplate := s.fillEmailTemplate(path, fields)
	content := s.createContent(toEmail, subject, messageTemplate)

	serverHost, _, _ := net.SplitHostPort(s.SenderConfig.ServerName)

	auth := smtp.PlainAuth("", "idler.email", s.SenderConfig.Password, serverHost)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         serverHost,
	}

	conn, err := tls.Dial(protocol, s.SenderConfig.ServerName, tlsConfig)
	if err != nil {
		return err
	}

	a, err := smtp.NewClient(conn, serverHost)
	if err != nil {
		return err
	}

	if err = a.Auth(auth); err != nil {
		return err
	}

	if err = a.Mail(s.SenderConfig.SenderAddress); err != nil {
		return err
	}

	if err = a.Rcpt(toEmail); err != nil {
		return err
	}

	w, err := a.Data()
	if err != nil {
		return err
	}

	if _, err = w.Write(content); err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	if err = a.Quit(); err != nil {
		return err
	}

	return nil
}
