package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeAddress represents a compute.cnrm.cloud.google.com ComputeAddress resource.
type ComputeAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeAddressSpec   `json:"spec,omitempty"`
	Status ComputeAddressStatus `json:"status,omitempty"`
}

// ComputeAddressSpec defines the desired state of ComputeAddress.
type ComputeAddressSpec struct {
	// Immutable. The static external IP address represented by this resource.
	// The IP address must be inside the specified subnetwork,
	// if any. Set by the API if undefined.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`
	// Immutable. The type of address to reserve.
	// Note: if you set this argument's value as 'INTERNAL' you need to leave the 'network_tier' argument unset in that resource block. Default value: "EXTERNAL" Possible values: ["INTERNAL", "EXTERNAL"].
	AddressType string `json:"addressType,omitempty" yaml:"addressType,omitempty"`
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"].
	IPVersion string `json:"ipVersion,omitempty" yaml:"ipVersion,omitempty"`
	// Immutable. The endpoint type of this address, which should be VM or NETLB. This is
	// used for deciding which type of endpoint this address can be used after
	// the external IPv6 address reservation. Possible values: ["VM", "NETLB"].
	Ipv6EndpointType string `json:"ipv6EndpointType,omitempty" yaml:"ipv6EndpointType,omitempty"`
	// Location represents the geographical location of the ComputeAddress. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The network in which to reserve the address. If global, the address
	// must be within the RFC1918 IP space. The network cannot be deleted
	// if there are any reserved IP ranges referring to it. This field can
	// only be used with INTERNAL type with the VPC_PEERING and
	// IPSEC_INTERCONNECT purposes.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM.
	// This argument should not be used when configuring Internal addresses, because [network tier cannot be set for internal traffic; it's always Premium](https://cloud.google.com/network-tiers/docs/overview). Possible values: ["PREMIUM", "STANDARD"].
	NetworkTier string `json:"networkTier,omitempty" yaml:"networkTier,omitempty"`
	// Immutable. The prefix length if the resource represents an IP range.
	PrefixLength int32 `json:"prefixLength,omitempty" yaml:"prefixLength,omitempty"`
	// Immutable. The purpose of this resource, which can be one of the following values.
	// * GCE_ENDPOINT for addresses that are used by VM instances, alias IP
	// ranges, load balancers, and similar resources.
	// * SHARED_LOADBALANCER_VIP for an address that can be used by multiple
	// internal load balancers.
	// * VPC_PEERING for addresses that are reserved for VPC peer networks.
	// * IPSEC_INTERCONNECT for addresses created from a private IP range that
	// are reserved for a VLAN attachment in an HA VPN over Cloud Interconnect
	// configuration. These addresses are regional resources.
	// * PRIVATE_SERVICE_CONNECT for a private network address that is used to
	// configure Private Service Connect. Only global internal addresses can use
	// this purpose.
	// This should only be set when using an Internal address.
	Purpose string `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The subnetwork in which to reserve the address. If an IP address is
	// specified, it must be within the subnetwork's IP range.  This field
	// can only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER
	// purposes.
	SubnetworkRef map[string]interface{} `json:"subnetworkRef,omitempty" yaml:"subnetworkRef,omitempty"`
}

// ComputeAddressStatus defines the observed state of ComputeAddress.
type ComputeAddressStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
