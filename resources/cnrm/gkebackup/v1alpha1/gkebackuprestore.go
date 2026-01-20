package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GKEBackupRestore is the Schema for the GKEBackupRestore API
type GKEBackupRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKEBackupRestoreSpec   `json:"spec,omitempty"`
	Status GKEBackupRestoreStatus `json:"status,omitempty"`
}

// GKEBackupRestoreSpec defines the desired state of GKEBackupRestore.
type GKEBackupRestoreSpec struct {
	// Required. Immutable. A reference to the [Backup][google.cloud.gkebackup.v1.Backup] used as the source from which this Restore will restore. Note that this Backup must be a sub-resource of the RestorePlan's [backup_plan][google.cloud.gkebackup.v1.RestorePlan.backup_plan].
	BackupRef map[string]interface{} `json:"backupRef,omitempty" yaml:"backupRef,omitempty"`
	// User specified descriptive string for this Restore.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Immutable. Filters resources for `Restore`. If not specified, the scope of the restore will remain the same as defined in the `RestorePlan`. If this is specified, and no resources are matched by the `inclusion_filters` or everything is excluded by the `exclusion_filters`, nothing will be restored. This filter can only be specified if the value of [namespaced_resource_restore_mode][google.cloud.gkebackup.v1.RestoreConfig.namespaced_resource_restore_mode] is set to `MERGE_SKIP_ON_CONFLICT`, `MERGE_REPLACE_VOLUME_ON_CONFLICT` or `MERGE_REPLACE_ON_CONFLICT`.
	Filter map[string]interface{} `json:"filter,omitempty" yaml:"filter,omitempty"`
	// A set of custom labels supplied by user.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// The GKEBackupRestore name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The RestorePlan from which this Restore is created.
	RestorePlanRef map[string]interface{} `json:"restorePlanRef,omitempty" yaml:"restorePlanRef,omitempty"`
	// Optional. Immutable. Overrides the volume data restore policies selected in the Restore Config for override-scoped resources.
	VolumeDataRestorePolicyOverrides []map[string]interface{} `json:"volumeDataRestorePolicyOverrides,omitempty" yaml:"volumeDataRestorePolicyOverrides,omitempty"`
}

// GKEBackupRestoreStatus defines the observed state of GKEBackupRestore.
type GKEBackupRestoreStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
