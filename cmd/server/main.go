package main

import (
	"flag"
	"fmt"
	"net/http"

	"golang-webserver-practise/internal/config"
	infra "golang-webserver-practise/internal/infrastructure"

	"github.com/labstack/echo/v4"
)

var appEnv string

func main() {
	flagParse()

	if err := config.Init(appEnv); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if _, err := infra.Init(); err != nil {
		panic(fmt.Errorf("DB init error: %s \n", err))
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func flagParse() {
	flag.StringVar(&appEnv, "e", "development", "environment")
	flag.Parse()
}
