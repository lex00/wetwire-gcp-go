package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkServicesHTTPRoute represents a networkservices.cnrm.cloud.google.com NetworkServicesHTTPRoute resource.
type NetworkServicesHTTPRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesHTTPRouteSpec   `json:"spec,omitempty"`
	Status NetworkServicesHTTPRouteStatus `json:"status,omitempty"`
}

// NetworkServicesHTTPRouteSpec defines the desired state of NetworkServicesHTTPRoute.
type NetworkServicesHTTPRouteSpec struct {
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Gateways []map[string]interface{} `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	// Required. Hostnames define a set of hosts that should match against the HTTP host header to select a HttpRoute to process the request. Hostname is the fully qualified domain name of a network host, as defined by RFC 1123 with the exception that ip addresses are not allowed. Wildcard hosts are supported as "*" (no prefix or suffix allowed).
	Hostnames []string `json:"hostnames,omitempty" yaml:"hostnames,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	Meshes []map[string]interface{} `json:"meshes,omitempty" yaml:"meshes,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Rules that define how traffic is routed and handled.
	Rules []map[string]interface{} `json:"rules,omitempty" yaml:"rules,omitempty"`
}

// NetworkServicesHTTPRouteStatus defines the observed state of NetworkServicesHTTPRoute.
type NetworkServicesHTTPRouteStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
