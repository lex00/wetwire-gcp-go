package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataprocNodeGroup is the Schema for the DataprocNodeGroup API
type DataprocNodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataprocNodeGroupSpec   `json:"spec,omitempty"`
	Status DataprocNodeGroupStatus `json:"status,omitempty"`
}

// DataprocNodeGroupSpec defines the desired state of DataprocNodeGroup.
type DataprocNodeGroupSpec struct {
	// Optional. Node group labels.
	// * Label **keys** must consist of from 1 to 63 characters and conform to
	// [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	// * Label **values** can be empty. If specified, they must consist of from
	// 1 to 63 characters and conform to [RFC 1035]
	// (https://www.ietf.org/rfc/rfc1035.txt).
	// * The node group must have no more than 32 labels.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Required.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Optional. The node group instance group configuration.
	NodeGroupConfig map[string]interface{} `json:"nodeGroupConfig,omitempty" yaml:"nodeGroupConfig,omitempty"`
	// Required.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The DataprocNodeGroup name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Node group roles.
	Roles []string `json:"roles,omitempty" yaml:"roles,omitempty"`
}

// DataprocNodeGroupStatus defines the observed state of DataprocNodeGroup.
type DataprocNodeGroupStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
