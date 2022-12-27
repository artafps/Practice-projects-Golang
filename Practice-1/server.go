package main

import (
	"api-serverapp/config"
	"api-serverapp/router"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config.AppConfig.Server.Port)
	server := echo.New()
	router.Setrouter(server)
	server.Start(":" + config.AppConfig.Server.Port)
}
