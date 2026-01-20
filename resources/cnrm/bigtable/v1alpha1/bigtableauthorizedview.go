package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigtableAuthorizedView is the Schema for the BigtableAuthorizedView API
type BigtableAuthorizedView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableAuthorizedViewSpec   `json:"spec,omitempty"`
	Status BigtableAuthorizedViewStatus `json:"status,omitempty"`
}

// BigtableAuthorizedViewSpec defines the desired state of BigtableAuthorizedView.
type BigtableAuthorizedViewSpec struct {
	// Set to true to make the AuthorizedView protected against deletion. The parent Table and containing Instance cannot be deleted if an AuthorizedView has this bit set.
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`
	// The BigtableAuthorizedView name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// An AuthorizedView permitting access to an explicit subset of a Table.
	SubsetView map[string]interface{} `json:"subsetView,omitempty" yaml:"subsetView,omitempty"`
	// TableRef defines the resource reference to BigtableTable, which "External" field holds the GCP identifier for the KRM object.
	TableRef map[string]interface{} `json:"tableRef,omitempty" yaml:"tableRef,omitempty"`
}

// BigtableAuthorizedViewStatus defines the observed state of BigtableAuthorizedView.
type BigtableAuthorizedViewStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
