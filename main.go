package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	. "github.com/harymindiar/mini/core"
	. "github.com/harymindiar/mini/routes"
)

var Kernel Kernel

func main() {

	App := new(Kernel)
	App.Environment = "dev"
	App.Echo = echo.New()

	// Middleware
	App.Echo.Use(middleware.Logger())
	App.Echo.Use(middleware.Recover())

	App.AddRoute(new(Users))

	// Start server
	App.Echo.Logger.Fatal(App.Echo.Start(":1323"))
}
