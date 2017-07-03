package handler

import (
	"github.com/harymindiar/mini/core"
	"github.com/harymindiar/mini/provider"
)

// Base handler
type Base struct {
	Container *core.Container
}

// NewHandler create handler
func NewHandler(c *core.Container) *Base {
	return &Base{
		Container: c,
	}
}

// Config get config
func (b *Base) Config() core.Config {
	config, _ := b.Container.Get(provider.ConfigServiceName)

	return config.(core.Config)
}
