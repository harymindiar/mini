package app

import (
	"github.com/gorilla/mux"
	"github.com/harymindiar/mini/handler"
	"github.com/harymindiar/mini/provider"
)

// Router here you need to register all the route of application
func (a *Application) Router() *mux.Router {
	// get router service
	container, _ := a.Container.Get(provider.RouterServiceName)
	router := container.(*mux.Router)

	baseHandler := handler.NewHandler(a.Container)
	router.HandleFunc("/", baseHandler.Index)

	return router
}
