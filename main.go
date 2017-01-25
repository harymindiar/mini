package main

import (
	. "github.com/harymindiar/mini/core"
	. "github.com/harymindiar/mini/route"
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

	// register service
	registerService()

	// lock container, to make sure all service not changed when the application ran
	App.LockContainer()

	// add routes
	App.AddRoute(new(Index))

	// Start server
	App.Echo.Logger.Fatal(App.Echo.Start(":1323"))
}

func registerService() {

}
