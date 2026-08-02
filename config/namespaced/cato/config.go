/*
Copyright 2021 Upbound Inc.
*/

// Package cato holds namespaced-scoped resource configuration overrides for the
// Cato provider. External names live in config/external_name.go and group/kind
// overrides in config/overrides.go, since those apply to both scopes. This
// package is the place for tweaks that only make sense for namespaced-scoped
// managed resources, such as cross-resource references.
package cato

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure applies the namespaced-scoped resource configurations.
func Configure(_ *ujconfig.Provider) {}
