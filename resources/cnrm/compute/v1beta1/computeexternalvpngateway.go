package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeExternalVPNGateway represents a compute.cnrm.cloud.google.com ComputeExternalVPNGateway resource.
type ComputeExternalVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeExternalVPNGatewaySpec   `json:"spec,omitempty"`
	Status ComputeExternalVPNGatewayStatus `json:"status,omitempty"`
}

// ComputeExternalVPNGatewaySpec defines the desired state of ComputeExternalVPNGateway.
type ComputeExternalVPNGatewaySpec struct {
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. A list of interfaces on this external VPN gateway.
	Interface []map[string]interface{} `json:"interface,omitempty" yaml:"interface,omitempty"`
	// Immutable. Indicates the redundancy type of this external VPN gateway Possible values: ["FOUR_IPS_REDUNDANCY", "SINGLE_IP_INTERNALLY_REDUNDANT", "TWO_IPS_REDUNDANCY"].
	RedundancyType string `json:"redundancyType,omitempty" yaml:"redundancyType,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeExternalVPNGatewayStatus defines the observed state of ComputeExternalVPNGateway.
type ComputeExternalVPNGatewayStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
