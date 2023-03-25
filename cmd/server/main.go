package main

import (
	"flag"
	"fmt"

	"golang-webserver-practise/internal/config"
	infra "golang-webserver-practise/internal/infrastructure"
	"golang-webserver-practise/internal/interfaces/routes"

	"github.com/labstack/echo/v4"
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

	if _, err := infra.Init(config.App.GormLogLevel()); err != nil {
		panic(fmt.Errorf("DB init error: %s \n", err))
	}

	e := echo.New()
	routes.RestRouting(e)

	fmt.Printf("running... mode:%s", appEnv)
	e.Logger.Fatal(e.Start(":" + Port))
}
