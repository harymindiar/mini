package core

import (
	"github.com/labstack/echo"
	. "github.com/harymindiar/mini/routes"
)

type Kernel struct {
	Environment string
	Echo *echo.Echo
}

func (k *Kernel) AddRoute(rc RouteCollection) {
	rc.Add(k.Echo);
}
