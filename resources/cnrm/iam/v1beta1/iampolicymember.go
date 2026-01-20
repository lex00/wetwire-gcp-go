package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMPolicyMember is the Schema for the iampolicies API
type IAMPolicyMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMPolicyMemberSpec   `json:"spec,omitempty"`
	Status IAMPolicyMemberStatus `json:"status,omitempty"`
}

// IAMPolicyMemberSpec defines the desired state of IAMPolicyMember.
type IAMPolicyMemberSpec struct {
	// Immutable. Optional. The condition under which the binding applies.
	Condition map[string]interface{} `json:"condition,omitempty" yaml:"condition,omitempty"`
	// Immutable. The IAM identity to be bound to the role. Exactly one of 'member' or 'memberFrom' must be used.
	Member string `json:"member,omitempty" yaml:"member,omitempty"`
	// Immutable. The IAM identity to be bound to the role. Exactly one of 'member' or 'memberFrom' must be used, and only one subfield within 'memberFrom' can be used.
	MemberFrom map[string]interface{} `json:"memberFrom,omitempty" yaml:"memberFrom,omitempty"`
	// Immutable. Required. The GCP resource to set the IAM policy on (e.g. organization, project...)
	ResourceRef map[string]interface{} `json:"resourceRef,omitempty" yaml:"resourceRef,omitempty"`
	// Immutable. Required. The role for which the Member will be bound.
	Role string `json:"role,omitempty" yaml:"role,omitempty"`
}

// IAMPolicyMemberStatus defines the observed state of IAMPolicyMember.
type IAMPolicyMemberStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
