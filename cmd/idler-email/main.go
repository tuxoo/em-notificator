package main

import (
	app "github/tuxoo/idler-email/internal/app/idler-email"
)

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
