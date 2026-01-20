package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeRouterInterface represents a compute.cnrm.cloud.google.com ComputeRouterInterface resource.
type ComputeRouterInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouterInterfaceSpec   `json:"spec,omitempty"`
	Status ComputeRouterInterfaceStatus `json:"status,omitempty"`
}

// ComputeRouterInterfaceSpec defines the desired state of ComputeRouterInterface.
type ComputeRouterInterfaceSpec struct {
	InterconnectAttachmentRef map[string]interface{} `json:"interconnectAttachmentRef,omitempty" yaml:"interconnectAttachmentRef,omitempty"`
	// Immutable. The IP address and range of the interface. The IP range must be in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IPRange string `json:"ipRange,omitempty" yaml:"ipRange,omitempty"`
	PrivateIPAddressRef map[string]interface{} `json:"privateIpAddressRef,omitempty" yaml:"privateIpAddressRef,omitempty"`
	// The interface the BGP peer is associated with.
	RedundantInterfaceRef map[string]interface{} `json:"redundantInterfaceRef,omitempty" yaml:"redundantInterfaceRef,omitempty"`
	// Immutable. The region this interface's router sits in. If not specified, the project region will be used. Changing this forces a new interface to be created.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	RouterRef map[string]interface{} `json:"routerRef,omitempty" yaml:"routerRef,omitempty"`
	SubnetworkRef map[string]interface{} `json:"subnetworkRef,omitempty" yaml:"subnetworkRef,omitempty"`
	VpnTunnelRef map[string]interface{} `json:"vpnTunnelRef,omitempty" yaml:"vpnTunnelRef,omitempty"`
}

// ComputeRouterInterfaceStatus defines the observed state of ComputeRouterInterface.
type ComputeRouterInterfaceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
