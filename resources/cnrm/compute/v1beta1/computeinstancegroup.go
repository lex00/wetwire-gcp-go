package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeInstanceGroup represents a compute.cnrm.cloud.google.com ComputeInstanceGroup resource.
type ComputeInstanceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeInstanceGroupSpec   `json:"spec,omitempty"`
	Status ComputeInstanceGroupStatus `json:"status,omitempty"`
}

// ComputeInstanceGroupSpec defines the desired state of ComputeInstanceGroup.
type ComputeInstanceGroupSpec struct {
	// Immutable. An optional textual description of the instance group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Instances []map[string]interface{} `json:"instances,omitempty" yaml:"instances,omitempty"`
	// The named port configuration.
	NamedPort []map[string]interface{} `json:"namedPort,omitempty" yaml:"namedPort,omitempty"`
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The zone that this instance group should be created in.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputeInstanceGroupStatus defines the observed state of ComputeInstanceGroup.
type ComputeInstanceGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
