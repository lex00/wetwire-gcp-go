package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeSharedVPCServiceProject represents a compute.cnrm.cloud.google.com ComputeSharedVPCServiceProject resource.
type ComputeSharedVPCServiceProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeSharedVPCServiceProjectSpec   `json:"spec,omitempty"`
	Status ComputeSharedVPCServiceProjectStatus `json:"status,omitempty"`
}

// ComputeSharedVPCServiceProjectSpec defines the desired state of ComputeSharedVPCServiceProject.
type ComputeSharedVPCServiceProjectSpec struct {
	// The deletion policy for the shared VPC service. Setting ABANDON allows the resource
	// to be abandoned rather than deleted. Possible values are: "ABANDON".
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
}

// ComputeSharedVPCServiceProjectStatus defines the observed state of ComputeSharedVPCServiceProject.
type ComputeSharedVPCServiceProjectStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
