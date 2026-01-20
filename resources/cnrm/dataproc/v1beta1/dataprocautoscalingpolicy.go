package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataprocAutoscalingPolicy represents a dataproc.cnrm.cloud.google.com DataprocAutoscalingPolicy resource.
type DataprocAutoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataprocAutoscalingPolicySpec   `json:"spec,omitempty"`
	Status DataprocAutoscalingPolicyStatus `json:"status,omitempty"`
}

// DataprocAutoscalingPolicySpec defines the desired state of DataprocAutoscalingPolicy.
type DataprocAutoscalingPolicySpec struct {
	BasicAlgorithm map[string]interface{} `json:"basicAlgorithm,omitempty" yaml:"basicAlgorithm,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig map[string]interface{} `json:"secondaryWorkerConfig,omitempty" yaml:"secondaryWorkerConfig,omitempty"`
	// Required. Describes how the autoscaler will operate for primary workers.
	WorkerConfig map[string]interface{} `json:"workerConfig,omitempty" yaml:"workerConfig,omitempty"`
}

// DataprocAutoscalingPolicyStatus defines the observed state of DataprocAutoscalingPolicy.
type DataprocAutoscalingPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
