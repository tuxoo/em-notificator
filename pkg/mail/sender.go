package mail

type SenderConfig struct {
	ServerName    string
	Username      string
	Password      string
	SenderAddress string
}

type Sender interface {
	ParsePath(path string) (sender, subject string)
	CreateContent(toEmail, sender, subject, text string) []byte
	FillEmailTemplate(path string, fields any) string
	Send(toEmail string, content []byte) error
}
