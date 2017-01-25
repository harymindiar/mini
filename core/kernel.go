package core

import (
	"github.com/labstack/echo"
)

type Kernel struct {
	Environment string
	Echo        *echo.Echo
	Container   *Container
}

func (k *Kernel) AddRoute(rc RouteCollection) {
	rc.Add(k.Echo)
	rc.SetApplication(k.Container)
}

func (k *Kernel) RegisterService(sp ServiceProvider) {
	k.Container.Set(sp.GetName(), sp.Register(k.Container))
}

func (k *Kernel) LockContainer() {
	k.Container.SetState(STATE_FREEZE)
}