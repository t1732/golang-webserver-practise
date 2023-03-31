package user

import (
	"golang-webserver-practise/internal/domain/model"
	"golang-webserver-practise/internal/registory"
)

type create interface {
	Exec() (*model.User, error)
	IsUnique() (bool, error)
}

type createImpl struct {
	repo registory.Repository
	user *model.User
}

func NewCreateImpl(repo registory.Repository, user *model.User) create {
	return &createImpl{repo: repo, user: user}
}

func (impl createImpl) Exec() (*model.User, error) {
	return impl.repo.NewUserRepository().Create(impl.user)
}

func (impl createImpl) IsUnique() (bool, error) {
	return impl.repo.NewUserRepository().IsUniqueEmail(impl.user.Email)
}
