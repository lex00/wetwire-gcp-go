package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// APIKeysKey is the Schema for the APIKeys Key resource.
type APIKeysKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIKeysKeySpec   `json:"spec,omitempty"`
	Status APIKeysKeyStatus `json:"status,omitempty"`
}

// APIKeysKeySpec defines the desired state of APIKeysKey.
type APIKeysKeySpec struct {
	// Human-readable display name of this key that you can modify. The maximum length is 63 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Key restrictions.
	Restrictions map[string]interface{} `json:"restrictions,omitempty" yaml:"restrictions,omitempty"`
}

// APIKeysKeyStatus defines the observed state of APIKeysKey.
type APIKeysKeyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
