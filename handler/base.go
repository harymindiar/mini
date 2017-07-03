package handler

import (
	"github.com/harymindiar/mini/core"
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
