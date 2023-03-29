package main

import (
	"context"
	"errors"
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
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	appEnv string
	port   string
	bindIP string
	db     *gorm.DB
)

func init() {
	appEnv := flag.String("e", "development", "environment")
	if err := config.Init(*appEnv); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var dbErr error
	db, dbErr = infra.Init(config.App.GormLogLevel())
	if dbErr != nil {
		panic(fmt.Errorf("DB init error: %s \n", dbErr))
	}

	flag.StringVar(&port, "p", "3000", "server port")
	flag.StringVar(&bindIP, "b", "", "binding ip address") // default: 0.0.0.0
}

func main() {
	flag.Parse()

	// Setup
	e := echo.New()
	e.Logger.SetLevel(config.App.LogLevel())
	routes.Init(e, db)

	// middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	e.HTTPErrorHandler = customHTTPErrorHandler

	// Start server
	fmt.Printf("running... %s mode", appEnv)
	go func() {
		if err := e.Start(bindIP + ":" + port); err != nil && err != http.ErrServerClosed {
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

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "Internal Server Error"

	var herr *echo.HTTPError
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		code = http.StatusNotFound
		msg = err.Error()
		c.Logger().Info(err)
	case errors.As(err, &herr):
		code = herr.Code
		msg = herr.Message.(string)
	default:
	}

	if code >= 500 {
		c.Logger().Error(err)
	} else {
		c.Logger().Info(err)
	}

	resp := ErrorResponse{Message: msg}
	if err := c.JSON(code, resp); err != nil {
		c.Logger().Error(err)
	}
}
