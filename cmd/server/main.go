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
	"go.uber.org/zap"
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
	db, dbErr = infra.Init(config.App().GormLogLevel())
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
	e.Debug = config.App().IsDevelopment()
	routes.Init(e, db)

	// middleware
	e.Use(loggerMiddleware())

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
	case errors.As(err, &herr):
		code = herr.Code
		msg = herr.Message.(string)
	default:
	}

	if code >= 500 {
		zap.S().Error(err)
	} else {
		zap.S().Info(err)
	}

	resp := ErrorResponse{Message: msg}
	if err := c.JSON(code, resp); err != nil {
		zap.S().Error(err)
	}
}

func loggerMiddleware() echo.MiddlewareFunc {
	var logger *zap.Logger
	var err error
	if config.App().IsDevelopment() {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(fmt.Errorf("Fatal zap new: %s \n", err))
	}

	zap.ReplaceGlobals(logger)

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:        true,
		LogURI:           true,
		LogStatus:        true,
		LogRemoteIP:      true,
		LogUserAgent:     true,
		LogProtocol:      true,
		LogLatency:       true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("protocol", v.Protocol),
				zap.Int("latency", int(v.Latency)),
				zap.String("content_length", v.ContentLength),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("user_agent", v.UserAgent),
			)

			return nil
		},
	})
}
