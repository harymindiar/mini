package app

import (
	"github.com/harymindiar/mini/provider"
)

// RegisterCoreProvider to register core services that used by application
func (a *Application) RegisterCoreProvider() {
	a.ProviderCollection.Add(&provider.Config{
		Config: a.Application.Config,
	})
	// register database provider when config provider registered
	a.ProviderCollection.Add(&provider.Database{})
}

// RegisterProvider to register provider
func (a *Application) RegisterProvider() {

}
