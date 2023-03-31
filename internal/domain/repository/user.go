package repository

import (
	"golang-webserver-practise/internal/domain/model"
)

type User interface {
	All() (*model.Users, error)
	FindById(int64) (*model.User, error)
	Create(*model.User) (*model.User, error)
	IsUniqueEmail(string) (bool, error)
}
