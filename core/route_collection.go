package core

import (
	"github.com/labstack/echo"
)

type RouteCollection interface {
	Add(e *echo.Echo) *echo.Echo
	SetApplication(a *Container)
	GetApplication() *Container
}
