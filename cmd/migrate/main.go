package main

import (
	"flag"
	"fmt"

	"golang-webserver-practise/internal/config"
	infra "golang-webserver-practise/internal/infrastructure"
)

var (
	appEnv string
	mode   string
)

func init() {
	flag.StringVar(&appEnv, "e", "development", "environment")
	flag.StringVar(&mode, "m", "apply", "run mode (apply, reset)")
}

func main() {
	flag.Parse()

	if err := config.Init(appEnv); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	switch mode {
	case "apply":
		if err := infra.Migrate(config.App.GormLogLevel()); err != nil {
			panic(fmt.Errorf("Fatal error migrate: %s \n", err))
		}
	case "reset":
		if err := infra.Reset(config.App.GormLogLevel()); err != nil {
			panic(fmt.Errorf("Fatal error reset migrate: %s \n", err))
		}
	}
}
