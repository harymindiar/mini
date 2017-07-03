package app

import (
	"github.com/harymindiar/mini/provider"
)

// RegisterCoreProvider to register core services that used by application
func (a *Application) RegisterCoreProvider() {
	a.ProviderCollection.Add(&provider.Config{
		Config: a.Application.Config,
	})
}

// RegisterProvider to register provider
func (a *Application) RegisterProvider() {

}
