package app

import (
	"github.com/gorilla/mux"
	"github.com/harymindiar/mini/handler"
	"github.com/harymindiar/mini/middleware"
	"github.com/urfave/negroni"
	"net/http"
)

// Handlers here you need to register all the route of application
func (a *Application) Handlers() http.Handler {
	baseHandler := handler.NewHandler(a.Container)

	r := mux.NewRouter()
	r.HandleFunc("/", baseHandler.Index)

	// default middleware
	n := negroni.New(negroni.NewRecovery())
	if a.IsDevelopment() {
		n.Use(middleware.NewLogger())
	}

	n.UseHandler(r)

	return n
}
