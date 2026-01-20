package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BatchTask is the Schema for the BatchTask API
type BatchTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BatchTaskSpec   `json:"spec,omitempty"`
	Status BatchTaskStatus `json:"status,omitempty"`
}

// BatchTaskSpec defines the desired state of BatchTask.
type BatchTaskSpec struct {
	// Immutable. The location where the alloydb cluster should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The BatchTask name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BatchTaskStatus defines the observed state of BatchTask.
type BatchTaskStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
