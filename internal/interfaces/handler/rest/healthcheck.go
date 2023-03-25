package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthcheckHandler interface {
	Show(c echo.Context) error
}

type helthcheck struct{}

func NewHealthcheck() HealthcheckHandler {
	return &helthcheck{}
}

func (_h *helthcheck) Show(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
