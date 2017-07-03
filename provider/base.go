package provider

import (
	"github.com/harymindiar/mini/core"
)

// Base provider
type Base struct {
	Container *core.Container
}

// SetContainer set container
func (b *Base) SetContainer(c *core.Container) {
	b.Container = c
}
