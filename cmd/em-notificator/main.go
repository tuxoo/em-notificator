package main

import (
	app "github/tuxoo/em-notificator/internal/app/em-notificator"
)

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
