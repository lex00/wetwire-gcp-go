package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CloudIdentityMembership is the Schema for the CloudIdentityMembership API
type CloudIdentityMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudIdentityMembershipSpec   `json:"spec,omitempty"`
	Status CloudIdentityMembershipStatus `json:"status,omitempty"`
}

// CloudIdentityMembershipSpec defines the desired state of CloudIdentityMembership.
type CloudIdentityMembershipSpec struct {
	// Immutable.
	GroupRef map[string]interface{} `json:"groupRef,omitempty" yaml:"groupRef,omitempty"`
	// Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned. Do not set. This is a legacy OUTPUT-ONLY field.
	MemberKey map[string]interface{} `json:"memberKey,omitempty" yaml:"memberKey,omitempty"`
	// Required. Immutable. The `EntityKey` of the member. Either `member_key` or `preferred_member_key` must be set when calling MembershipsService.CreateMembership but not both; both shall be set when returned.
	PreferredMemberKey map[string]interface{} `json:"preferredMemberKey,omitempty" yaml:"preferredMemberKey,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Format: {membershipID}. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.
	Roles []map[string]interface{} `json:"roles,omitempty" yaml:"roles,omitempty"`
}

// CloudIdentityMembershipStatus defines the observed state of CloudIdentityMembership.
type CloudIdentityMembershipStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
