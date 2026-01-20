package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EdgeNetworkSubnet represents a edgenetwork.cnrm.cloud.google.com EdgeNetworkSubnet resource.
type EdgeNetworkSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeNetworkSubnetSpec   `json:"spec,omitempty"`
	Status EdgeNetworkSubnetStatus `json:"status,omitempty"`
}

// EdgeNetworkSubnetSpec defines the desired state of EdgeNetworkSubnet.
type EdgeNetworkSubnetSpec struct {
	// Immutable. A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The ranges of ipv4 addresses that are owned by this subnetwork, in CIDR format.
	Ipv4Cidr []string `json:"ipv4Cidr,omitempty" yaml:"ipv4Cidr,omitempty"`
	// Immutable. The ranges of ipv6 addresses that are owned by this subnetwork, in CIDR format.
	Ipv6Cidr []string `json:"ipv6Cidr,omitempty" yaml:"ipv6Cidr,omitempty"`
	// Immutable. The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The subnetId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. VLAN ID for this subnetwork. If not specified, one is assigned automatically.
	VlanID int32 `json:"vlanId,omitempty" yaml:"vlanId,omitempty"`
	// Immutable. The name of the target Distributed Cloud Edge zone.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// EdgeNetworkSubnetStatus defines the observed state of EdgeNetworkSubnet.
type EdgeNetworkSubnetStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
