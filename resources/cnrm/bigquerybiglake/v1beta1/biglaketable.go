package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigLakeTable is the Schema for the BigLakeTable API
type BigLakeTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigLakeTableSpec   `json:"spec,omitempty"`
	Status BigLakeTableStatus `json:"status,omitempty"`
}

// BigLakeTableSpec defines the desired state of BigLakeTable.
type BigLakeTableSpec struct {
	// Options of a Hive table.
	HiveOptions map[string]interface{} `json:"hiveOptions,omitempty" yaml:"hiveOptions,omitempty"`
	// Required. The parent resource where this table will be created. Format: projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}
	ParentDatabaseRef map[string]interface{} `json:"parentDatabaseRef,omitempty" yaml:"parentDatabaseRef,omitempty"`
	// The BigLake Table ID. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The table type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// BigLakeTableStatus defines the observed state of BigLakeTable.
type BigLakeTableStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
