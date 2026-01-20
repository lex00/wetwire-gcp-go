package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkConnectivitySpoke represents a networkconnectivity.cnrm.cloud.google.com NetworkConnectivitySpoke resource.
type NetworkConnectivitySpoke struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkConnectivitySpokeSpec   `json:"spec,omitempty"`
	Status NetworkConnectivitySpokeStatus `json:"status,omitempty"`
}

// NetworkConnectivitySpokeSpec defines the desired state of NetworkConnectivitySpoke.
type NetworkConnectivitySpokeSpec struct {
	// An optional description of the spoke.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable.
	HubRef map[string]interface{} `json:"hubRef,omitempty" yaml:"hubRef,omitempty"`
	// Immutable. A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.
	LinkedInterconnectAttachments map[string]interface{} `json:"linkedInterconnectAttachments,omitempty" yaml:"linkedInterconnectAttachments,omitempty"`
	// Immutable. The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances map[string]interface{} `json:"linkedRouterApplianceInstances,omitempty" yaml:"linkedRouterApplianceInstances,omitempty"`
	// Immutable. VPC network that is associated with the spoke.
	LinkedVPCNetwork map[string]interface{} `json:"linkedVPCNetwork,omitempty" yaml:"linkedVPCNetwork,omitempty"`
	// Immutable. The URIs of linked VPN tunnel resources
	LinkedVpnTunnels map[string]interface{} `json:"linkedVpnTunnels,omitempty" yaml:"linkedVpnTunnels,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// NetworkConnectivitySpokeStatus defines the observed state of NetworkConnectivitySpoke.
type NetworkConnectivitySpokeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
