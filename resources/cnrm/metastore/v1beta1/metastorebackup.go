package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MetastoreBackup is the Schema for the MetastoreBackup API
type MetastoreBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MetastoreBackupSpec   `json:"spec,omitempty"`
	Status MetastoreBackupStatus `json:"status,omitempty"`
}

// MetastoreBackupSpec defines the desired state of MetastoreBackup.
type MetastoreBackupSpec struct {
	// The description of the backup.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The MetastoreBackup name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The MetastoreService that the backup belongs to.
	ServiceRef map[string]interface{} `json:"serviceRef,omitempty" yaml:"serviceRef,omitempty"`
}

// MetastoreBackupStatus defines the observed state of MetastoreBackup.
type MetastoreBackupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
