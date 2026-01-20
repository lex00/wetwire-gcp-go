package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeGlobalNetworkEndpoint represents a compute.cnrm.cloud.google.com ComputeGlobalNetworkEndpoint resource.
type ComputeGlobalNetworkEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeGlobalNetworkEndpointSpec   `json:"spec,omitempty"`
	Status ComputeGlobalNetworkEndpointStatus `json:"status,omitempty"`
}

// ComputeGlobalNetworkEndpointSpec defines the desired state of ComputeGlobalNetworkEndpoint.
type ComputeGlobalNetworkEndpointSpec struct {
	// Immutable. Fully qualified domain name of network endpoint.
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`
	// Immutable. The global network endpoint group this endpoint is part of.
	GlobalNetworkEndpointGroup string `json:"globalNetworkEndpointGroup,omitempty" yaml:"globalNetworkEndpointGroup,omitempty"`
	// Immutable. IPv4 address external endpoint.
	IPAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The port of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeGlobalNetworkEndpointStatus defines the observed state of ComputeGlobalNetworkEndpoint.
type ComputeGlobalNetworkEndpointStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
