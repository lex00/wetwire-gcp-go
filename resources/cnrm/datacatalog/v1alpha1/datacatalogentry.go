package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataCatalogEntry is the Schema for the DataCatalogEntry API
type DataCatalogEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCatalogEntrySpec   `json:"spec,omitempty"`
	Status DataCatalogEntryStatus `json:"status,omitempty"`
}

// DataCatalogEntrySpec defines the desired state of DataCatalogEntry.
type DataCatalogEntrySpec struct {
	// Business Context of the entry. Not supported for BigQuery datasets
	BusinessContext map[string]interface{} `json:"businessContext,omitempty" yaml:"businessContext,omitempty"`
	// Specification that applies to Cloud Bigtable system. Only settable when `integrated_system` is equal to `CLOUD_BIGTABLE`
	CloudBigtableSystemSpec map[string]interface{} `json:"cloudBigtableSystemSpec,omitempty" yaml:"cloudBigtableSystemSpec,omitempty"`
	// Specification that applies to a data source connection. Valid only for entries with the `DATA_SOURCE_CONNECTION` type.
	DataSourceConnectionSpec map[string]interface{} `json:"dataSourceConnectionSpec,omitempty" yaml:"dataSourceConnectionSpec,omitempty"`
	// Specification that applies to a table resource. Valid only for entries with the `TABLE` or `EXPLORE` type.
	DatabaseTableSpec map[string]interface{} `json:"databaseTableSpec,omitempty" yaml:"databaseTableSpec,omitempty"`
	// Specification that applies to a dataset.
	DatasetSpec map[string]interface{} `json:"datasetSpec,omitempty" yaml:"datasetSpec,omitempty"`
	// Entry description that can consist of several sentences or paragraphs
	// that describe entry contents.
	// The description must not contain Unicode non-characters as well as C0
	// and C1 control codes except tabs (HT), new lines (LF), carriage returns
	// (CR), and page breaks (FF).
	// The maximum size is 2000 bytes when encoded in UTF-8.
	// Default value is an empty string.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Display name of an entry.
	// The maximum size is 500 bytes when encoded in UTF-8.
	// Default value is an empty string.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Reference to the entry group that contains the entry.
	EntryGroupRef map[string]interface{} `json:"entryGroupRef,omitempty" yaml:"entryGroupRef,omitempty"`
	// FeatureonlineStore spec for Vertex AI Feature Store.
	FeatureOnlineStoreSpec map[string]interface{} `json:"featureOnlineStoreSpec,omitempty" yaml:"featureOnlineStoreSpec,omitempty"`
	// Specification that applies to a fileset resource. Valid only for entries with the `FILESET` type.
	FilesetSpec map[string]interface{} `json:"filesetSpec,omitempty" yaml:"filesetSpec,omitempty"`
	// Specification that applies to a Cloud Storage fileset. Valid only for entries with the `FILESET` type.
	GcsFilesetSpec map[string]interface{} `json:"gcsFilesetSpec,omitempty" yaml:"gcsFilesetSpec,omitempty"`
	// Cloud labels attached to the entry.
	// In Data Catalog, you can create and modify labels attached only to custom
	// entries. Synced entries have unmodifiable labels that come from the source
	// system.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Specification that applies to Looker sysstem. Only settable when `user_specified_system` is equal to `LOOKER`
	LookerSystemSpec map[string]interface{} `json:"lookerSystemSpec,omitempty" yaml:"lookerSystemSpec,omitempty"`
	// Model specification.
	ModelSpec map[string]interface{} `json:"modelSpec,omitempty" yaml:"modelSpec,omitempty"`
	// The DataCatalogEntry name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Specification that applies to a user-defined function or procedure. Valid only for entries with the `ROUTINE` type.
	RoutineSpec map[string]interface{} `json:"routineSpec,omitempty" yaml:"routineSpec,omitempty"`
	// Schema of the entry. An entry might not have any schema attached to it.
	Schema map[string]interface{} `json:"schema,omitempty" yaml:"schema,omitempty"`
	// Specification that applies to a Service resource.
	ServiceSpec map[string]interface{} `json:"serviceSpec,omitempty" yaml:"serviceSpec,omitempty"`
	// Timestamps from the underlying resource, not from the Data Catalog
	// entry.
	// Output only when the entry has a system listed in the `IntegratedSystem`
	// enum. For entries with `user_specified_system`, this field is optional
	// and defaults to an empty timestamp.
	SourceSystemTimestamps map[string]interface{} `json:"sourceSystemTimestamps,omitempty" yaml:"sourceSystemTimestamps,omitempty"`
	// Specification that applies to a relational database system. Only settable when `user_specified_system` is equal to `SQL_DATABASE`
	SqlDatabaseSystemSpec map[string]interface{} `json:"sqlDatabaseSystemSpec,omitempty" yaml:"sqlDatabaseSystemSpec,omitempty"`
	// The type of the entry.
	// For details, see [`EntryType`](#entrytype).
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// Resource usage statistics.
	UsageSignal map[string]interface{} `json:"usageSignal,omitempty" yaml:"usageSignal,omitempty"`
	// Indicates the entry's source system that Data Catalog doesn't
	// automatically integrate with.
	// The `user_specified_system` string has the following limitations:
	// * Is case insensitive.
	// * Must begin with a letter or underscore.
	// * Can only contain letters, numbers, and underscores.
	// * Must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem string `json:"userSpecifiedSystem,omitempty" yaml:"userSpecifiedSystem,omitempty"`
	// Custom entry type that doesn't match any of the values allowed for input
	// and listed in the `EntryType` enum.
	// When creating an entry, first check the type values in the enum.
	// If there are no appropriate types for the new entry,
	// provide a custom value, for example, `my_special_type`.
	// The `user_specified_type` string has the following limitations:
	// * Is case insensitive.
	// * Must begin with a letter or underscore.
	// * Can only contain letters, numbers, and underscores.
	// * Must be at least 1 character and at most 64 characters long.
	UserSpecifiedType string `json:"userSpecifiedType,omitempty" yaml:"userSpecifiedType,omitempty"`
}

// DataCatalogEntryStatus defines the observed state of DataCatalogEntry.
type DataCatalogEntryStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
