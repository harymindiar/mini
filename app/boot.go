package app

import (
	"github.com/gorilla/handlers"
	"github.com/harymindiar/mini/core"
	"net/http"
)

// ConfigPath path of configuration file
const ConfigPath = "."

// Application client
type Application struct {
	*core.Application
}

// NewApp create application
func NewApp() *Application {
	return &Application{
		Application: core.NewApp(ConfigPath),
	}
}

// Run the application
func (a *Application) Run() {
	a.RegisterCoreProvider()
	a.RegisterProvider()
	a.RegisterHandler()
	http.ListenAndServe(":"+a.Port, handlers.CORS()(a.Router))
}
