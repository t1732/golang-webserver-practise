package routes

import (
	"net/http"
	"sort"

	"golang-webserver-practise/internal/config"
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
	e.POST("/users", user.Create)

	// NOTE: 開発時の routes 定義情報確認用
	if config.App().IsDevelopment() {
		e.GET("/routes", func(c echo.Context) error {
			return c.JSON(http.StatusOK, sortingRoutes(e.Routes()))
		})
	}
}

func sortingRoutes(routes []*echo.Route) []*echo.Route {
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Name < routes[j].Name
	})
	return routes
}
