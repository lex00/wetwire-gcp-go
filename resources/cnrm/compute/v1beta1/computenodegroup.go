package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNodeGroup represents a compute.cnrm.cloud.google.com ComputeNodeGroup resource.
type ComputeNodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNodeGroupSpec   `json:"spec,omitempty"`
	Status ComputeNodeGroupStatus `json:"status,omitempty"`
}

// ComputeNodeGroupSpec defines the desired state of ComputeNodeGroup.
type ComputeNodeGroupSpec struct {
	// Immutable. If you use sole-tenant nodes for your workloads, you can use the node
	// group autoscaler to automatically manage the sizes of your node groups.
	AutoscalingPolicy map[string]interface{} `json:"autoscalingPolicy,omitempty" yaml:"autoscalingPolicy,omitempty"`
	// Immutable. An optional textual description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The initial number of nodes in the node group. One of 'initial_size' or 'size' must be specified.
	InitialSize int32 `json:"initialSize,omitempty" yaml:"initialSize,omitempty"`
	// Immutable. Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
	MaintenancePolicy string `json:"maintenancePolicy,omitempty" yaml:"maintenancePolicy,omitempty"`
	// Immutable. contains properties for the timeframe of maintenance.
	MaintenanceWindow map[string]interface{} `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`
	// The node template to which this node group belongs.
	NodeTemplateRef map[string]interface{} `json:"nodeTemplateRef,omitempty" yaml:"nodeTemplateRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. Share settings for the node group.
	ShareSettings map[string]interface{} `json:"shareSettings,omitempty" yaml:"shareSettings,omitempty"`
	// Immutable. The total number of nodes in the node group. One of 'initial_size' or 'size' must be specified.
	Size int32 `json:"size,omitempty" yaml:"size,omitempty"`
	// Immutable. Zone where this node group is located.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputeNodeGroupStatus defines the observed state of ComputeNodeGroup.
type ComputeNodeGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
