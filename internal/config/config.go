package config

import (
	"github.com/spf13/viper"
	. "github/tuxoo/em-notificator/pkg/mail"
	"strings"
)

type (
	Config struct {
		Mail SenderConfig
		Grpc GrpcConfig
	}

	GrpcConfig struct {
		Port int
	}
)

func Init(path string) (*Config, error) {
	viper.AutomaticEnv()

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	if err := parseEnv(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshalConfig(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0]) // folder
	viper.SetConfigName(path[1]) // config file name

	return viper.ReadInConfig()
}

func parseEnv() error {
	if err := parseMailEnv(); err != nil {
		return err
	}

	if err := parseLineEnv("grpc", "port"); err != nil {
		return err
	}

	return nil
}

func parseLineEnv(prefix, name string) error {
	viper.SetEnvPrefix(prefix)
	return viper.BindEnv(name)
}

func parseMailEnv() error {
	if err := viper.BindEnv("mail.server", "MAIL_SERVER_NAME"); err != nil {
		return err
	}

	if err := viper.BindEnv("mail.user", "MAIL_USERNAME"); err != nil {
		return err
	}

	if err := viper.BindEnv("mail.password", "MAIL_PASSWORD"); err != nil {
		return err
	}

	if err := viper.BindEnv("mail.sender.name", "MAIL_SENDER_NAME"); err != nil {
		return err
	}

	if err := viper.BindEnv("mail.sender.address", "MAIL_SENDER_ADDRESS"); err != nil {
		return err
	}

	return nil
}

func unmarshalConfig(cfg *Config) error {
	if err := viper.UnmarshalKey("grpc", &cfg.Grpc); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.Mail.ServerName = viper.GetString("mail.server")
	cfg.Mail.Username = viper.GetString("mail.user")
	cfg.Mail.Password = viper.GetString("mail.password")
	cfg.Mail.SenderAddress = viper.GetString("mail.sender.address")

	cfg.Grpc.Port = viper.GetInt("grpc.port")
}
