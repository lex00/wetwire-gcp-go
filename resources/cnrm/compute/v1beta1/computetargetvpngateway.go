package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeTargetVPNGateway represents a compute.cnrm.cloud.google.com ComputeTargetVPNGateway resource.
type ComputeTargetVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetVPNGatewaySpec   `json:"spec,omitempty"`
	Status ComputeTargetVPNGatewayStatus `json:"status,omitempty"`
}

// ComputeTargetVPNGatewaySpec defines the desired state of ComputeTargetVPNGateway.
type ComputeTargetVPNGatewaySpec struct {
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The network this VPN gateway is accepting traffic for.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. The region this gateway should sit in.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeTargetVPNGatewayStatus defines the observed state of ComputeTargetVPNGateway.
type ComputeTargetVPNGatewayStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
