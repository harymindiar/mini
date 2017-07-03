package core

import (
	"fmt"
)

// Provider interface
type Provider interface {
	SetContainer(container *Container)
	GetName() string
	Boot() error
}

// ProviderCollection provider collection
type ProviderCollection struct {
	provider  map[string]bool
	Container *Container
}

// NewProviderCollection create provider collection
func NewProviderCollection(c *Container) *ProviderCollection {
	return &ProviderCollection{
		provider:  make(map[string]bool, 0),
		Container: c,
	}
}

// Add provider to collection
func (pc *ProviderCollection) Add(p Provider) error {
	if _, ok := pc.provider[p.GetName()]; ok {
		return fmt.Errorf("Provider %s already exist", p.GetName())
	}

	p.SetContainer(pc.Container)
	p.Boot()
	pc.provider[p.GetName()] = true

	return nil
}
