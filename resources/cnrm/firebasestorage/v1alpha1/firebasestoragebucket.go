package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirebaseStorageBucket represents a firebasestorage.cnrm.cloud.google.com FirebaseStorageBucket resource.
type FirebaseStorageBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirebaseStorageBucketSpec   `json:"spec,omitempty"`
	Status FirebaseStorageBucketStatus `json:"status,omitempty"`
}

// FirebaseStorageBucketSpec defines the desired state of FirebaseStorageBucket.
type FirebaseStorageBucketSpec struct {
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The bucketId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// FirebaseStorageBucketStatus defines the observed state of FirebaseStorageBucket.
type FirebaseStorageBucketStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
