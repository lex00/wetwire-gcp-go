package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeVPNTunnel represents a compute.cnrm.cloud.google.com ComputeVPNTunnel resource.
type ComputeVPNTunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeVPNTunnelSpec   `json:"spec,omitempty"`
	Status ComputeVPNTunnelStatus `json:"status,omitempty"`
}

// ComputeVPNTunnelSpec defines the desired state of ComputeVPNTunnel.
type ComputeVPNTunnelSpec struct {
	// Immutable. An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. IKE protocol version to use when establishing the VPN tunnel with
	// peer VPN gateway.
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	IkeVersion int32 `json:"ikeVersion,omitempty" yaml:"ikeVersion,omitempty"`
	// Immutable. Local traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example '192.168.0.0/16'. The ranges should be disjoint.
	// Only IPv4 is supported.
	LocalTrafficSelector []string `json:"localTrafficSelector,omitempty" yaml:"localTrafficSelector,omitempty"`
	// Immutable. The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface int32 `json:"peerExternalGatewayInterface,omitempty" yaml:"peerExternalGatewayInterface,omitempty"`
	// The peer side external VPN gateway to which this VPN tunnel
	// is connected.
	PeerExternalGatewayRef map[string]interface{} `json:"peerExternalGatewayRef,omitempty" yaml:"peerExternalGatewayRef,omitempty"`
	// The peer side HA GCP VPN gateway to which this VPN tunnel is
	// connected. If provided, the VPN tunnel will automatically use the
	// same VPN gateway interface ID in the peer GCP VPN gateway.
	PeerGCPGatewayRef map[string]interface{} `json:"peerGCPGatewayRef,omitempty" yaml:"peerGCPGatewayRef,omitempty"`
	// Immutable. IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIP string `json:"peerIp,omitempty" yaml:"peerIp,omitempty"`
	// Immutable. The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Remote traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example '192.168.0.0/16'. The ranges should be disjoint.
	// Only IPv4 is supported.
	RemoteTrafficSelector []string `json:"remoteTrafficSelector,omitempty" yaml:"remoteTrafficSelector,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The router to be used for dynamic routing.
	RouterRef map[string]interface{} `json:"routerRef,omitempty" yaml:"routerRef,omitempty"`
	// Immutable. Shared secret used to set the secure session between the Cloud VPN
	// gateway and the peer VPN gateway.
	SharedSecret map[string]interface{} `json:"sharedSecret,omitempty" yaml:"sharedSecret,omitempty"`
	// The ComputeTargetVPNGateway with which this VPN tunnel is
	// associated.
	TargetVPNGatewayRef map[string]interface{} `json:"targetVPNGatewayRef,omitempty" yaml:"targetVPNGatewayRef,omitempty"`
	// Immutable. The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface int32 `json:"vpnGatewayInterface,omitempty" yaml:"vpnGatewayInterface,omitempty"`
	// The ComputeVPNGateway with which this VPN tunnel is associated.
	// This must be used if a High Availability VPN gateway resource is
	// created.
	VpnGatewayRef map[string]interface{} `json:"vpnGatewayRef,omitempty" yaml:"vpnGatewayRef,omitempty"`
}

// ComputeVPNTunnelStatus defines the observed state of ComputeVPNTunnel.
type ComputeVPNTunnelStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
