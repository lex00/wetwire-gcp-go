package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GKEBackupRestorePlan is the Schema for the GKEBackupRestorePlan API
type GKEBackupRestorePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKEBackupRestorePlanSpec   `json:"spec,omitempty"`
	Status GKEBackupRestorePlanStatus `json:"status,omitempty"`
}

// GKEBackupRestorePlanSpec defines the desired state of GKEBackupRestorePlan.
type GKEBackupRestorePlanSpec struct {
	// Required. Immutable. A reference to the [BackupPlan][google.cloud.gkebackup.v1.BackupPlan] from which Backups may be used as the source for Restores created via this RestorePlan.
	BackupPlanRef map[string]interface{} `json:"backupPlanRef,omitempty" yaml:"backupPlanRef,omitempty"`
	// Required. Immutable. The target cluster into which Restores created via this RestorePlan will restore data. NOTE: the cluster's region must be the same as the RestorePlan.
	ClusterRef map[string]interface{} `json:"clusterRef,omitempty" yaml:"clusterRef,omitempty"`
	// Optional. User specified descriptive string for this RestorePlan.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. A set of custom labels supplied by user.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The GKEBackupRestorePlan name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Configuration of Restores created via this RestorePlan.
	RestoreConfig map[string]interface{} `json:"restoreConfig,omitempty" yaml:"restoreConfig,omitempty"`
}

// GKEBackupRestorePlanStatus defines the observed state of GKEBackupRestorePlan.
type GKEBackupRestorePlanStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
