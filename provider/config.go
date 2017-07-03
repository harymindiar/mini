package provider

import (
	"github.com/harymindiar/mini/core"
)

const (
	// ConfigProviderName name of config provider
	ConfigProviderName = "core.config_provider"

	// ConfigServiceName service name of config
	ConfigServiceName = "core.config"
)

// Config struct config provider
type Config struct {
	Base
	Config core.Config
}

// GetName get provider name
func (c Config) GetName() string {
	return ConfigProviderName
}

// Boot booting provider
func (c *Config) Boot() error {
	c.Container.Set(ConfigServiceName, c.Config)

	return nil
}
