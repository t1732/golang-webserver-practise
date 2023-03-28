package registory

import (
	"golang-webserver-practise/internal/domain/repository"
	"golang-webserver-practise/internal/infrastructure/dao"

	"gorm.io/gorm"
)

type Repository interface {
	NewDBinfoRepository() repository.DBinfo
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryImpl(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) NewDBinfoRepository() repository.DBinfo {
	return dao.NewDBinfoImpl(r.db)
}
