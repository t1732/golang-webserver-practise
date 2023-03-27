package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HealthcheckHandler interface {
	Show(c echo.Context) error
}

type helthcheck struct {
	db *gorm.DB
}

func NewHealthcheck(db *gorm.DB) HealthcheckHandler {
	return &helthcheck{db}
}

func (h *helthcheck) Show(c echo.Context) error {
	sqlDB, err := h.db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}
