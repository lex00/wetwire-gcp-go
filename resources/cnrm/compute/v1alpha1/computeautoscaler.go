package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeAutoscaler represents a compute.cnrm.cloud.google.com ComputeAutoscaler resource.
type ComputeAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeAutoscalerSpec   `json:"spec,omitempty"`
	Status ComputeAutoscalerStatus `json:"status,omitempty"`
}

// ComputeAutoscalerSpec defines the desired state of ComputeAutoscaler.
type ComputeAutoscalerSpec struct {
	// The configuration parameters for the autoscaling algorithm. You can
	// define one or more of the policies for an autoscaler: cpuUtilization,
	// customMetricUtilizations, and loadBalancingUtilization.
	// If none of these are specified, the default will be to autoscale based
	// on cpuUtilization to 0.6 or 60%.
	AutoscalingPolicy map[string]interface{} `json:"autoscalingPolicy,omitempty" yaml:"autoscalingPolicy,omitempty"`
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	TargetRef map[string]interface{} `json:"targetRef,omitempty" yaml:"targetRef,omitempty"`
	// Immutable. URL of the zone where the instance group resides.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputeAutoscalerStatus defines the observed state of ComputeAutoscaler.
type ComputeAutoscalerStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
