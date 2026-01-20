package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMServiceAccount represents a iam.cnrm.cloud.google.com IAMServiceAccount resource.
type IAMServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMServiceAccountSpec   `json:"spec,omitempty"`
	Status IAMServiceAccountStatus `json:"status,omitempty"`
}

// IAMServiceAccountSpec defines the desired state of IAMServiceAccount.
type IAMServiceAccountSpec struct {
	// A text description of the service account. Must be less than or equal to 256 UTF-8 bytes.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Whether the service account is disabled. Defaults to false.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// The display name for the service account. Can be updated without creating a new resource.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. Optional. The accountId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// IAMServiceAccountStatus defines the observed state of IAMServiceAccount.
type IAMServiceAccountStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
