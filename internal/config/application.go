package config

import (
	"fmt"

	"golang-webserver-practise/pkg/sliceutil"
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

	return nil
}

func (e *AppEnvError) Error() string {
	return e.message
}

func (a *AppConfig) IsDevelopment() bool {
	return App.Env == appEnvValues[0]
}

func (a *AppConfig) IsStaging() bool {
	return App.Env == appEnvValues[1]
}

func (a *AppConfig) IsProduction() bool {
	return App.Env == appEnvValues[2]
}
