package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirebaseProject represents a firebase.cnrm.cloud.google.com FirebaseProject resource.
type FirebaseProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirebaseProjectSpec   `json:"spec,omitempty"`
	Status FirebaseProjectStatus `json:"status,omitempty"`
}

// FirebaseProjectSpec defines the desired state of FirebaseProject.
type FirebaseProjectSpec struct {
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The project of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// FirebaseProjectStatus defines the observed state of FirebaseProject.
type FirebaseProjectStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
