package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeRouterPeer represents a compute.cnrm.cloud.google.com ComputeRouterPeer resource.
type ComputeRouterPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouterPeerSpec   `json:"spec,omitempty"`
	Status ComputeRouterPeerStatus `json:"status,omitempty"`
}

// ComputeRouterPeerSpec defines the desired state of ComputeRouterPeer.
type ComputeRouterPeerSpec struct {
	// User-specified flag to indicate which mode to use for advertisement.
	// Valid values of this enum field are: 'DEFAULT', 'CUSTOM' Default value: "DEFAULT" Possible values: ["DEFAULT", "CUSTOM"].
	AdvertiseMode string `json:"advertiseMode,omitempty" yaml:"advertiseMode,omitempty"`
	// User-specified list of prefix groups to advertise in custom
	// mode, which currently supports the following option:
	// * 'ALL_SUBNETS': Advertises all of the router's own VPC subnets.
	// This excludes any routes learned for subnets that use VPC Network
	// Peering.
	// Note that this field can only be populated if advertiseMode is 'CUSTOM'
	// and overrides the list defined for the router (in the "bgp" message).
	// These groups are advertised in addition to any specified prefixes.
	// Leave this field blank to advertise no custom groups.
	AdvertisedGroups []string `json:"advertisedGroups,omitempty" yaml:"advertisedGroups,omitempty"`
	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is 'CUSTOM' and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	AdvertisedIPRanges []map[string]interface{} `json:"advertisedIpRanges,omitempty" yaml:"advertisedIpRanges,omitempty"`
	// The priority of routes advertised to this BGP peer.
	// Where there is more than one matching route of maximum
	// length, the routes with the lowest priority value win.
	AdvertisedRoutePriority int32 `json:"advertisedRoutePriority,omitempty" yaml:"advertisedRoutePriority,omitempty"`
	// BFD configuration for the BGP peering.
	Bfd map[string]interface{} `json:"bfd,omitempty" yaml:"bfd,omitempty"`
	// The status of the BGP peer connection. If set to false, any active session
	// with the peer is terminated and all associated routing information is removed.
	// If set to true, the peer connection can be established with routing information.
	// The default is true.
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`
	// Enable IPv6 traffic over BGP Peer. If not specified, it is disabled by default.
	EnableIpv6 bool `json:"enableIpv6,omitempty" yaml:"enableIpv6,omitempty"`
	// IP address of the interface inside Google Cloud Platform.
	// Only IPv4 is supported.
	IPAddress map[string]interface{} `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
	// IPv6 address of the interface inside Google Cloud Platform.
	// The address must be in the range 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64.
	// If you do not specify the next hop addresses, Google Cloud automatically
	// assigns unused addresses from the 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64 range for you.
	Ipv6NexthopAddress string `json:"ipv6NexthopAddress,omitempty" yaml:"ipv6NexthopAddress,omitempty"`
	// Peer BGP Autonomous System Number (ASN).
	// Each BGP interface may use a different value.
	PeerAsn int32 `json:"peerAsn,omitempty" yaml:"peerAsn,omitempty"`
	// IP address of the BGP interface outside Google Cloud Platform.
	// Only IPv4 is supported. Required if 'ip_address' is set.
	PeerIPAddress string `json:"peerIpAddress,omitempty" yaml:"peerIpAddress,omitempty"`
	// IPv6 address of the BGP interface outside Google Cloud Platform.
	// The address must be in the range 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64.
	// If you do not specify the next hop addresses, Google Cloud automatically
	// assigns unused addresses from the 2600:2d00:0:2::/64 or 2600:2d00:0:3::/64 range for you.
	PeerIpv6NexthopAddress string `json:"peerIpv6NexthopAddress,omitempty" yaml:"peerIpv6NexthopAddress,omitempty"`
	// Immutable. Region where the router and BgpPeer reside.
	// If it is not provided, the provider region is used.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The URI of the VM instance that is used as third-party router
	// appliances such as Next Gen Firewalls, Virtual Routers, or Router
	// Appliances. The VM instance must be located in zones contained in
	// the same region as this Cloud Router. The VM instance is the peer
	// side of the BGP session.
	RouterApplianceInstanceRef map[string]interface{} `json:"routerApplianceInstanceRef,omitempty" yaml:"routerApplianceInstanceRef,omitempty"`
	// The interface the BGP peer is associated with.
	RouterInterfaceRef map[string]interface{} `json:"routerInterfaceRef,omitempty" yaml:"routerInterfaceRef,omitempty"`
	// The Cloud Router in which this BGP peer will be configured.
	RouterRef map[string]interface{} `json:"routerRef,omitempty" yaml:"routerRef,omitempty"`
}

// ComputeRouterPeerStatus defines the observed state of ComputeRouterPeer.
type ComputeRouterPeerStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
