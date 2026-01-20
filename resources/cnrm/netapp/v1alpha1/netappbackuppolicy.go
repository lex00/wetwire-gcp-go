package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetAppBackupPolicy is the Schema for the NetAppBackupPolicy API
type NetAppBackupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetAppBackupPolicySpec   `json:"spec,omitempty"`
	Status NetAppBackupPolicyStatus `json:"status,omitempty"`
}

// NetAppBackupPolicySpec defines the desired state of NetAppBackupPolicy.
type NetAppBackupPolicySpec struct {
	// Number of daily backups to keep. Note that the minimum daily backup limit is 2.
	DailyBackupLimit int32 `json:"dailyBackupLimit,omitempty" yaml:"dailyBackupLimit,omitempty"`
	// Description of the backup policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// If enabled, make backups automatically according to the schedules. This will be applied to all volumes that have this policy attached and enforced on volume level. If not specified, default is true.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Number of monthly backups to keep. Note that the sum of daily, weekly and monthly backups should be greater than 1.
	MonthlyBackupLimit int32 `json:"monthlyBackupLimit,omitempty" yaml:"monthlyBackupLimit,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The GCP resource identifier. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Number of weekly backups to keep. Note that the sum of daily, weekly and monthly backups should be greater than 1.
	WeeklyBackupLimit int32 `json:"weeklyBackupLimit,omitempty" yaml:"weeklyBackupLimit,omitempty"`
}

// NetAppBackupPolicyStatus defines the observed state of NetAppBackupPolicy.
type NetAppBackupPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
