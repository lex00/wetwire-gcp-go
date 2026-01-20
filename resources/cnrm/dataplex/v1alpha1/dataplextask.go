package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataplexTask is the Schema for the DataplexTask API
type DataplexTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataplexTaskSpec   `json:"spec,omitempty"`
	Status DataplexTaskStatus `json:"status,omitempty"`
}

// DataplexTaskSpec defines the desired state of DataplexTask.
type DataplexTaskSpec struct {
	// Optional. Description of the task.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. User friendly display name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Required. Spec related to how a task is executed.
	ExecutionSpec map[string]interface{} `json:"executionSpec,omitempty" yaml:"executionSpec,omitempty"`
	// LakeRef defines the resource reference to DataplexLake, which "External" field holds the GCP identifier for the KRM object.
	LakeRef map[string]interface{} `json:"lakeRef,omitempty" yaml:"lakeRef,omitempty"`
	// Config related to running scheduled Notebooks. Exactly one of spark or notebook must be set.
	Notebook map[string]interface{} `json:"notebook,omitempty" yaml:"notebook,omitempty"`
	// The DataplexTask name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Config related to running custom Spark tasks. Exactly one of spark or notebook must be set.
	Spark map[string]interface{} `json:"spark,omitempty" yaml:"spark,omitempty"`
	// Required. Spec related to how often and when a task should be triggered.
	TriggerSpec map[string]interface{} `json:"triggerSpec,omitempty" yaml:"triggerSpec,omitempty"`
}

// DataplexTaskStatus defines the observed state of DataplexTask.
type DataplexTaskStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
