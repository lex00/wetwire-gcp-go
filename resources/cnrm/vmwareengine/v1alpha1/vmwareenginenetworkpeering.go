package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VMwareEngineNetworkPeering is the Schema for the VMwareEngineNetworkPeering API
type VMwareEngineNetworkPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VMwareEngineNetworkPeeringSpec   `json:"spec,omitempty"`
	Status VMwareEngineNetworkPeeringStatus `json:"status,omitempty"`
}

// VMwareEngineNetworkPeeringSpec defines the desired state of VMwareEngineNetworkPeering.
type VMwareEngineNetworkPeeringSpec struct {
	// Optional. User-provided description for this network peering.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. True if full mesh connectivity is created and managed automatically between peered networks; false otherwise. Currently this field is always true because Google Compute Engine automatically creates and manages subnetwork routes between two VPC networks when peering state is 'ACTIVE'.
	ExchangeSubnetRoutes bool `json:"exchangeSubnetRoutes,omitempty" yaml:"exchangeSubnetRoutes,omitempty"`
	// Optional. True if custom routes are exported to the peered network; false otherwise. The default value is true.
	ExportCustomRoutes bool `json:"exportCustomRoutes,omitempty" yaml:"exportCustomRoutes,omitempty"`
	// Optional. True if all subnet routes with a public IP address range are exported; false otherwise. The default value is true. IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportCustomRoutesWithPublicIP bool `json:"exportCustomRoutesWithPublicIP,omitempty" yaml:"exportCustomRoutesWithPublicIP,omitempty"`
	// Optional. True if custom routes are imported from the peered network; false otherwise. The default value is true.
	ImportCustomRoutes bool `json:"importCustomRoutes,omitempty" yaml:"importCustomRoutes,omitempty"`
	// Optional. True if all subnet routes with public IP address range are imported; false otherwise. The default value is true. IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported to peers and are not controlled by this field.
	ImportCustomRoutesWithPublicIP bool `json:"importCustomRoutesWithPublicIP,omitempty" yaml:"importCustomRoutesWithPublicIP,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Maximum transmission unit (MTU) in bytes. The default value is `1500`. If a value of `0` is provided for this field, VMware Engine uses the default value instead.
	PeerMTU int32 `json:"peerMTU,omitempty" yaml:"peerMTU,omitempty"`
	// Required. The name of the network to peer with a standard VMware Engine network. The provided network can be a consumer VPC network or another standard VMware Engine network.
	PeerNetwork map[string]interface{} `json:"peerNetwork,omitempty" yaml:"peerNetwork,omitempty"`
	// Required. The type of the network to peer with the VMware Engine network.
	PeerNetworkType string `json:"peerNetworkType,omitempty" yaml:"peerNetworkType,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The VMwareEngineNetworkPeering name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The relative resource name of the VMware Engine network.
	VmwareEngineNetworkRef map[string]interface{} `json:"vmwareEngineNetworkRef,omitempty" yaml:"vmwareEngineNetworkRef,omitempty"`
}

// VMwareEngineNetworkPeeringStatus defines the observed state of VMwareEngineNetworkPeering.
type VMwareEngineNetworkPeeringStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
