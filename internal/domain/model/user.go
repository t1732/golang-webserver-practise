package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Email     string    `json:"email" validate:"required,email"`
	Name      string    `json:"name" validate:"required"`
}

type Users []User
