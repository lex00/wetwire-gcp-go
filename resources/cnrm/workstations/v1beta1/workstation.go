package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Workstation is the Schema for the Workstation API
type Workstation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkstationSpec   `json:"spec,omitempty"`
	Status WorkstationStatus `json:"status,omitempty"`
}

// WorkstationSpec defines the desired state of Workstation.
type WorkstationSpec struct {
	// Optional. Client-specified annotations.
	Annotations []map[string]interface{} `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	// Optional. Human-readable name for this workstation.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. [Labels](https://cloud.google.com/workstations/docs/label-resources) that are applied to the workstation and that are also propagated to the underlying Compute Engine resources.
	Labels []map[string]interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Parent is a reference to the parent WorkstationConfig for this Workstation.
	ParentRef map[string]interface{} `json:"parentRef,omitempty" yaml:"parentRef,omitempty"`
	// The Workstation name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// WorkstationStatus defines the observed state of Workstation.
type WorkstationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
