// Package registry provides a central registry for GCP Config Connector resource types.
package registry

import (
	"fmt"
	"strings"
	"sync"
)

// TypeInfo describes a registered Config Connector resource type.
type TypeInfo struct {
	Package    string // Go package alias (e.g., "computev1beta1")
	Group      string // API group (e.g., "compute.cnrm.cloud.google.com")
	Version    string // API version (e.g., "v1beta1")
	Kind       string // Resource kind (e.g., "ComputeInstance")
	APIVersion string // Full apiVersion (e.g., "compute.cnrm.cloud.google.com/v1beta1")
}

// PackageInfo holds information about a registered package.
type PackageInfo struct {
	Group      string
	Version    string
	APIVersion string
}

// Registry holds registered type information.
type Registry struct {
	types    map[string]TypeInfo
	packages map[string]PackageInfo
	kinds    map[string]TypeInfo
	mu       sync.RWMutex
}

// DefaultRegistry is the global registry instance.
var DefaultRegistry = NewRegistry()

// NewRegistry creates a new type registry.
func NewRegistry() *Registry {
	return &Registry{
		types:    make(map[string]TypeInfo),
		packages: make(map[string]PackageInfo),
		kinds:    make(map[string]TypeInfo),
	}
}

// Register adds a type to the registry.
func (r *Registry) Register(info TypeInfo) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := fmt.Sprintf("%s.%s", info.Package, info.Kind)
	r.types[key] = info

	r.packages[info.Package] = PackageInfo{
		Group:      info.Group,
		Version:    info.Version,
		APIVersion: info.APIVersion,
	}

	r.kinds[info.Kind] = info
}

// IsKnownType checks if a type name is a known Config Connector resource type.
func (r *Registry) IsKnownType(typeName string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if _, ok := r.types[typeName]; ok {
		return true
	}

	parts := strings.Split(typeName, ".")
	if len(parts) == 2 {
		pkg := parts[0]
		if _, ok := r.packages[pkg]; ok {
			return true
		}
	}

	if _, ok := r.kinds[typeName]; ok {
		return true
	}

	return false
}

// IsKnownPackage checks if a package name is registered.
func (r *Registry) IsKnownPackage(pkg string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, ok := r.packages[pkg]
	return ok
}

// GetTypeInfo returns type information for a type name.
func (r *Registry) GetTypeInfo(typeName string) (TypeInfo, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if info, ok := r.types[typeName]; ok {
		return info, true
	}

	if info, ok := r.kinds[typeName]; ok {
		return info, true
	}

	return TypeInfo{}, false
}

// APIVersionForPackage returns the API version string for a package.
func (r *Registry) APIVersionForPackage(pkg string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if info, ok := r.packages[pkg]; ok {
		return info.APIVersion
	}
	return ""
}

// ListPackages returns all registered package names.
func (r *Registry) ListPackages() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	pkgs := make([]string, 0, len(r.packages))
	for pkg := range r.packages {
		pkgs = append(pkgs, pkg)
	}
	return pkgs
}

// RegisterBulk registers multiple types at once.
func (r *Registry) RegisterBulk(types []TypeInfo) {
	for _, t := range types {
		r.Register(t)
	}
}
