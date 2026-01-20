package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkServicesGateway represents a networkservices.cnrm.cloud.google.com NetworkServicesGateway resource.
type NetworkServicesGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesGatewaySpec   `json:"spec,omitempty"`
	Status NetworkServicesGatewayStatus `json:"status,omitempty"`
}

// NetworkServicesGatewaySpec defines the desired state of NetworkServicesGateway.
type NetworkServicesGatewaySpec struct {
	// One or more addresses with ports in format of ":" that the Gateway must receive traffic on. The proxy binds to the ports specified. IP address can be anything that is allowed by the underlying infrastructure (auto-allocation, static IP, BYOIP).
	Addresses []string `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. One or more ports that the Gateway must receive traffic on. The proxy binds to the ports specified. Gateway listen on 0.0.0.0 on the ports specified below.
	Ports []int64 `json:"ports,omitempty" yaml:"ports,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Required. Immutable. Scope determines how configuration across multiple Gateway instances are merged. The configuration for multiple Gateway instances with the same scope will be merged as presented as a single coniguration to the proxy/load balancer. Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`
	ServerTlsPolicyRef map[string]interface{} `json:"serverTlsPolicyRef,omitempty" yaml:"serverTlsPolicyRef,omitempty"`
	// Immutable. Immutable. The type of the customer managed gateway. Possible values: TYPE_UNSPECIFIED, OPEN_MESH, SECURE_WEB_GATEWAY
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// NetworkServicesGatewayStatus defines the observed state of NetworkServicesGateway.
type NetworkServicesGatewayStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
