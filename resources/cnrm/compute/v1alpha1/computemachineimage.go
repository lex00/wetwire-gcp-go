package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeMachineImage represents a compute.cnrm.cloud.google.com ComputeMachineImage resource.
type ComputeMachineImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeMachineImageSpec   `json:"spec,omitempty"`
	Status ComputeMachineImageStatus `json:"status,omitempty"`
}

// ComputeMachineImageSpec defines the desired state of ComputeMachineImage.
type ComputeMachineImageSpec struct {
	// Immutable. A text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	// Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	GuestFlush bool `json:"guestFlush,omitempty" yaml:"guestFlush,omitempty"`
	// Immutable. Encrypts the machine image using a customer-supplied encryption key.
	// After you encrypt a machine image with a customer-supplied key, you must
	// provide the same key if you use the machine image later (e.g. to create a
	// instance from the image).
	MachineImageEncryptionKey map[string]interface{} `json:"machineImageEncryptionKey,omitempty" yaml:"machineImageEncryptionKey,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	SourceInstanceRef map[string]interface{} `json:"sourceInstanceRef,omitempty" yaml:"sourceInstanceRef,omitempty"`
}

// ComputeMachineImageStatus defines the observed state of ComputeMachineImage.
type ComputeMachineImageStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
