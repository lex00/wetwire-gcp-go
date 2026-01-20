package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupDRBackupPlanAssociation is the Schema for the BackupDRBackupPlanAssociation API
type BackupDRBackupPlanAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupDRBackupPlanAssociationSpec   `json:"spec,omitempty"`
	Status BackupDRBackupPlanAssociationStatus `json:"status,omitempty"`
}

// BackupDRBackupPlanAssociationSpec defines the desired state of BackupDRBackupPlanAssociation.
type BackupDRBackupPlanAssociationSpec struct {
	// Required. The backup plan which needs to be applied on workload.
	BackupPlanRef map[string]interface{} `json:"backupPlanRef,omitempty" yaml:"backupPlanRef,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Required. Immutable. Resource name of workload on which backupplan is applied
	Resource map[string]interface{} `json:"resource,omitempty" yaml:"resource,omitempty"`
	// The BackupDRBackupPlanAssociation name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Immutable. Resource type of workload on which backupplan is applied
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`
}

// BackupDRBackupPlanAssociationStatus defines the observed state of BackupDRBackupPlanAssociation.
type BackupDRBackupPlanAssociationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
