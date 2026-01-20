package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigtableMaterializedView is the Schema for the BigtableMaterializedView API
type BigtableMaterializedView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableMaterializedViewSpec   `json:"spec,omitempty"`
	Status BigtableMaterializedViewStatus `json:"status,omitempty"`
}

// BigtableMaterializedViewSpec defines the desired state of BigtableMaterializedView.
type BigtableMaterializedViewSpec struct {
	// Optional. Set to true to make the MaterializedView protected against deletion.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`
	// InstanceRef defines the resource reference to BigtableInstance, which "External" field holds the GCP identifier for the KRM object.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// Immutable. MaterializedView's select query.
	Query string `json:"query,omitempty" yaml:"query,omitempty"`
	// The BigtableMaterializedView name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BigtableMaterializedViewStatus defines the observed state of BigtableMaterializedView.
type BigtableMaterializedViewStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
