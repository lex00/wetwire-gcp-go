package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BatchJob is the Schema for the BatchJob API
type BatchJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BatchJobSpec   `json:"spec,omitempty"`
	Status BatchJobStatus `json:"status,omitempty"`
}

// BatchJobSpec defines the desired state of BatchJob.
type BatchJobSpec struct {
	// Compute resource allocation for all TaskGroups in the Job.
	AllocationPolicy map[string]interface{} `json:"allocationPolicy,omitempty" yaml:"allocationPolicy,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Log preservation policy for the Job.
	LogsPolicy map[string]interface{} `json:"logsPolicy,omitempty" yaml:"logsPolicy,omitempty"`
	// Notification configurations.
	Notifications []map[string]interface{} `json:"notifications,omitempty" yaml:"notifications,omitempty"`
	// Priority of the Job. The valid value range is [0, 100). Default value is 0. Higher value indicates higher priority. A job with higher priority value is more likely to run earlier if all other requirements are satisfied.
	Priority int64 `json:"priority,omitempty" yaml:"priority,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The BatchJob name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. TaskGroups in the Job. Only one TaskGroup is supported now.
	TaskGroups []map[string]interface{} `json:"taskGroups,omitempty" yaml:"taskGroups,omitempty"`
}

// BatchJobStatus defines the observed state of BatchJob.
type BatchJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
