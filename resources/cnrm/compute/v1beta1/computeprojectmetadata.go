package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeProjectMetadata represents a compute.cnrm.cloud.google.com ComputeProjectMetadata resource.
type ComputeProjectMetadata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeProjectMetadataSpec   `json:"spec,omitempty"`
	Status ComputeProjectMetadataStatus `json:"status,omitempty"`
}

// ComputeProjectMetadataSpec defines the desired state of ComputeProjectMetadata.
type ComputeProjectMetadataSpec struct {
	// A series of key value pairs.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// ComputeProjectMetadataStatus defines the observed state of ComputeProjectMetadata.
type ComputeProjectMetadataStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
