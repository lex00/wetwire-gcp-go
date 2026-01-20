package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNetworkEndpointGroup represents a compute.cnrm.cloud.google.com ComputeNetworkEndpointGroup resource.
type ComputeNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkEndpointGroupSpec   `json:"spec,omitempty"`
	Status ComputeNetworkEndpointGroupStatus `json:"status,omitempty"`
}

// ComputeNetworkEndpointGroupSpec defines the desired state of ComputeNetworkEndpointGroup.
type ComputeNetworkEndpointGroupSpec struct {
	// Immutable. The default port used if the port number is not specified in the
	// network endpoint.
	DefaultPort int32 `json:"defaultPort,omitempty" yaml:"defaultPort,omitempty"`
	// Immutable. An optional description of this resource. Provide this property when
	// you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Location represents the geographical location of the ComputeNetworkEndpointGroup. Specify a zone name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Type of network endpoints in this network endpoint group.
	// NON_GCP_PRIVATE_IP_PORT is used for hybrid connectivity network
	// endpoint groups (see https://cloud.google.com/load-balancing/docs/hybrid).
	// Note that NON_GCP_PRIVATE_IP_PORT can only be used with Backend Services
	// that 1) have the following load balancing schemes: EXTERNAL, EXTERNAL_MANAGED,
	// INTERNAL_MANAGED, and INTERNAL_SELF_MANAGED and 2) support the RATE or
	// CONNECTION balancing modes.
	// Possible values include: GCE_VM_IP, GCE_VM_IP_PORT, and NON_GCP_PRIVATE_IP_PORT. Default value: "GCE_VM_IP_PORT" Possible values: ["GCE_VM_IP", "GCE_VM_IP_PORT", "NON_GCP_PRIVATE_IP_PORT"].
	NetworkEndpointType string `json:"networkEndpointType,omitempty" yaml:"networkEndpointType,omitempty"`
	// The network to which all network endpoints in the NEG belong. Uses
	// "default" project network if unspecified.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional subnetwork to which all network endpoints in the NEG belong.
	SubnetworkRef map[string]interface{} `json:"subnetworkRef,omitempty" yaml:"subnetworkRef,omitempty"`
}

// ComputeNetworkEndpointGroupStatus defines the observed state of ComputeNetworkEndpointGroup.
type ComputeNetworkEndpointGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
