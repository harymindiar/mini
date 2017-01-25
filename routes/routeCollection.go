package routes

import (
	"github.com/labstack/echo"
)

type RouteCollection interface {
	Add(e *echo.Echo) *echo.Echo
}
