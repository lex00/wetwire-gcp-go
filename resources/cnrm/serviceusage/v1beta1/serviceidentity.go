package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceIdentity is the Schema for the ServiceIdentity API
type ServiceIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceIdentitySpec   `json:"spec,omitempty"`
	Status ServiceIdentityStatus `json:"status,omitempty"`
}

// ServiceIdentitySpec defines the desired state of ServiceIdentity.
type ServiceIdentitySpec struct {
	// The project that this service identity belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The service name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ServiceIdentityStatus defines the observed state of ServiceIdentity.
type ServiceIdentityStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
