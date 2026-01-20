package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudBuildWorkerPool is the Schema for the CloudBuild WorkerPool API
type CloudBuildWorkerPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudBuildWorkerPoolSpec   `json:"spec,omitempty"`
	Status CloudBuildWorkerPoolStatus `json:"status,omitempty"`
}

// CloudBuildWorkerPoolSpec defines the desired state of CloudBuildWorkerPool.
type CloudBuildWorkerPoolSpec struct {
	// A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Legacy Private Pool configuration.
	PrivatePoolV1Config map[string]interface{} `json:"privatePoolV1Config,omitempty" yaml:"privatePoolV1Config,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The GCP resource identifier. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// CloudBuildWorkerPoolStatus defines the observed state of CloudBuildWorkerPool.
type CloudBuildWorkerPoolStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
