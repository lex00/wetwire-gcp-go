package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryDataset is the Schema for the BigQueryDataset API
type BigQueryDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryDatasetSpec   `json:"spec,omitempty"`
	Status BigQueryDatasetStatus `json:"status,omitempty"`
}

// BigQueryDatasetSpec defines the desired state of BigQueryDataset.
type BigQueryDatasetSpec struct {
	// An array of objects that define dataset access for one or more entities.
	Access []map[string]interface{} `json:"access,omitempty" yaml:"access,omitempty"`
	// Optional. Defines the default collation specification of future tables
	// created in the dataset. If a table is created in this dataset without
	// table-level default collation, then the table inherits the dataset default
	// collation, which is applied to the string fields that do not have explicit
	// collation specified. A change to this field affects only tables created
	// afterwards, and does not alter the existing tables.
	// The following values are supported:
	// * 'und:ci': undetermined locale, case-insensitive.
	// * '': empty string. Default to case-sensitive behavior.
	DefaultCollation string `json:"defaultCollation,omitempty" yaml:"defaultCollation,omitempty"`
	// The default encryption key for all tables in the dataset. After this property is set, the encryption key of all newly-created tables in the dataset is set to this value unless the table creation request or query explicitly overrides the key.
	DefaultEncryptionConfiguration map[string]interface{} `json:"defaultEncryptionConfiguration,omitempty" yaml:"defaultEncryptionConfiguration,omitempty"`
	// This default partition expiration, expressed in milliseconds.
	// When new time-partitioned tables are created in a dataset where this
	// property is set, the table will inherit this value, propagated as the
	// `TimePartitioning.expirationMs` property on the new table.  If you set
	// `TimePartitioning.expirationMs` explicitly when creating a table,
	// the `defaultPartitionExpirationMs` of the containing dataset is ignored.
	// When creating a partitioned table, if `defaultPartitionExpirationMs`
	// is set, the `defaultTableExpirationMs` value is ignored and the table
	// will not be inherit a table expiration deadline.
	DefaultPartitionExpirationMs int64 `json:"defaultPartitionExpirationMs,omitempty" yaml:"defaultPartitionExpirationMs,omitempty"`
	// Optional. The default lifetime of all tables in the dataset, in milliseconds. The minimum lifetime value is 3600000 milliseconds (one hour). To clear an existing default expiration with a PATCH request, set to 0. Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.
	DefaultTableExpirationMs int64 `json:"defaultTableExpirationMs,omitempty" yaml:"defaultTableExpirationMs,omitempty"`
	// Optional. A user-friendly description of the dataset.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. A descriptive name for the dataset.
	FriendlyName string `json:"friendlyName,omitempty" yaml:"friendlyName,omitempty"`
	// Optional. TRUE if the dataset and its table names are case-insensitive, otherwise FALSE. By default, this is FALSE, which means the dataset and its table names are case-sensitive. This field does not affect routine references.
	IsCaseInsensitive bool `json:"isCaseInsensitive,omitempty" yaml:"isCaseInsensitive,omitempty"`
	// Optional. The geographic location where the dataset should reside. See https://cloud.google.com/bigquery/docs/locations for supported locations.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days). The default value is 168 hours if this is not set.
	MaxTimeTravelHours string `json:"maxTimeTravelHours,omitempty" yaml:"maxTimeTravelHours,omitempty"`
	// Optional. The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The BigQueryDataset name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Updates storage_billing_model for the dataset.
	StorageBillingModel string `json:"storageBillingModel,omitempty" yaml:"storageBillingModel,omitempty"`
}

// BigQueryDatasetStatus defines the observed state of BigQueryDataset.
type BigQueryDatasetStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
