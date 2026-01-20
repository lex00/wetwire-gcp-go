package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNetwork represents a compute.cnrm.cloud.google.com ComputeNetwork resource.
type ComputeNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkSpec   `json:"spec,omitempty"`
	Status ComputeNetworkStatus `json:"status,omitempty"`
}

// ComputeNetworkSpec defines the desired state of ComputeNetwork.
type ComputeNetworkSpec struct {
	// Immutable. When set to 'true', the network is created in "auto subnet mode" and
	// it will create a subnet for each region automatically across the
	// '10.128.0.0/9' address range.
	// When set to 'false', the network is created in "custom subnet mode" so
	// the user can explicitly connect subnetwork resources.
	AutoCreateSubnetworks bool `json:"autoCreateSubnetworks,omitempty" yaml:"autoCreateSubnetworks,omitempty"`
	// If set to 'true', default routes ('0.0.0.0/0') will be deleted
	// immediately after network creation. Defaults to 'false'.
	DeleteDefaultRoutesOnCreate bool `json:"deleteDefaultRoutesOnCreate,omitempty" yaml:"deleteDefaultRoutesOnCreate,omitempty"`
	// Immutable. An optional description of this resource. The resource must be
	// recreated to modify this field.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Enable ULA internal ipv6 on this network. Enabling this feature will assign
	// a /48 from google defined ULA prefix fd20::/20.
	EnableUlaInternalIpv6 bool `json:"enableUlaInternalIpv6,omitempty" yaml:"enableUlaInternalIpv6,omitempty"`
	// Immutable. When enabling ula internal ipv6, caller optionally can specify the /48 range
	// they want from the google defined ULA prefix fd20::/20. The input must be a
	// valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will
	// fail if the speficied /48 is already in used by another resource.
	// If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field.
	InternalIpv6Range string `json:"internalIpv6Range,omitempty" yaml:"internalIpv6Range,omitempty"`
	// Immutable. Maximum Transmission Unit in bytes. The default value is 1460 bytes.
	// The minimum value for this field is 1300 and the maximum value is 8896 bytes (jumbo frames).
	// Note that packets larger than 1500 bytes (standard Ethernet) can be subject to TCP-MSS clamping or dropped
	// with an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or other VPCs
	// with varying MTUs.
	Mtu int32 `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	// Set the order that Firewall Rules and Firewall Policies are evaluated. Default value: "AFTER_CLASSIC_FIREWALL" Possible values: ["BEFORE_CLASSIC_FIREWALL", "AFTER_CLASSIC_FIREWALL"].
	NetworkFirewallPolicyEnforcementOrder string `json:"networkFirewallPolicyEnforcementOrder,omitempty" yaml:"networkFirewallPolicyEnforcementOrder,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The network-wide routing mode to use. If set to 'REGIONAL', this
	// network's cloud routers will only advertise routes with subnetworks
	// of this network in the same region as the router. If set to 'GLOBAL',
	// this network's cloud routers will advertise routes with all
	// subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"].
	RoutingMode string `json:"routingMode,omitempty" yaml:"routingMode,omitempty"`
}

// ComputeNetworkStatus defines the observed state of ComputeNetwork.
type ComputeNetworkStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
