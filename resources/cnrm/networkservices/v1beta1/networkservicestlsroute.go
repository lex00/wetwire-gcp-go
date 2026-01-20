package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkServicesTLSRoute represents a networkservices.cnrm.cloud.google.com NetworkServicesTLSRoute resource.
type NetworkServicesTLSRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesTLSRouteSpec   `json:"spec,omitempty"`
	Status NetworkServicesTLSRouteStatus `json:"status,omitempty"`
}

// NetworkServicesTLSRouteSpec defines the desired state of NetworkServicesTLSRoute.
type NetworkServicesTLSRouteSpec struct {
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Gateways []map[string]interface{} `json:"gateways,omitempty" yaml:"gateways,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	Meshes []map[string]interface{} `json:"meshes,omitempty" yaml:"meshes,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Rules that define how traffic is routed and handled. At least one RouteRule must be supplied. If there are multiple rules then the action taken will be the first rule to match.
	Rules []map[string]interface{} `json:"rules,omitempty" yaml:"rules,omitempty"`
}

// NetworkServicesTLSRouteStatus defines the observed state of NetworkServicesTLSRoute.
type NetworkServicesTLSRouteStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
