package controllers

import (
	"net/http"

	"github.com/david-torres/go-webapp-boilerplate/services"
	"github.com/labstack/echo"
)

// HelloWorldController contains methods pertaining to rendering and manipulating resources via http
type HelloWorldController struct {
	hello *services.HelloWorldService
}

// NewHelloWorldController creates an instance of HelloWorldController
func NewHelloWorldController(hello *services.HelloWorldService) *HelloWorldController {
	return &HelloWorldController{hello: hello}
}

// Hello retrieves a blank
func (hc HelloWorldController) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, hc.hello.SayHello())
}
