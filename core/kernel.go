package core

import (
	"github.com/labstack/echo"
)

type Kernel struct {
	Environment string
	Echo        *echo.Echo
	Container   ContainerInterface
}

func (k *Kernel) AddRoute(rc RouteCollection) {
	rc.Add(k.Echo)
	rc.SetApplication(k.Container)
}
