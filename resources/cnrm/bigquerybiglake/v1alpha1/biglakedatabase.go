package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigLakeDatabase is the Schema for the BigLakeDatabase API
type BigLakeDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigLakeDatabaseSpec   `json:"spec,omitempty"`
	Status BigLakeDatabaseStatus `json:"status,omitempty"`
}

// BigLakeDatabaseSpec defines the desired state of BigLakeDatabase.
type BigLakeDatabaseSpec struct {
	// Options of a Hive database.
	HiveOptions map[string]interface{} `json:"hiveOptions,omitempty" yaml:"hiveOptions,omitempty"`
	// Required. Defines the parent path of the resource.
	ParentCatalogRef map[string]interface{} `json:"parentCatalogRef,omitempty" yaml:"parentCatalogRef,omitempty"`
	// The BigLakeDatabase name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The database type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// BigLakeDatabaseStatus defines the observed state of BigLakeDatabase.
type BigLakeDatabaseStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
