package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupDRBackupVault is the Schema for the BackupDRBackupVault API
type BackupDRBackupVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupDRBackupVaultSpec   `json:"spec,omitempty"`
	Status BackupDRBackupVaultStatus `json:"status,omitempty"`
}

// BackupDRBackupVaultSpec defines the desired state of BackupDRBackupVault.
type BackupDRBackupVaultSpec struct {
	// Optional. Note: This field is added for future use case and will not be
	// supported in the current release.
	// Access restriction for the backup vault.
	// Default value is WITHIN_ORGANIZATION if not provided during creation.
	AccessRestriction string `json:"accessRestriction,omitempty" yaml:"accessRestriction,omitempty"`
	// Optional. User annotations. See https://google.aip.dev/128#annotations Stores small amounts of arbitrary data.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Required. The default and minimum enforced retention for each backup within the backup vault.  The enforced retention for each backup can be extended.
	BackupMinimumEnforcedRetentionDuration string `json:"backupMinimumEnforcedRetentionDuration,omitempty" yaml:"backupMinimumEnforcedRetentionDuration,omitempty"`
	// Optional. The description of the BackupVault instance (2048 characters or less).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Time after which the BackupVault resource is locked.
	EffectiveTime string `json:"effectiveTime,omitempty" yaml:"effectiveTime,omitempty"`
	// Optional. If set to true, allows deletion of a backup vault even when it contains inactive data sources. This overrides the default restriction that prevents deletion of backup vaults with any data sources, even if those data sources are inactive.
	IgnoreInactiveDatasources bool `json:"ignoreInactiveDatasources,omitempty" yaml:"ignoreInactiveDatasources,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The BackupDRBackupVault name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BackupDRBackupVaultStatus defines the observed state of BackupDRBackupVault.
type BackupDRBackupVaultStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
