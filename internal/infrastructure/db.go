package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang-webserver-practise/internal/infrastructure/schemas"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type InitOption struct {
	Debug bool
	Dsn   string
}

type Connection struct {
	DB *gorm.DB
}

func Init(option *InitOption) (*Connection, error) {
	fmt.Println("database connecting...")

	logLv := logger.Warn
	if option.Debug {
		logLv = logger.Info
	}

	newLog := log.New(os.Stdout, "\r\n", log.LstdFlags)
	newLogger := logger.New(newLog, logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logLv,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	db, err := gorm.Open(mysql.Open(option.Dsn), &gorm.Config{
		SkipDefaultTransaction: true, // デフォルトのトランザクション機能を無効化
		PrepareStmt:            true, // プリペアードステートメントキャッシュ有効化
		Logger:                 newLogger,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("database connection done")

	return &Connection{DB: db}, nil
}

func (c *Connection) Migrate() error {
	return schemas.MigrateUserTable(c.DB)
}

func (c *Connection) Reset() error {
	return schemas.ResetUserTable(c.DB)
}
