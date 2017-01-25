package main

import (
	. "github.com/harymindiar/mini/core"
	. "github.com/harymindiar/mini/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	App := new(Kernel)
	App.Environment = "dev"
	App.Echo = echo.New()

	// Middleware
	App.Echo.Use(middleware.Logger())
	App.Echo.Use(middleware.Recover())

	// set container
	App.Container = new(Container)

	// add routes
	App.AddRoute(new(Users))

	// Start server
	App.Echo.Logger.Fatal(App.Echo.Start(":1323"))
}
