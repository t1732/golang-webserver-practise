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
	v.BindEnv("host.address")
	v.BindEnv("host.port")
	v.BindEnv("host.dbname")
	v.BindEnv("user.name")
	v.BindEnv("user.password")

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&DB); err != nil {
		return err
	}

	return nil
}
