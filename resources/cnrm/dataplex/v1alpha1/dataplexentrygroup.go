package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataplexEntryGroup is the Schema for the DataplexEntryGroup API
type DataplexEntryGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataplexEntryGroupSpec   `json:"spec,omitempty"`
	Status DataplexEntryGroupStatus `json:"status,omitempty"`
}

// DataplexEntryGroupSpec defines the desired state of DataplexEntryGroup.
type DataplexEntryGroupSpec struct {
	// Optional. Description of the EntryGroup.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. User friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The DataplexEntryGroup name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DataplexEntryGroupStatus defines the observed state of DataplexEntryGroup.
type DataplexEntryGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
