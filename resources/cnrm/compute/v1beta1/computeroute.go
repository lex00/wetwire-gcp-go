package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeRoute represents a compute.cnrm.cloud.google.com ComputeRoute resource.
type ComputeRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouteSpec   `json:"spec,omitempty"`
	Status ComputeRouteStatus `json:"status,omitempty"`
}

// ComputeRouteSpec defines the desired state of ComputeRoute.
type ComputeRouteSpec struct {
	// Immutable. An optional description of this resource. Provide this property
	// when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The destination range of outgoing packets that this route applies to.
	// Only IPv4 is supported.
	DestRange string `json:"destRange,omitempty" yaml:"destRange,omitempty"`
	// The network that this route applies to.
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. URL to a gateway that should handle matching packets.
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway'
	// * 'projects/project/global/gateways/default-internet-gateway'
	// * 'global/gateways/default-internet-gateway'
	// * The string 'default-internet-gateway'.
	NextHopGateway string `json:"nextHopGateway,omitempty" yaml:"nextHopGateway,omitempty"`
	// A forwarding rule of type loadBalancingScheme=INTERNAL that should
	// handle matching packets.  Note that this can only be used when the
	// destinationRange is a public (non-RFC 1918) IP CIDR range.
	NextHopILBRef map[string]interface{} `json:"nextHopILBRef,omitempty" yaml:"nextHopILBRef,omitempty"`
	// Instance that should handle matching packets.
	NextHopInstanceRef map[string]interface{} `json:"nextHopInstanceRef,omitempty" yaml:"nextHopInstanceRef,omitempty"`
	// Immutable. Network IP address of an instance that should handle matching packets.
	NextHopIP string `json:"nextHopIp,omitempty" yaml:"nextHopIp,omitempty"`
	// The ComputeVPNTunnel that should handle matching packets
	NextHopVPNTunnelRef map[string]interface{} `json:"nextHopVPNTunnelRef,omitempty" yaml:"nextHopVPNTunnelRef,omitempty"`
	// Immutable. The priority of this route. Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	// Default value is 1000. Valid range is 0 through 65535.
	Priority int32 `json:"priority,omitempty" yaml:"priority,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. A list of instance tags to which this route applies.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// ComputeRouteStatus defines the observed state of ComputeRoute.
type ComputeRouteStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
