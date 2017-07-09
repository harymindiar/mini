package provider

import (
	"github.com/gorilla/mux"
)

const (
	// RouterProviderName provider name of router
	RouterProviderName = "core.router_provider"

	// RouterServiceName service name of router
	RouterServiceName = "core.router"
)

// Route provider of routes
type Route struct {
	Base
}

// GetName get name of provider
func (r Route) GetName() string {
	return RouterProviderName
}

// Boot boot the provider
func (r *Route) Boot() error {

	r.Container.Set(RouterServiceName, mux.NewRouter())

	return nil
}
