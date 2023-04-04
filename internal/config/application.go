package config

import (
	"fmt"

	"golang-webserver-practise/pkg/envloader"
	"golang-webserver-practise/pkg/sliceutil"
)

var appEnvValues = []string{"development", "staging", "production"}

type AppConfig struct {
	env           string
	maxConnection int // 最大同時接続数
}

type AppConfigError struct {
	message string
}

// Application 設定の初期化
func initApplicationConfig(env string) error {
	if !sliceutil.ContainsChar(appEnvValues, env) {
		return &AppConfigError{fmt.Sprintf("%s is an unauthorized environmental name.", env)}
	}

	appCnf := &AppConfig{}
	appCnf.env = env
	appCnf.maxConnection = envloader.GetInt("MAX_CONNECTION", 2)
	_appCnf = *appCnf

	return nil
}

func (e *AppConfigError) Error() string {
	return e.message
}

// 実行モード名
func (a *AppConfig) Env() string {
	return a.env
}

func (a *AppConfig) IsDevelopment() bool {
	return a.env == appEnvValues[0]
}

func (a *AppConfig) IsStaging() bool {
	return a.env == appEnvValues[1]
}

func (a *AppConfig) IsProduction() bool {
	return a.env == appEnvValues[2]
}

// 最大同時接続数
func (a *AppConfig) MaxConnection() int {
	return a.maxConnection
}
