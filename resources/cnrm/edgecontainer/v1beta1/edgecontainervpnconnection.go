package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EdgeContainerVpnConnection represents a edgecontainer.cnrm.cloud.google.com EdgeContainerVpnConnection resource.
type EdgeContainerVpnConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeContainerVpnConnectionSpec   `json:"spec,omitempty"`
	Status EdgeContainerVpnConnectionStatus `json:"status,omitempty"`
}

// EdgeContainerVpnConnectionSpec defines the desired state of EdgeContainerVpnConnection.
type EdgeContainerVpnConnectionSpec struct {
	ClusterRef map[string]interface{} `json:"clusterRef,omitempty" yaml:"clusterRef,omitempty"`
	// Immutable. Whether this VPN connection has HA enabled on cluster side. If enabled, when creating VPN connection we will attempt to use 2 ANG floating IPs.
	EnableHighAvailability bool `json:"enableHighAvailability,omitempty" yaml:"enableHighAvailability,omitempty"`
	// Immutable. Google Cloud Platform location.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. NAT gateway IP, or WAN IP address. If a customer has multiple NAT IPs, the customer needs to configure NAT such that only one external IP maps to the GMEC Anthos cluster.
	// This is empty if NAT is not used.
	NatGatewayIP string `json:"natGatewayIp,omitempty" yaml:"natGatewayIp,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The VPN connection Cloud Router name.
	Router string `json:"router,omitempty" yaml:"router,omitempty"`
	// Immutable. The network ID of VPC to connect to.
	Vpc string `json:"vpc,omitempty" yaml:"vpc,omitempty"`
	// Project detail of the VPC network. Required if VPC is in a different project than the cluster project.
	VpcProject map[string]interface{} `json:"vpcProject,omitempty" yaml:"vpcProject,omitempty"`
}

// EdgeContainerVpnConnectionStatus defines the observed state of EdgeContainerVpnConnection.
type EdgeContainerVpnConnectionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
