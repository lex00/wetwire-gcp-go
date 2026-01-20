package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeSharedVPCHostProject represents a compute.cnrm.cloud.google.com ComputeSharedVPCHostProject resource.
type ComputeSharedVPCHostProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeSharedVPCHostProjectSpec   `json:"spec,omitempty"`
	Status ComputeSharedVPCHostProjectStatus `json:"status,omitempty"`
}

// ComputeSharedVPCHostProjectSpec defines the desired state of ComputeSharedVPCHostProject.
type ComputeSharedVPCHostProjectSpec struct {
	// TODO: Add spec fields from CRD schema
}

// ComputeSharedVPCHostProjectStatus defines the observed state of ComputeSharedVPCHostProject.
type ComputeSharedVPCHostProjectStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
