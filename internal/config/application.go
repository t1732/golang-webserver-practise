package config

import (
	"fmt"

	"golang-webserver-practise/pkg/sliceutil"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm/logger"
)

var appEnvValues = []string{"development", "staging", "production"}

type AppConfig struct {
	Env string
}

type AppEnvError struct {
	message string
}

// Application 設定の初期化
func initApplicationConfig(env string) error {
	if !sliceutil.ContainsChar(appEnvValues, env) {
		return &AppEnvError{fmt.Sprintf("%s is an unauthorized environmental name.", env)}
	}

	appCnf := &AppConfig{}
	appCnf.Env = env
	App = *appCnf

	return nil
}

func (e *AppEnvError) Error() string {
	return e.message
}

func (a *AppConfig) IsDevelopment() bool {
	return a.Env == appEnvValues[0]
}

func (a *AppConfig) IsStaging() bool {
	return a.Env == appEnvValues[1]
}

func (a *AppConfig) IsProduction() bool {
	return a.Env == appEnvValues[2]
}

func (a *AppConfig) LogLevel() log.Lvl {
	if a.IsDevelopment() {
		return log.DEBUG
	}

	return log.WARN
}

func (a *AppConfig) GormLogLevel() logger.LogLevel {
	if a.IsDevelopment() {
		return logger.Info
	}

	return logger.Warn
}
