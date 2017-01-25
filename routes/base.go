package routes

import (
	. "github.com/harymindiar/mini/core"
)

type Base struct {
	application ContainerInterface
}

func (b *Base) SetApplication(a ContainerInterface) {
	b.application = a
}

func (b *Base) GetApplication() ContainerInterface {
	return b.application
}