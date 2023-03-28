package registory

import (
	"golang-webserver-practise/internal/domain/repository"
	"golang-webserver-practise/internal/infrastructure/dao"

	"gorm.io/gorm"
)

type Repository interface {
	NewDBinfoRepository() repository.DBinfo
	NewUserRepository() repository.User
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryImpl(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (impl *repositoryImpl) NewDBinfoRepository() repository.DBinfo {
	return dao.NewDBinfoImpl(impl.db)
}

func (impl *repositoryImpl) NewUserRepository() repository.User {
	return dao.NewUserImpl(impl.db)
}
