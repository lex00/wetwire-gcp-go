package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeDiskResourcePolicyAttachment represents a compute.cnrm.cloud.google.com ComputeDiskResourcePolicyAttachment resource.
type ComputeDiskResourcePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeDiskResourcePolicyAttachmentSpec   `json:"spec,omitempty"`
	Status ComputeDiskResourcePolicyAttachmentStatus `json:"status,omitempty"`
}

// ComputeDiskResourcePolicyAttachmentSpec defines the desired state of ComputeDiskResourcePolicyAttachment.
type ComputeDiskResourcePolicyAttachmentSpec struct {
	DiskRef map[string]interface{} `json:"diskRef,omitempty" yaml:"diskRef,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. A reference to the zone where the disk resides.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// ComputeDiskResourcePolicyAttachmentStatus defines the observed state of ComputeDiskResourcePolicyAttachment.
type ComputeDiskResourcePolicyAttachmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
