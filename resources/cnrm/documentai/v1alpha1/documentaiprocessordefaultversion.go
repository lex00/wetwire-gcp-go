package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DocumentAIProcessorDefaultVersion represents a documentai.cnrm.cloud.google.com DocumentAIProcessorDefaultVersion resource.
type DocumentAIProcessorDefaultVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DocumentAIProcessorDefaultVersionSpec   `json:"spec,omitempty"`
	Status DocumentAIProcessorDefaultVersionStatus `json:"status,omitempty"`
}

// DocumentAIProcessorDefaultVersionSpec defines the desired state of DocumentAIProcessorDefaultVersion.
type DocumentAIProcessorDefaultVersionSpec struct {
	// Immutable. Optional. The processor of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The version to set. Using 'stable' or 'rc' will cause the API to return the latest version in that release channel.
	// Apply 'lifecycle.ignore_changes' to the 'version' field to suppress this diff.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

// DocumentAIProcessorDefaultVersionStatus defines the observed state of DocumentAIProcessorDefaultVersion.
type DocumentAIProcessorDefaultVersionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
