package routes

import (
	"golang-webserver-practise/internal/interfaces/handler"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	healthcheckRouting(e)
}

func healthcheckRouting(e *echo.Echo) {
	e.GET("/healthcheck", handler.NewHealthcheck().Show)
}
