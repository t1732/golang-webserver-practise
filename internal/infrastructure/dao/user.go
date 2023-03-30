package dao

import (
	"golang-webserver-practise/internal/domain/model"
	"golang-webserver-practise/internal/infrastructure/dto"

	"gorm.io/gorm"
)

type User interface {
	All() (*model.Users, error)
	FindById(int64) (*model.User, error)
}

type userImpl struct {
	db *gorm.DB
}

func NewUserImpl(db *gorm.DB) User {
	return &userImpl{db: db}
}

func (impl *userImpl) All() (*model.Users, error) {
	var users dto.Users

	if err := impl.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users.ConvertToModel(), nil
}

func (impl *userImpl) FindById(id int64) (*model.User, error) {
	var user dto.User

	if err := impl.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user.ConvertToModel(), nil
}
