package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigtableTable is the Schema for the BigtableTable API
type BigtableTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableTableSpec   `json:"spec,omitempty"`
	Status BigtableTableStatus `json:"status,omitempty"`
}

// BigtableTableSpec defines the desired state of BigtableTable.
type BigtableTableSpec struct {
	// Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days..
	ChangeStreamRetention string `json:"changeStreamRetention,omitempty" yaml:"changeStreamRetention,omitempty"`
	// The names of the column families that should be created immediately upon table creation, specified by name. The values that may be set are specified here. At least one column family must be specified.
	ColumnFamily []map[string]interface{} `json:"columnFamily,omitempty" yaml:"columnFamily,omitempty"`
	// NOTE: DeletionProtection proto field is changed from string (1.38) to bool (1.40) in cloud.google.com/go/bigtable/admin/apiv2/adminpb
	// Set to true to make the table protected against data loss. i.e. deleting
	// the following resources through Admin APIs are prohibited:
	// * The table.
	// * The column families in the table.
	// * The instance containing the table.
	// Note one can still delete the data stored in the table through Data APIs.
	DeletionProtection string `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`
	// Immutable. The instance to create the table in.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// The BigtableTable name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A list of predefined keys to split the table on.
	SplitKeys []string `json:"splitKeys,omitempty" yaml:"splitKeys,omitempty"`
}

// BigtableTableStatus defines the observed state of BigtableTable.
type BigtableTableStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
