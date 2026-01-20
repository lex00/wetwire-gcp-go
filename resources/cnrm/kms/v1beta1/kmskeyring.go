package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KMSKeyRing represents a KMS KeyRing.
type KMSKeyRing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KMSKeyRingSpec   `json:"spec,omitempty"`
	Status KMSKeyRingStatus `json:"status,omitempty"`
}

// KMSKeyRingSpec defines the desired state of KMSKeyRing.
type KMSKeyRingSpec struct {
	// Immutable. The location for the KeyRing. A full list of valid locations can be found by running 'gcloud kms locations list'.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// KMSKeyRingStatus defines the observed state of KMSKeyRing.
type KMSKeyRingStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
