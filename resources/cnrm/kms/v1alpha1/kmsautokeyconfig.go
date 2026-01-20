package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KMSAutokeyConfig is the Schema for the KMSAutokeyConfig API
type KMSAutokeyConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSAutokeyConfigSpec   `json:"spec,omitempty"`
	Status KMSAutokeyConfigStatus `json:"status,omitempty"`
}

// KMSAutokeyConfigSpec defines the desired state of KMSAutokeyConfig.
type KMSAutokeyConfigSpec struct {
	// Immutable. The folder that this resource belongs to.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// The Project that this resource belongs to.
	KeyProject map[string]interface{} `json:"keyProject,omitempty" yaml:"keyProject,omitempty"`
}

// KMSAutokeyConfigStatus defines the observed state of KMSAutokeyConfig.
type KMSAutokeyConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
