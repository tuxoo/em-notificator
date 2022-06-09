package main

import (
	app "github/eugene-krivtsov/idler-email/internal/app/idler-email"
)

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
