package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceDirectoryNamespace represents a servicedirectory.cnrm.cloud.google.com ServiceDirectoryNamespace resource.
type ServiceDirectoryNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDirectoryNamespaceSpec   `json:"spec,omitempty"`
	Status ServiceDirectoryNamespaceStatus `json:"status,omitempty"`
}

// ServiceDirectoryNamespaceSpec defines the desired state of ServiceDirectoryNamespace.
type ServiceDirectoryNamespaceSpec struct {
	// Immutable. The location for the Namespace.
	// A full list of valid locations can be found by running
	// 'gcloud beta service-directory locations list'.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The namespaceId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ServiceDirectoryNamespaceStatus defines the observed state of ServiceDirectoryNamespace.
type ServiceDirectoryNamespaceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
