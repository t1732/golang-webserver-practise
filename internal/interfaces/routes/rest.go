package routes

import (
	"golang-webserver-practise/internal/interfaces/handler/rest"

	"github.com/labstack/echo/v4"
)

func RestRouting(e *echo.Echo) {
	healthcheckRouting(e)
}

func healthcheckRouting(e *echo.Echo) {
	r := rest.NewHealthcheck()
	e.GET("/healthcheck", r.Show)
}
