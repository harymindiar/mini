package app

import (
	"github.com/harymindiar/mini/handler"
)

// RegisterHandler here you need to register all the route of application
func (a *Application) RegisterHandler() {
	baseHandler := handler.NewHandler(a.Container)
	a.Application.AddHandlerFunc("/", baseHandler.Index)
}
