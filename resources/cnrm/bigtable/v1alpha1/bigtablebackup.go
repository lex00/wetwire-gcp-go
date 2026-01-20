package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigtableBackup is the Schema for the BigtableBackup API
type BigtableBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableBackupSpec   `json:"spec,omitempty"`
	Status BigtableBackupStatus `json:"status,omitempty"`
}

// BigtableBackupSpec defines the desired state of BigtableBackup.
type BigtableBackupSpec struct {
	// Indicates the backup type of the backup.
	BackupType string `json:"backupType,omitempty" yaml:"backupType,omitempty"`
	// ClusterRef defines the resource reference to BigtableCluster, which "External" field holds the GCP identifier for the KRM object.
	ClusterRef map[string]interface{} `json:"clusterRef,omitempty" yaml:"clusterRef,omitempty"`
	// Required. The expiration time of the backup.
	// When creating a backup or updating its `expire_time`, the value must be
	// greater than the backup creation time by:
	// - At least 6 hours
	// - At most 90 days
	// Once the `expire_time` has passed, Cloud Bigtable will delete the backup.
	ExpireTime string `json:"expireTime,omitempty" yaml:"expireTime,omitempty"`
	// The time at which the hot backup will be converted to a standard backup.
	// Once the `hot_to_standard_time` has passed, Cloud Bigtable will convert the
	// hot backup to a standard backup. This value must be greater than the backup
	// creation time by:
	// - At least 24 hours
	// This field only applies for hot backups. When creating or updating a
	// standard backup, attempting to set this field will fail the request.
	HotToStandardTime string `json:"hotToStandardTime,omitempty" yaml:"hotToStandardTime,omitempty"`
	// The BigtableBackup name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Immutable. Name of the table from which this backup was created. This needs to be in the same instance as the backup. Values are of the form `projects/{project}/instances/{instance}/tables/{source_table}`.
	SourceTableRef map[string]interface{} `json:"sourceTableRef,omitempty" yaml:"sourceTableRef,omitempty"`
}

// BigtableBackupStatus defines the observed state of BigtableBackup.
type BigtableBackupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
