package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GKEHubFeature represents a gkehub.cnrm.cloud.google.com GKEHubFeature resource.
type GKEHubFeature struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKEHubFeatureSpec   `json:"spec,omitempty"`
	Status GKEHubFeatureStatus `json:"status,omitempty"`
}

// GKEHubFeatureSpec defines the desired state of GKEHubFeature.
type GKEHubFeatureSpec struct {
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	Spec map[string]interface{} `json:"spec,omitempty" yaml:"spec,omitempty"`
}

// GKEHubFeatureStatus defines the observed state of GKEHubFeature.
type GKEHubFeatureStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
