package dao

import "gorm.io/gorm"

type DBinfo interface {
	Ping() error
}

type dbinfoImpl struct {
	db *gorm.DB
}

func NewDBinfoImpl(db *gorm.DB) DBinfo {
	return &dbinfoImpl{db: db}
}

func (h *dbinfoImpl) Ping() error {
	sqlDB, err := h.db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		return err
	}

	return nil
}
