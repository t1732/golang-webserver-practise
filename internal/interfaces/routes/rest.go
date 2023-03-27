package routes

import (
	"golang-webserver-practise/internal/interfaces/handler"

	"github.com/labstack/echo/v4"
)

func RestRouting(e *echo.Echo) {
	healthcheckRouting(e)
}

func healthcheckRouting(e *echo.Echo) {
	r := handler.NewHealthcheck()
	e.GET("/healthcheck", r.Show)
}
