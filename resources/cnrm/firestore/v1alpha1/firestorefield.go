package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirestoreField is the Schema for the FirestoreField API
type FirestoreField struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirestoreFieldSpec   `json:"spec,omitempty"`
	Status FirestoreFieldStatus `json:"status,omitempty"`
}

// FirestoreFieldSpec defines the desired state of FirestoreField.
type FirestoreFieldSpec struct {
	// The collectionGroup of which this field is a part.
	CollectionGroup string `json:"collectionGroup,omitempty" yaml:"collectionGroup,omitempty"`
	// The FirestoreDatabase containing the collection group for this field.
	DatabaseRef map[string]interface{} `json:"databaseRef,omitempty" yaml:"databaseRef,omitempty"`
	// The index configuration for this field. If unset, field indexing will revert to the configuration defined by the `ancestor_field`. To explicitly remove all indexes for this field, specify an index config with an empty list of indexes.
	IndexConfig map[string]interface{} `json:"indexConfig,omitempty" yaml:"indexConfig,omitempty"`
	// The FirestoreField name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The TTL configuration for this `Field`.
	TtlConfig map[string]interface{} `json:"ttlConfig,omitempty" yaml:"ttlConfig,omitempty"`
}

// FirestoreFieldStatus defines the observed state of FirestoreField.
type FirestoreFieldStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
