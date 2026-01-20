package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeInstanceGroupNamedPort represents a compute.cnrm.cloud.google.com ComputeInstanceGroupNamedPort resource.
type ComputeInstanceGroupNamedPort struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInstanceGroupNamedPortSpec   `json:"spec,omitempty"`
	Status ComputeInstanceGroupNamedPortStatus `json:"status,omitempty"`
}

// ComputeInstanceGroupNamedPortSpec defines the desired state of ComputeInstanceGroupNamedPort.
type ComputeInstanceGroupNamedPortSpec struct {
	GroupRef map[string]interface{} `json:"groupRef,omitempty" yaml:"groupRef,omitempty"`
	// Immutable. The port number, which can be a value between 1 and 65535.
	Port int32 `json:"port,omitempty" yaml:"port,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The zone of the instance group.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputeInstanceGroupNamedPortStatus defines the observed state of ComputeInstanceGroupNamedPort.
type ComputeInstanceGroupNamedPortStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
