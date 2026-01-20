package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageFolder is the Schema for the StorageFolder API
type StorageFolder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageFolderSpec   `json:"spec,omitempty"`
	Status StorageFolderStatus `json:"status,omitempty"`
}

// StorageFolderSpec defines the desired state of StorageFolder.
type StorageFolderSpec struct {
	// Required. The host project of the application.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The StorageFolder name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The storage bucket where the folder will be created in.
	StoragebucketRef map[string]interface{} `json:"storagebucketRef,omitempty" yaml:"storagebucketRef,omitempty"`
}

// StorageFolderStatus defines the observed state of StorageFolder.
type StorageFolderStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
