package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang-webserver-practise/internal/config"
	"golang-webserver-practise/internal/infrastructure/schemas"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(llv logger.LogLevel) (*gorm.DB, error) {
	fmt.Println("database connecting...")

	newLog := log.New(os.Stdout, "\r\n", log.LstdFlags)
	newLogger := logger.New(newLog, logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  llv,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	dsn, err := config.DB().GetDsn()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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

func Migrate(llv logger.LogLevel) error {
	db, err := Init(llv)
	if err != nil {
		return err
	}
	return schemas.MigrateUserTable(db)
}

func Reset(llv logger.LogLevel) error {
	db, err := Init(llv)
	if err != nil {
		return err
	}
	return schemas.ResetUserTable(db)
}
