package dao

import (
	"errors"
	"golang-webserver-practise/internal/domain/model"
	"golang-webserver-practise/internal/infrastructure/dto"

	"gorm.io/gorm"
)

type User interface {
	All() (*model.Users, error)
	FindById(int64) (*model.User, error)
	Create(*model.User) (*model.User, error)
	IsUniqueEmail(string) (bool, error)
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

// 指定したメールアドレスが既存データに存在しないか
func (impl *userImpl) IsUniqueEmail(email string) (bool, error) {
	var user = dto.User{Email: email}

	if err := impl.db.Where(&user).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

func (impl *userImpl) Create(u *model.User) (*model.User, error) {
	user := &dto.User{
		Email: u.Email,
		Name:  u.Name,
	}

	if err := impl.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user.ConvertToModel(), nil
}
