package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GKEBackupBackup is the Schema for the GKEBackupBackup API
type GKEBackupBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKEBackupBackupSpec   `json:"spec,omitempty"`
	Status GKEBackupBackupStatus `json:"status,omitempty"`
}

// GKEBackupBackupSpec defines the desired state of GKEBackupBackup.
type GKEBackupBackupSpec struct {
	// Required. The BackupPlan from which this Backup is created.
	BackupPlanRef map[string]interface{} `json:"backupPlanRef,omitempty" yaml:"backupPlanRef,omitempty"`
	// Optional. Minimum age for this Backup (in days). If this field is set to a
	// non-zero value, the Backup will be "locked" against deletion (either manual
	// or automatic deletion) for the number of days provided (measured from the
	// creation time of the Backup).  MUST be an integer value between 0-90
	// (inclusive).
	// Defaults to parent BackupPlan's
	// [backup_delete_lock_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_delete_lock_days]
	// setting and may only be increased
	// (either at creation time or in a subsequent update).
	DeleteLockDays int32 `json:"deleteLockDays,omitempty" yaml:"deleteLockDays,omitempty"`
	// Optional. User specified descriptive string for this Backup.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. A set of custom labels supplied by user.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// The GKEBackupBackup name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. The age (in days) after which this Backup will be automatically
	// deleted. Must be an integer value >= 0:
	// - If 0, no automatic deletion will occur for this Backup.
	// - If not 0, this must be >=
	// [delete_lock_days][google.cloud.gkebackup.v1.Backup.delete_lock_days] and
	// <= 365.
	// Once a Backup is created, this value may only be increased.
	// Defaults to the parent BackupPlan's
	// [backup_retain_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_retain_days]
	// value.
	RetainDays int32 `json:"retainDays,omitempty" yaml:"retainDays,omitempty"`
}

// GKEBackupBackupStatus defines the observed state of GKEBackupBackup.
type GKEBackupBackupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
