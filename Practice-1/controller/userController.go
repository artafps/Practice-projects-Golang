package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersearchInput struct {
	Fname string `json:"fname"`
	Fage  string `json:"fage"`
}

func Getaddress(c echo.Context) error {
	userInput := new(UsersearchInput)
	err := c.Bind(userInput)
	if err != nil {
		return err
	}

	fmt.Println(userInput)
	return c.HTML(http.StatusOK, "<center><h1>Hello World</h1></center>")

}
func Getabout(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, `{server:api,port:3000,address:/users/get/`+id+"}")
}
