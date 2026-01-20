package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigtableLogicalView is the Schema for the BigtableLogicalView API
type BigtableLogicalView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableLogicalViewSpec   `json:"spec,omitempty"`
	Status BigtableLogicalViewStatus `json:"status,omitempty"`
}

// BigtableLogicalViewSpec defines the desired state of BigtableLogicalView.
type BigtableLogicalViewSpec struct {
	// Optional. Set to true to make the LogicalView protected against deletion.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`
	// InstanceRef defines the resource reference to BigtableInstance, which "External" field holds the GCP identifier for the KRM object.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// The BigtableLogicalView's select query.
	Query string `json:"query,omitempty" yaml:"query,omitempty"`
	// The BigtableLogicalView ID. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BigtableLogicalViewStatus defines the observed state of BigtableLogicalView.
type BigtableLogicalViewStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
