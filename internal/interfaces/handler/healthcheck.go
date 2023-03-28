package handler

import (
	"net/http"

	"golang-webserver-practise/internal/registory"

	"github.com/labstack/echo/v4"
)

type HealthcheckHandler interface {
	Show(c echo.Context) error
}

type helthcheckImpl struct {
	repo registory.Repository
}

func NewHealthcheckImpl(repo registory.Repository) HealthcheckHandler {
	return &helthcheckImpl{repo}
}

func (h *helthcheckImpl) Show(c echo.Context) error {
	if err := h.repo.NewDBinfoRepository().Ping(); err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}
