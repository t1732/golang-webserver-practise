package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang-webserver-practise/internal/config"
	infra "golang-webserver-practise/internal/infrastructure"
	"golang-webserver-practise/internal/interfaces/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// Setup
	e := echo.New()
	e.Logger.SetLevel(config.App.LogLevel())
	routes.Init(e)

	// middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	// Start server
	fmt.Printf("running... %s mode", appEnv)
	go func() {
		if err := e.Start(":" + Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// FYI: https://echo.labstack.com/cookbook/graceful-shutdown/
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
