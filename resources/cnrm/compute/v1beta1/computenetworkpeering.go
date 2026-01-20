package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNetworkPeering represents a compute.cnrm.cloud.google.com ComputeNetworkPeering resource.
type ComputeNetworkPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkPeeringSpec   `json:"spec,omitempty"`
	Status ComputeNetworkPeeringStatus `json:"status,omitempty"`
}

// ComputeNetworkPeeringSpec defines the desired state of ComputeNetworkPeering.
type ComputeNetworkPeeringSpec struct {
	// Whether to export the custom routes to the peer network. Defaults to false.
	ExportCustomRoutes bool `json:"exportCustomRoutes,omitempty" yaml:"exportCustomRoutes,omitempty"`
	// Immutable.
	ExportSubnetRoutesWithPublicIP bool `json:"exportSubnetRoutesWithPublicIp,omitempty" yaml:"exportSubnetRoutesWithPublicIp,omitempty"`
	// Whether to export the custom routes from the peer network. Defaults to false.
	ImportCustomRoutes bool `json:"importCustomRoutes,omitempty" yaml:"importCustomRoutes,omitempty"`
	// Immutable.
	ImportSubnetRoutesWithPublicIP bool `json:"importSubnetRoutesWithPublicIp,omitempty" yaml:"importSubnetRoutesWithPublicIp,omitempty"`
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	PeerNetworkRef map[string]interface{} `json:"peerNetworkRef,omitempty" yaml:"peerNetworkRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Which IP version(s) of traffic and routes are allowed to be imported or exported between peer networks. The default value is IPV4_ONLY. Possible values: ["IPV4_ONLY", "IPV4_IPV6"].
	StackType string `json:"stackType,omitempty" yaml:"stackType,omitempty"`
}

// ComputeNetworkPeeringStatus defines the observed state of ComputeNetworkPeering.
type ComputeNetworkPeeringStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
