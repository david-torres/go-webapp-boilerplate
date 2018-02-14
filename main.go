package main

import (
	"github.com/david-torres/go-webapp-quickstart/controllers"
	"github.com/david-torres/go-webapp-quickstart/services"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	// init the web server
	e := echo.New()

	// init app-wide middleware
	e.Pre(mw.RemoveTrailingSlash())
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	// security measures
	e.Use(mw.BodyLimit("2M")) // sets maximum request body size
	e.Use(mw.CSRF())          // default protection against session riding
	e.Use(mw.Secure())        // default protection against XSS, content sniffing, clickjacking, and other infections

	// init services
	helloService := services.NewHelloWorldService()

	// init controllers
	hc := controllers.NewHelloWorldController(helloService)

	// app entrypoint
	e.File("/", "public/index.html")

	// static assets
	e.Static("/static", "public/static")

	// helloWorld routes
	helloWorldRoutes := e.Group("/hello")
	helloWorldRoutes.GET("/world", hc.Hello)

	// start the server
	e.Logger.Fatal(e.Start(":3000"))
}
