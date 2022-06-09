package mail

type SenderConfig struct {
	ServerName    string
	Username      string
	Password      string
	SenderName    string
	SenderAddress string
}

type Sender interface {
	Send()
}
