package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HealthcareDataset represents a healthcare.cnrm.cloud.google.com HealthcareDataset resource.
type HealthcareDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HealthcareDatasetSpec   `json:"spec,omitempty"`
	Status HealthcareDatasetStatus `json:"status,omitempty"`
}

// HealthcareDatasetSpec defines the desired state of HealthcareDataset.
type HealthcareDatasetSpec struct {
	// Immutable. The location for the Dataset.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`
}

// HealthcareDatasetStatus defines the observed state of HealthcareDataset.
type HealthcareDatasetStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
