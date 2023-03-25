package infrastructure

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"golang-webserver-practise/internal/config"
	"golang-webserver-practise/internal/infrastructure/schemas"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

func Init(llv logger.LogLevel) (*gorm.DB, error) {
	fmt.Println("database connecting...")

	env := os.Getenv("SERVER_ENV")
	if env == "" {
		env = "development"
	}

	t, err := template.
		New("dsn").
		Parse("{{.User.Name}}:{{.User.Password}}@tcp({{.Host.Address}}:{{.Host.Port}})/{{.Host.DBname}}?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err = t.Execute(&b, config.DB); err != nil {
		return nil, err
	}

	newLog := log.New(os.Stdout, "\r\n", log.LstdFlags)
	newLogger := logger.New(newLog, logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  llv,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	db, err = gorm.Open(mysql.Open(b.String()), &gorm.Config{
		SkipDefaultTransaction: true, // デフォルトのトランザクション機能を無効化
		PrepareStmt:            true, // プリペアードステートメントキャッシュ有効化
		Logger:                 newLogger,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("database connection done")

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}

func Migrate(llv logger.LogLevel) {
	Init(llv)
	schemas.MigrateUserTable(db)
}

func Reset(llv logger.LogLevel) {
	Init(llv)
	schemas.ResetUserTable(db)
}
