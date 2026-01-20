package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNetworkEndpoint represents a compute.cnrm.cloud.google.com ComputeNetworkEndpoint resource.
type ComputeNetworkEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkEndpointSpec   `json:"spec,omitempty"`
	Status ComputeNetworkEndpointStatus `json:"status,omitempty"`
}

// ComputeNetworkEndpointSpec defines the desired state of ComputeNetworkEndpoint.
type ComputeNetworkEndpointSpec struct {
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// Immutable. IPv4 address of network endpoint. The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	IPAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
	NetworkEndpointGroupRef map[string]interface{} `json:"networkEndpointGroupRef,omitempty" yaml:"networkEndpointGroupRef,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The port of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Zone where the containing network endpoint group is located.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputeNetworkEndpointStatus defines the observed state of ComputeNetworkEndpoint.
type ComputeNetworkEndpointStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
