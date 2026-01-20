package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMAccessBoundaryPolicy represents a iam.cnrm.cloud.google.com IAMAccessBoundaryPolicy resource.
type IAMAccessBoundaryPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMAccessBoundaryPolicySpec   `json:"spec,omitempty"`
	Status IAMAccessBoundaryPolicyStatus `json:"status,omitempty"`
}

// IAMAccessBoundaryPolicySpec defines the desired state of IAMAccessBoundaryPolicy.
type IAMAccessBoundaryPolicySpec struct {
	// The display name of the rule.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Rules to be applied.
	Rules []map[string]interface{} `json:"rules,omitempty" yaml:"rules,omitempty"`
}

// IAMAccessBoundaryPolicyStatus defines the observed state of IAMAccessBoundaryPolicy.
type IAMAccessBoundaryPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
