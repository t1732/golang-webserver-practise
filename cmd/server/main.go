package main

import (
	"flag"
	"fmt"
	"net/http"

	"golang-webserver-practise/internal/config"
	infra "golang-webserver-practise/internal/infrastructure"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/logger"
)

var (
	appEnv string
	Port   string
)

func init() {
	flag.StringVar(&appEnv, "e", "development", "environment")
	flag.StringVar(&Port, "p", "3000", "server port")
}

func main() {
	flag.Parse()

	if err := config.Init(appEnv); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	logLevel := logger.Warn
	if config.App.IsDevelopment() {
		logLevel = logger.Info
	}
	if _, err := infra.Init(logLevel); err != nil {
		panic(fmt.Errorf("DB init error: %s \n", err))
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Printf("running... mode:%s", appEnv)
	e.Logger.Fatal(e.Start(":" + Port))
}
