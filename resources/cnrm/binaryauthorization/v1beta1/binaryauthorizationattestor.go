package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BinaryAuthorizationAttestor represents a binaryauthorization.cnrm.cloud.google.com BinaryAuthorizationAttestor resource.
type BinaryAuthorizationAttestor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BinaryAuthorizationAttestorSpec   `json:"spec,omitempty"`
	Status BinaryAuthorizationAttestorStatus `json:"status,omitempty"`
}

// BinaryAuthorizationAttestorSpec defines the desired state of BinaryAuthorizationAttestor.
type BinaryAuthorizationAttestorSpec struct {
	// Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// This specifies how an attestation will be read, and how it will be used during policy enforcement.
	UserOwnedDrydockNote map[string]interface{} `json:"userOwnedDrydockNote,omitempty" yaml:"userOwnedDrydockNote,omitempty"`
}

// BinaryAuthorizationAttestorStatus defines the observed state of BinaryAuthorizationAttestor.
type BinaryAuthorizationAttestorStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
