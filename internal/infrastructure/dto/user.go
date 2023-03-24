package dto

import (
	"time"

	"golang-webserver-practise/internal/domain/model"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `gorm:"size:255;index:,unique"`
	Name      string `gorm:"size:255;index"`
}

type Users []Users

func (u *User) ConvertToModel() *model.User {
	return &model.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Email:     u.Email,
		Name:      u.Name,
	}
}
