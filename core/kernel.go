package core

import (
	. "github.com/harymindiar/mini/routes"
	"github.com/labstack/echo"
)

type Kernel struct {
	Environment string
	Echo        *echo.Echo
}

func (k *Kernel) AddRoute(rc RouteCollection) {
	rc.Add(k.Echo)
}
