package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirestoreBackupSchedule is the Schema for the FirestoreBackupSchedule API
type FirestoreBackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirestoreBackupScheduleSpec   `json:"spec,omitempty"`
	Status FirestoreBackupScheduleStatus `json:"status,omitempty"`
}

// FirestoreBackupScheduleSpec defines the desired state of FirestoreBackupSchedule.
type FirestoreBackupScheduleSpec struct {
	// For a schedule that runs daily.
	DailyRecurrence map[string]interface{} `json:"dailyRecurrence,omitempty" yaml:"dailyRecurrence,omitempty"`
	// The database that this resource belongs to.
	DatabaseRef map[string]interface{} `json:"databaseRef,omitempty" yaml:"databaseRef,omitempty"`
	// At what relative time in the future, compared to its creation time,
	// the backup should be deleted, e.g. keep backups for 7 days.
	// The maximum supported retention period is 14 weeks.
	Retention string `json:"retention,omitempty" yaml:"retention,omitempty"`
	// For a schedule that runs weekly on a specific day.
	WeeklyRecurrence map[string]interface{} `json:"weeklyRecurrence,omitempty" yaml:"weeklyRecurrence,omitempty"`
}

// FirestoreBackupScheduleStatus defines the observed state of FirestoreBackupSchedule.
type FirestoreBackupScheduleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
