package schemas

import (
	"golang-webserver-practise/internal/infrastructure/dto"

	"gorm.io/gorm"
)

func MigrateUserTable(db *gorm.DB) error {
	if err := db.AutoMigrate(&dto.User{}); err != nil {
		return err
	}

	return nil
}

func ResetUserTable(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&dto.User{}); err != nil {
		return err
	}

	if err := MigrateUserTable(db); err != nil {
		return err
	}

	return nil
}
