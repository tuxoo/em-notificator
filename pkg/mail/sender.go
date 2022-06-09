package mail

type SenderConfig struct {
	ServerName    string
	Username      string
	Password      string
	SenderName    string
	SenderAddress string
}

type Sender interface {
	GetSubjectByPath(path string) string
	FillEmailTemplate(path string, fields any) string
	Send(toEmail, path, fields string) error
}
