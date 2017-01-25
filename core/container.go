package core

import (
	"fmt"
	"sync"
)

const STATE_FREEZE = 1

const ERROR_MESSAGE_STATE_FROZEN = "Can't add or modified %s. All service have been frozen"

type ContainerInterface interface {
	Add(name string, object interface{})
	Remove(name string)
	Get(name string) (object interface{}, ok bool)
}

type Container struct {
	mux   sync.RWMutex
	m     map[string]interface{}
	state int
}

// Add a service
func (c *Container) Set(name string, object interface{}) {
	if c.state == STATE_FREEZE {
		panic(fmt.Sprintf(ERROR_MESSAGE_STATE_FROZEN, name))
	}

	if c.m == nil {
		c.m = make(map[string]interface{})
	}

	c.mux.Lock()
	c.m[name] = object
	c.mux.Unlock()
}

func (c *Container) Get(name string) (object interface{}, ok bool) {
	c.mux.RLock()
	object, ok = c.m[name]
	c.mux.RUnlock()

	return object, ok
}

func (c *Container) SetState(state int) {
	c.state = state
}

type ServiceProvider interface {
	GetName() string
	Register(app *Container) interface{}
}
