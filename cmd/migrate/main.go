package main

import (
	"golang-webserver-practise/internal/config"
	infra "golang-webserver-practise/internal/infrastructure"

	"gorm.io/gorm/logger"
)

func main() {
	logLevel := logger.Warn
	if config.App.IsDevelopment() {
		logLevel = logger.Info
	}
	infra.Reset(logLevel)
	// infra.Migrate()
}
