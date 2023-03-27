package config

import (
	"strings"

	"github.com/spf13/viper"
)

type dbConfig struct {
	Host struct {
		Address string
		Port    int
		DBname  string
	}
	User struct {
		Name     string
		Password string
	}
}

func initDatabaseConfig() error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("database")
	v.AddConfigPath("config")

	// 環境変数に DB_ prefix を追加
	v.SetEnvPrefix("db")
	// 環境変数名をアンダーバー区切りにする
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 環境変数がある場合はそれを優先する
	if err := v.BindEnv("host.address"); err != nil {
		return err
	}
	if err := v.BindEnv("host.port"); err != nil {
		return err
	}
	if err := v.BindEnv("host.dbname"); err != nil {
		return err
	}
	if err := v.BindEnv("user.name"); err != nil {
		return err
	}
	if err := v.BindEnv("user.password"); err != nil {
		return err
	}

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&DB); err != nil {
		return err
	}

	return nil
}
