package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecureSourceManagerRepository is the Schema for the SecureSourceManagerRepository API
type SecureSourceManagerRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecureSourceManagerRepositorySpec   `json:"spec,omitempty"`
	Status SecureSourceManagerRepositoryStatus `json:"status,omitempty"`
}

// SecureSourceManagerRepositorySpec defines the desired state of SecureSourceManagerRepository.
type SecureSourceManagerRepositorySpec struct {
	// Input only. Initial configurations for the repository.
	InitialConfig map[string]interface{} `json:"initialConfig,omitempty" yaml:"initialConfig,omitempty"`
	// The name of the instance in which the repository is hosted, formatted as `projects/{project_number}/locations/{location_id}/instances/{instance_id}`
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// Immutable. Location of the instance.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The SecureSourceManagerRepository name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// SecureSourceManagerRepositoryStatus defines the observed state of SecureSourceManagerRepository.
type SecureSourceManagerRepositoryStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
