package router

import (
	"api-serverapp/controller"

	"github.com/labstack/echo/v4"
)

func Setrouter(e *echo.Echo) error {
	e.GET("/", controller.Getaddress)
	e.GET("/:id", controller.Getabout)

	return nil
}
