package metrics

import (
	"errors"
	"reflect"
)

// Module is a simple interface for module Metrics
type Module interface {
	// Return subsystem name for the given metrics handler
	Name() string
}

// MetricsProvider is a Metrics module registry provider
// Use it as a single store to collect all the metrics so
// you can provide them as handles to the systems that need them
type MetricsProvider struct {
	mm map[string]Module
}

// NewMetricsProvider returns a new instance of a MetricsProvider
// with all the supplied modules set.
// Panics if multiple modules with the same name
// but different types are given.
// Todo: Initial creation shouldnt allow same name at all
func NewMetricsProvider(modules ...Module) *MetricsProvider {
	mp := MetricsProvider{make(map[string]Module)}
	for _, mod := range modules {
		name := mod.Name()
		if name == "" {
			panic(errors.New("Provided metrics module has empty name"))
		}
		err := mp.Set(mod)
		if err != nil {
			panic(err)
		}
	}

	return &mp
}

// Set a module in the provider.
// Trying to set a module with the same name but different
// type will return an error
func (mp *MetricsProvider) Set(module Module) error {
	name := module.Name()
	origmod, exists := mp.mm[name]
	if exists {
		// type assert so we can't overwrite different types
		if reflect.TypeOf(origmod) != reflect.TypeOf(module) {
			return errors.New("Updating existing module with different type")
		}
	}
	mp.mm[name] = module
	return nil
}

// Get the module from the provider with the given name
// may return nil if no module is set for that name
func (mp *MetricsProvider) Get(name string) Module {
	mod, _ := mp.mm[name]
	return mod
}
