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
		panic(fmt.Errorf("Fatal error config: %s \n", err))
	}

	dsn, err := config.DB().GetDsn()
	if err != nil {
		panic(fmt.Errorf("Fatal error database dsn get: %s \n", err))
	}

	option := &infra.InitOption{Debug: true, Dsn: dsn}
	c, err := infra.Init(option)
	if err != nil {
		panic(fmt.Errorf("Fatal error database init: %s \n", err))
	}

	switch mode {
	case "apply":
		if err := c.Migrate(); err != nil {
			panic(fmt.Errorf("Fatal error db migrate: %s \n", err))
		}
	case "reset":
		if err := c.Reset(); err != nil {
			panic(fmt.Errorf("Fatal error db reset: %s \n", err))
		}
	}
}
