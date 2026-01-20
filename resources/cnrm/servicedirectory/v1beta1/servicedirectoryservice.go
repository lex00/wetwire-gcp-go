package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceDirectoryService represents a servicedirectory.cnrm.cloud.google.com ServiceDirectoryService resource.
type ServiceDirectoryService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDirectoryServiceSpec   `json:"spec,omitempty"`
	Status ServiceDirectoryServiceStatus `json:"status,omitempty"`
}

// ServiceDirectoryServiceSpec defines the desired state of ServiceDirectoryService.
type ServiceDirectoryServiceSpec struct {
	// The ServiceDirectoryNamespace that this service belongs to.
	NamespaceRef map[string]interface{} `json:"namespaceRef,omitempty" yaml:"namespaceRef,omitempty"`
	// Immutable. Optional. The serviceId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ServiceDirectoryServiceStatus defines the observed state of ServiceDirectoryService.
type ServiceDirectoryServiceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
