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

// Router provider of routes
type Router struct {
	Base
}

// GetName get name of provider
func (r Router) GetName() string {
	return RouterProviderName
}

// Boot boot the provider
func (r *Router) Boot() error {

	r.Container.Set(RouterServiceName, mux.NewRouter())

	return nil
}
