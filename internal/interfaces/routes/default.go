package routes

import (
	"golang-webserver-practise/internal/interfaces/handler"
	"golang-webserver-practise/internal/registory"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {
	healthcheckRouting(e, db)
}

func healthcheckRouting(e *echo.Echo, db *gorm.DB) {
	repo := registory.NewRepositoryImpl(db)

	e.GET("/healthcheck", handler.NewHealthcheckImpl(repo).Show)

	user := handler.NewUserImpl(repo)
	e.GET("/users", user.Index)
	e.GET("/users/:id", user.Show)
}
