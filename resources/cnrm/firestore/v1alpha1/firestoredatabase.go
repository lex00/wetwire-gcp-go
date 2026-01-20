package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirestoreDatabase is the Schema for the FirestoreDatabase API
type FirestoreDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirestoreDatabaseSpec   `json:"spec,omitempty"`
	Status FirestoreDatabaseStatus `json:"status,omitempty"`
}

// FirestoreDatabaseSpec defines the desired state of FirestoreDatabase.
type FirestoreDatabaseSpec struct {
	// The concurrency control mode to use for this database. See https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases#concurrencymode for more info.
	ConcurrencyMode string `json:"concurrencyMode,omitempty" yaml:"concurrencyMode,omitempty"`
	// State of delete protection for the database.
	DeleteProtectionState string `json:"deleteProtectionState,omitempty" yaml:"deleteProtectionState,omitempty"`
	// The location of the database. Available locations are listed at https://cloud.google.com/firestore/docs/locations.
	LocationID string `json:"locationID,omitempty" yaml:"locationID,omitempty"`
	// Whether to enable the PITR feature on this database. See https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases#pointintimerecoveryenablement for more info.
	PointInTimeRecoveryEnablement string `json:"pointInTimeRecoveryEnablement,omitempty" yaml:"pointInTimeRecoveryEnablement,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The FirestoreDatabase name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// FirestoreDatabaseStatus defines the observed state of FirestoreDatabase.
type FirestoreDatabaseStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
