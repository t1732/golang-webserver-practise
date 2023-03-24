package schemas

import (
	"golang-webserver-practise/internal/infrastructure/dto"
	"gorm.io/gorm"
)

func MigrateUserTable(db *gorm.DB) {
	db.AutoMigrate(&dto.User{})
}

func ResetUserTable(db *gorm.DB) {
	db.Migrator().DropTable(&dto.User{})
	MigrateUserTable(db)
}
