package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageManagedFolder is the Schema for the StorageManagedFolder API
type StorageManagedFolder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageManagedFolderSpec   `json:"spec,omitempty"`
	Status StorageManagedFolderStatus `json:"status,omitempty"`
}

// StorageManagedFolderSpec defines the desired state of StorageManagedFolder.
type StorageManagedFolderSpec struct {
	// Required. The host project of the application.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The StorageManagedFolder name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The storage bucket where the folder will be created in.
	StoragebucketRef map[string]interface{} `json:"storagebucketRef,omitempty" yaml:"storagebucketRef,omitempty"`
}

// StorageManagedFolderStatus defines the observed state of StorageManagedFolder.
type StorageManagedFolderStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
