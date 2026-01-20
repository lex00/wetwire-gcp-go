package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeRegionDiskResourcePolicyAttachment represents a compute.cnrm.cloud.google.com ComputeRegionDiskResourcePolicyAttachment resource.
type ComputeRegionDiskResourcePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRegionDiskResourcePolicyAttachmentSpec   `json:"spec,omitempty"`
	Status ComputeRegionDiskResourcePolicyAttachmentStatus `json:"status,omitempty"`
}

// ComputeRegionDiskResourcePolicyAttachmentSpec defines the desired state of ComputeRegionDiskResourcePolicyAttachment.
type ComputeRegionDiskResourcePolicyAttachmentSpec struct {
	DiskRef map[string]interface{} `json:"diskRef,omitempty" yaml:"diskRef,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. A reference to the region where the disk resides.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeRegionDiskResourcePolicyAttachmentStatus defines the observed state of ComputeRegionDiskResourcePolicyAttachment.
type ComputeRegionDiskResourcePolicyAttachmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
