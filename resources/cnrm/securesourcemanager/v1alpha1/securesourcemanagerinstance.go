package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecureSourceManagerInstance is the Schema for the SecureSourceManagerInstance API
type SecureSourceManagerInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecureSourceManagerInstanceSpec   `json:"spec,omitempty"`
	Status SecureSourceManagerInstanceStatus `json:"status,omitempty"`
}

// SecureSourceManagerInstanceSpec defines the desired state of SecureSourceManagerInstance.
type SecureSourceManagerInstanceSpec struct {
	// Optional. Immutable. Customer-managed encryption key name.
	KmsKeyRef map[string]interface{} `json:"kmsKeyRef,omitempty" yaml:"kmsKeyRef,omitempty"`
	// Optional. Labels as key value pairs.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable. Location of the instance.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. PrivateConfig includes settings for private instance.
	PrivateConfig map[string]interface{} `json:"privateConfig,omitempty" yaml:"privateConfig,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SecureSourceManagerInstanceStatus defines the observed state of SecureSourceManagerInstance.
type SecureSourceManagerInstanceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
