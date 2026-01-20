package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeVPNGateway represents a compute.cnrm.cloud.google.com ComputeVPNGateway resource.
type ComputeVPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeVPNGatewaySpec   `json:"spec,omitempty"`
	Status ComputeVPNGatewayStatus `json:"status,omitempty"`
}

// ComputeVPNGatewaySpec defines the desired state of ComputeVPNGateway.
type ComputeVPNGatewaySpec struct {
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The network this VPN gateway is accepting traffic for.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. The region this gateway should sit in.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The stack type for this VPN gateway to identify the IP protocols that are enabled.
	// If not specified, IPV4_ONLY will be used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6"].
	StackType string `json:"stackType,omitempty" yaml:"stackType,omitempty"`
	// Immutable. A list of interfaces on this VPN gateway.
	VpnInterfaces []map[string]interface{} `json:"vpnInterfaces,omitempty" yaml:"vpnInterfaces,omitempty"`
}

// ComputeVPNGatewayStatus defines the observed state of ComputeVPNGateway.
type ComputeVPNGatewayStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
