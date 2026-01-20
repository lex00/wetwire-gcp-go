package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpannerDatabase represents a spanner.cnrm.cloud.google.com SpannerDatabase resource.
type SpannerDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpannerDatabaseSpec   `json:"spec,omitempty"`
	Status SpannerDatabaseStatus `json:"status,omitempty"`
}

// SpannerDatabaseSpec defines the desired state of SpannerDatabase.
type SpannerDatabaseSpec struct {
	// Immutable. The dialect of the Cloud Spanner Database.
	// If it is not provided, "GOOGLE_STANDARD_SQL" will be used. Possible values: ["GOOGLE_STANDARD_SQL", "POSTGRESQL"].
	DatabaseDialect string `json:"databaseDialect,omitempty" yaml:"databaseDialect,omitempty"`
	// An optional list of DDL statements to run inside the newly created
	// database. Statements can create tables, indexes, etc. These statements
	// execute atomically with the creation of the database: if there is an
	// error in any statement, the database is not created.
	Ddl []string `json:"ddl,omitempty" yaml:"ddl,omitempty"`
	EnableDropProtection bool `json:"enableDropProtection,omitempty" yaml:"enableDropProtection,omitempty"`
	// Immutable. Encryption configuration for the database.
	EncryptionConfig map[string]interface{} `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`
	// The instance to create the database on.
	InstanceRef map[string]interface{} `json:"instanceRef,omitempty" yaml:"instanceRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The retention period for the database. The retention period must be between 1 hour
	// and 7 days, and can be specified in days, hours, minutes, or seconds. For example,
	// the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is 1h.
	// If this property is used, you must avoid adding new DDL statements to 'ddl' that
	// update the database's version_retention_period.
	VersionRetentionPeriod string `json:"versionRetentionPeriod,omitempty" yaml:"versionRetentionPeriod,omitempty"`
}

// SpannerDatabaseStatus defines the observed state of SpannerDatabase.
type SpannerDatabaseStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
