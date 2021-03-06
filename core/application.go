package core

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
	"sync"
)

// Config to get configuration
type Config interface {
	GetString(string) string
	GetBool(string) bool
	SetConfigName(string)
	AddConfigPath(string)
	ReadInConfig() error
	GetStringMap(key string) map[string]interface{}
}

// Application core of application
type Application struct {
	Environment        string
	Port               string
	Debug              bool
	Config             Config
	Container          *Container
	ProviderCollection *ProviderCollection
}

// NewApp create kernel
func NewApp(configPath string) *Application {
	a := &Application{
		Config:    viper.New(),
		Container: NewContainer(),
	}
	a.ProviderCollection = NewProviderCollection(a.Container)

	a.Config.SetConfigName("config")
	a.Config.AddConfigPath(configPath)
	a.Config.ReadInConfig()
	a.Environment = a.Config.GetString("environment")
	a.Port = a.Config.GetString("port")
	a.Debug = a.Config.GetBool("debug")

	return a
}

// Serve serve the handler
func (a *Application) Serve(h http.Handler) {
	log.Printf("Start to serve on port :%s", a.Port)
	http.ListenAndServe(":"+a.Port, h)
}

// IsDevelopment get current enviroment is development or not
func (a Application) IsDevelopment() bool {
	if a.Environment == "dev" {
		return true
	}

	return false
}

// Container to hold services
type Container struct {
	mux sync.RWMutex
	m   map[string]interface{}
}

// NewContainer create container
func NewContainer() *Container {
	return &Container{
		m: make(map[string]interface{}),
	}
}

// Set to set a service
func (c *Container) Set(name string, object interface{}) {
	c.mux.Lock()
	c.m[name] = object
	c.mux.Unlock()
}

// Get to get service
func (c *Container) Get(name string) (object interface{}, ok bool) {
	c.mux.RLock()
	object, ok = c.m[name]
	c.mux.RUnlock()
	if !ok {
		return nil, false
	}

	return object, ok
}
