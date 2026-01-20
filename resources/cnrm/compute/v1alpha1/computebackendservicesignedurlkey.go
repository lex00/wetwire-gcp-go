package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeBackendServiceSignedURLKey represents a compute.cnrm.cloud.google.com ComputeBackendServiceSignedURLKey resource.
type ComputeBackendServiceSignedURLKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeBackendServiceSignedURLKeySpec   `json:"spec,omitempty"`
	Status ComputeBackendServiceSignedURLKeyStatus `json:"status,omitempty"`
}

// ComputeBackendServiceSignedURLKeySpec defines the desired state of ComputeBackendServiceSignedURLKey.
type ComputeBackendServiceSignedURLKeySpec struct {
	BackendServiceRef map[string]interface{} `json:"backendServiceRef,omitempty" yaml:"backendServiceRef,omitempty"`
	// Immutable. 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	KeyValue map[string]interface{} `json:"keyValue,omitempty" yaml:"keyValue,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeBackendServiceSignedURLKeyStatus defines the observed state of ComputeBackendServiceSignedURLKey.
type ComputeBackendServiceSignedURLKeyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
