package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AlloyDBBackup is the Schema for the AlloyDBBackup API
type AlloyDBBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlloyDBBackupSpec   `json:"spec,omitempty"`
	Status AlloyDBBackupStatus `json:"status,omitempty"`
}

// AlloyDBBackupSpec defines the desired state of AlloyDBBackup.
type AlloyDBBackupSpec struct {
	// The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
	ClusterNameRef map[string]interface{} `json:"clusterNameRef,omitempty" yaml:"clusterNameRef,omitempty"`
	// Immutable. User-provided description of the backup.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
	EncryptionConfig map[string]interface{} `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`
	// Immutable. The location where the alloydb backup should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The backupId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// AlloyDBBackupStatus defines the observed state of AlloyDBBackup.
type AlloyDBBackupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
