package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VMwareEngineNetwork is the Schema for the VMwareEngineNetwork API
type VMwareEngineNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VMwareEngineNetworkSpec   `json:"spec,omitempty"`
	Status VMwareEngineNetworkStatus `json:"status,omitempty"`
}

// VMwareEngineNetworkSpec defines the desired state of VMwareEngineNetwork.
type VMwareEngineNetworkSpec struct {
	// User-provided description for this VMware Engine network.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The VMwareEngineNetwork name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. VMware Engine network type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// VMwareEngineNetworkStatus defines the observed state of VMwareEngineNetwork.
type VMwareEngineNetworkStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
