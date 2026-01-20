package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GKEBackupBackupPlan is the Schema for the GKEBackupBackupPlan API
type GKEBackupBackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKEBackupBackupPlanSpec   `json:"spec,omitempty"`
	Status GKEBackupBackupPlanStatus `json:"status,omitempty"`
}

// GKEBackupBackupPlanSpec defines the desired state of GKEBackupBackupPlan.
type GKEBackupBackupPlanSpec struct {
	// Optional. Defines the configuration of Backups created via this BackupPlan.
	BackupConfig map[string]interface{} `json:"backupConfig,omitempty" yaml:"backupConfig,omitempty"`
	// Optional. Defines a schedule for automatic Backup creation via this BackupPlan.
	BackupSchedule map[string]interface{} `json:"backupSchedule,omitempty" yaml:"backupSchedule,omitempty"`
	// Required. Immutable. The source cluster from which Backups will be created via this BackupPlan.
	ClusterRef map[string]interface{} `json:"clusterRef,omitempty" yaml:"clusterRef,omitempty"`
	// Optional. This flag indicates whether this BackupPlan has been deactivated.
	// Setting this field to True locks the BackupPlan such that no further
	// updates will be allowed (except deletes), including the deactivated field
	// itself. It also prevents any new Backups from being created via this
	// BackupPlan (including scheduled Backups).
	// Default: False
	Deactivated bool `json:"deactivated,omitempty" yaml:"deactivated,omitempty"`
	// Optional. User specified descriptive string for this BackupPlan.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. A set of custom labels supplied by user.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The GKEBackupBackupPlan name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. RetentionPolicy governs lifecycle of Backups created under this plan.
	RetentionPolicy map[string]interface{} `json:"retentionPolicy,omitempty" yaml:"retentionPolicy,omitempty"`
}

// GKEBackupBackupPlanStatus defines the observed state of GKEBackupBackupPlan.
type GKEBackupBackupPlanStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
