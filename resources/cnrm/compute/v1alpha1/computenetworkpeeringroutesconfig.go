package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNetworkPeeringRoutesConfig represents a compute.cnrm.cloud.google.com ComputeNetworkPeeringRoutesConfig resource.
type ComputeNetworkPeeringRoutesConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkPeeringRoutesConfigSpec   `json:"spec,omitempty"`
	Status ComputeNetworkPeeringRoutesConfigStatus `json:"status,omitempty"`
}

// ComputeNetworkPeeringRoutesConfigSpec defines the desired state of ComputeNetworkPeeringRoutesConfig.
type ComputeNetworkPeeringRoutesConfigSpec struct {
	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes bool `json:"exportCustomRoutes,omitempty" yaml:"exportCustomRoutes,omitempty"`
	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes bool `json:"importCustomRoutes,omitempty" yaml:"importCustomRoutes,omitempty"`
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The peering of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeNetworkPeeringRoutesConfigStatus defines the observed state of ComputeNetworkPeeringRoutesConfig.
type ComputeNetworkPeeringRoutesConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
