package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MonitoringGroup represents a monitoring.cnrm.cloud.google.com MonitoringGroup resource.
type MonitoringGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringGroupSpec   `json:"spec,omitempty"`
	Status MonitoringGroupStatus `json:"status,omitempty"`
}

// MonitoringGroupSpec defines the desired state of MonitoringGroup.
type MonitoringGroupSpec struct {
	// A user-assigned name for this group, used only for display purposes.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The filter used to determine which monitored resources belong to this group.
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
	// If true, the members of this group are considered to be a cluster. The system can perform additional analysis on groups that are clusters.
	IsCluster bool `json:"isCluster,omitempty" yaml:"isCluster,omitempty"`
	ParentRef map[string]interface{} `json:"parentRef,omitempty" yaml:"parentRef,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// MonitoringGroupStatus defines the observed state of MonitoringGroup.
type MonitoringGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
