package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeBackendBucketSignedURLKey represents a compute.cnrm.cloud.google.com ComputeBackendBucketSignedURLKey resource.
type ComputeBackendBucketSignedURLKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeBackendBucketSignedURLKeySpec   `json:"spec,omitempty"`
	Status ComputeBackendBucketSignedURLKeyStatus `json:"status,omitempty"`
}

// ComputeBackendBucketSignedURLKeySpec defines the desired state of ComputeBackendBucketSignedURLKey.
type ComputeBackendBucketSignedURLKeySpec struct {
	BackendBucketRef map[string]interface{} `json:"backendBucketRef,omitempty" yaml:"backendBucketRef,omitempty"`
	// Immutable. 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	KeyValue map[string]interface{} `json:"keyValue,omitempty" yaml:"keyValue,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeBackendBucketSignedURLKeyStatus defines the observed state of ComputeBackendBucketSignedURLKey.
type ComputeBackendBucketSignedURLKeyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
