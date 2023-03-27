package routes

import (
	"golang-webserver-practise/internal/interfaces/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {
	healthcheckRouting(e, db)
}

func healthcheckRouting(e *echo.Echo, db *gorm.DB) {
	e.GET("/healthcheck", handler.NewHealthcheck(db).Show)
}
