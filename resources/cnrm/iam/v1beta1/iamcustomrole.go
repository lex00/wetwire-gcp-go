package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMCustomRole represents a iam.cnrm.cloud.google.com IAMCustomRole resource.
type IAMCustomRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMCustomRoleSpec   `json:"spec,omitempty"`
	Status IAMCustomRoleStatus `json:"status,omitempty"`
}

// IAMCustomRoleSpec defines the desired state of IAMCustomRole.
type IAMCustomRoleSpec struct {
	// A human-readable description for the role.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
	Permissions []string `json:"permissions,omitempty" yaml:"permissions,omitempty"`
	// Immutable. Optional. The roleId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The current launch stage of the role. Defaults to GA.
	Stage string `json:"stage,omitempty" yaml:"stage,omitempty"`
	// A human-readable title for the role.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}

// IAMCustomRoleStatus defines the observed state of IAMCustomRole.
type IAMCustomRoleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
