package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeGlobalNetworkEndpointGroup represents a compute.cnrm.cloud.google.com ComputeGlobalNetworkEndpointGroup resource.
type ComputeGlobalNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeGlobalNetworkEndpointGroupSpec   `json:"spec,omitempty"`
	Status ComputeGlobalNetworkEndpointGroupStatus `json:"status,omitempty"`
}

// ComputeGlobalNetworkEndpointGroupSpec defines the desired state of ComputeGlobalNetworkEndpointGroup.
type ComputeGlobalNetworkEndpointGroupSpec struct {
	// Immutable. The default port used if the port number is not specified in the
	// network endpoint.
	DefaultPort int32 `json:"defaultPort,omitempty" yaml:"defaultPort,omitempty"`
	// Immutable. An optional description of this resource. Provide this property when
	// you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Type of network endpoints in this network endpoint group. Possible values: ["INTERNET_IP_PORT", "INTERNET_FQDN_PORT"].
	NetworkEndpointType string `json:"networkEndpointType,omitempty" yaml:"networkEndpointType,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeGlobalNetworkEndpointGroupStatus defines the observed state of ComputeGlobalNetworkEndpointGroup.
type ComputeGlobalNetworkEndpointGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
