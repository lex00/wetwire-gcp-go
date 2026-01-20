package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeResourcePolicy represents a compute.cnrm.cloud.google.com ComputeResourcePolicy resource.
type ComputeResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeResourcePolicySpec   `json:"spec,omitempty"`
	Status ComputeResourcePolicyStatus `json:"status,omitempty"`
}

// ComputeResourcePolicySpec defines the desired state of ComputeResourcePolicy.
type ComputeResourcePolicySpec struct {
	// Immutable. An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Replication consistency group for asynchronous disk replication.
	DiskConsistencyGroupPolicy map[string]interface{} `json:"diskConsistencyGroupPolicy,omitempty" yaml:"diskConsistencyGroupPolicy,omitempty"`
	// Immutable. Resource policy for instances used for placement configuration.
	GroupPlacementPolicy map[string]interface{} `json:"groupPlacementPolicy,omitempty" yaml:"groupPlacementPolicy,omitempty"`
	// Immutable. Resource policy for scheduling instance operations.
	InstanceSchedulePolicy map[string]interface{} `json:"instanceSchedulePolicy,omitempty" yaml:"instanceSchedulePolicy,omitempty"`
	// Immutable. Region where resource policy resides.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Policy for creating snapshots of persistent disks.
	SnapshotSchedulePolicy map[string]interface{} `json:"snapshotSchedulePolicy,omitempty" yaml:"snapshotSchedulePolicy,omitempty"`
}

// ComputeResourcePolicyStatus defines the observed state of ComputeResourcePolicy.
type ComputeResourcePolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
