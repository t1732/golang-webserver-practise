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

type Users []User

func (u *User) ConvertToModel() *model.User {
	return &model.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Email:     u.Email,
		Name:      u.Name,
	}
}

func (users Users) ConvertToModel() *model.Users {
	result := make(model.Users, len(users))

	for i, u := range users {
		result[i] = *u.ConvertToModel()
	}

	return &result
}
