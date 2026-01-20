package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpannerBackupSchedule is the Schema for the SpannerBackupSchedule API
type SpannerBackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpannerBackupScheduleSpec   `json:"spec,omitempty"`
	Status SpannerBackupScheduleStatus `json:"status,omitempty"`
}

// SpannerBackupScheduleSpec defines the desired state of SpannerBackupSchedule.
type SpannerBackupScheduleSpec struct {
	// Optional. The encryption configuration that will be used to encrypt the backup. If this field is not specified, the backup will use the same encryption configuration as the database.
	EncryptionConfig map[string]interface{} `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`
	// The schedule creates only full backups.
	FullBackupSpec map[string]interface{} `json:"fullBackupSpec,omitempty" yaml:"fullBackupSpec,omitempty"`
	// The schedule creates incremental backup chains.
	IncrementalBackupSpec map[string]interface{} `json:"incrementalBackupSpec,omitempty" yaml:"incrementalBackupSpec,omitempty"`
	// The SpannerBackupSchedule name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. The retention duration of a backup that must be at least 6 hours and at most 366 days. The backup is eligible to be automatically deleted once the retention period has elapsed.
	RetentionDuration string `json:"retentionDuration,omitempty" yaml:"retentionDuration,omitempty"`
	// Required. The Spanner Database that this backup applies to.
	SpannerDatabaseRef map[string]interface{} `json:"spannerDatabaseRef,omitempty" yaml:"spannerDatabaseRef,omitempty"`
	// Optional. The schedule specification based on which the backup creations are triggered.
	Spec map[string]interface{} `json:"spec,omitempty" yaml:"spec,omitempty"`
}

// SpannerBackupScheduleStatus defines the observed state of SpannerBackupSchedule.
type SpannerBackupScheduleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
