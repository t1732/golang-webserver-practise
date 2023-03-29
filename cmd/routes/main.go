package main

import (
	"encoding/json"
	"fmt"
	"golang-webserver-practise/internal/interfaces/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	routes.Init(e, &gorm.DB{})

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(fmt.Errorf("routes loading error: %s \n", err))
	}
	fmt.Println(string(data))
}
