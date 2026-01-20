package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetAppBackupVault is the Schema for the NetAppBackupVault API
type NetAppBackupVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetAppBackupVaultSpec   `json:"spec,omitempty"`
	Status NetAppBackupVaultStatus `json:"status,omitempty"`
}

// NetAppBackupVaultSpec defines the desired state of NetAppBackupVault.
type NetAppBackupVaultSpec struct {
	// Description of the backup vault.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The NetAppBackupVault name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// NetAppBackupVaultStatus defines the observed state of NetAppBackupVault.
type NetAppBackupVaultStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
