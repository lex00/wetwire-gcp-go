package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryTable is the Schema for the BigQueryTable API
type BigQueryTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryTableSpec   `json:"spec,omitempty"`
	Status BigQueryTableStatus `json:"status,omitempty"`
}

// BigQueryTableSpec defines the desired state of BigQueryTable.
type BigQueryTableSpec struct {
	// Clustering specification for the table. Must be specified with time-based partitioning, data in the table will be first partitioned and subsequently clustered.
	Clustering []string `json:"clustering,omitempty" yaml:"clustering,omitempty"`
	// DatasetRef defines the resource reference to BigQueryDataset, which "External" field holds the GCP identifier for the KRM object.
	DatasetRef map[string]interface{} `json:"datasetRef,omitempty" yaml:"datasetRef,omitempty"`
	// A user-friendly description of this table.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Custom encryption configuration (e.g., Cloud KMS keys).
	EncryptionConfiguration map[string]interface{} `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`
	// The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.  The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
	ExpirationTime int64 `json:"expirationTime,omitempty" yaml:"expirationTime,omitempty"`
	// Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
	ExternalDataConfiguration map[string]interface{} `json:"externalDataConfiguration,omitempty" yaml:"externalDataConfiguration,omitempty"`
	// A descriptive name for this table.
	FriendlyName string `json:"friendlyName,omitempty" yaml:"friendlyName,omitempty"`
	// When using `alpha.cnrm.cloud.google.com/reconciler:direct` annotion, use labels field to set the labels for this resource on GCP. Otherwise, use .metadata.labels. Please refer to https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/4274 for context. The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// The materialized view definition.
	MaterializedView map[string]interface{} `json:"materializedView,omitempty" yaml:"materializedView,omitempty"`
	// The maximum staleness of data that could be returned when the table (or stale MV) is queried. Staleness encoded as a string encoding of sql IntervalValue type.
	MaxStaleness string `json:"maxStaleness,omitempty" yaml:"maxStaleness,omitempty"`
	// If specified, configures range partitioning for this table.
	RangePartitioning map[string]interface{} `json:"rangePartitioning,omitempty" yaml:"rangePartitioning,omitempty"`
	// If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
	RequirePartitionFilter bool `json:"requirePartitionFilter,omitempty" yaml:"requirePartitionFilter,omitempty"`
	// The BigQueryTable name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Describes the schema of this table.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
	// Tables Primary Key and Foreign Key information
	TableConstraints map[string]interface{} `json:"tableConstraints,omitempty" yaml:"tableConstraints,omitempty"`
	// If specified, configures time-based partitioning for this table.
	TimePartitioning map[string]interface{} `json:"timePartitioning,omitempty" yaml:"timePartitioning,omitempty"`
	// The view definition.
	View map[string]interface{} `json:"view,omitempty" yaml:"view,omitempty"`
}

// BigQueryTableStatus defines the observed state of BigQueryTable.
type BigQueryTableStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
