package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupDRBackupPlan is the Schema for the BackupDRBackupPlan API
type BackupDRBackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupDRBackupPlanSpec   `json:"spec,omitempty"`
	Status BackupDRBackupPlanStatus `json:"status,omitempty"`
}

// BackupDRBackupPlanSpec defines the desired state of BackupDRBackupPlan.
type BackupDRBackupPlanSpec struct {
	// Required. The backup rules for this `BackupPlan`. There must be at least one `BackupRule` message.
	BackupRules []map[string]interface{} `json:"backupRules,omitempty" yaml:"backupRules,omitempty"`
	// Required. Resource name of backup vault which will be used as storage location for backups.
	BackupVaultRef map[string]interface{} `json:"backupVaultRef,omitempty" yaml:"backupVaultRef,omitempty"`
	// Optional. The description of the `BackupPlan` resource.
	// The description allows for additional details about `BackupPlan` and its
	// use cases to be provided. An example description is the following:  "This
	// is a backup plan that performs a daily backup at 6pm and retains data for 3
	// months". The description must be at most 2048 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The BackupDRBackupPlan name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The resource type to which the `BackupPlan` will be applied. Examples include, "compute.googleapis.com/Instance", "sqladmin.googleapis.com/Instance", or "alloydb.googleapis.com/Cluster".
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`
}

// BackupDRBackupPlanStatus defines the observed state of BackupDRBackupPlan.
type BackupDRBackupPlanStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
