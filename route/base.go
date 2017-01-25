package route

import (
	. "github.com/harymindiar/mini/core"
)

type Base struct {
	application *Container
}

func (b *Base) SetApplication(a *Container) {
	b.application = a
}

func (b *Base) GetApplication() *Container {
	return b.application
}
