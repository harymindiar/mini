package app

import (
	"github.com/gorilla/mux"
	"github.com/harymindiar/mini/handler"
	"net/http"
)

// Handlers here you need to register all the route of application
func (a *Application) Handlers() http.Handler {
	router := mux.NewRouter()

	baseHandler := handler.NewHandler(a.Container)
	router.HandleFunc("/", baseHandler.Index)

	return router
}
