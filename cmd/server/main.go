package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	infra "golang-webserver-practise/internal/infrastructure"
)

func main() {
	_, err := infra.Init()
	if err != nil {
		fmt.Println("db init error: ", err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
