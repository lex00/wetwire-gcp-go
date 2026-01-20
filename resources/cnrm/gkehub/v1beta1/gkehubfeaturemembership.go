package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GKEHubFeatureMembership is the Schema for the gkehub API
type GKEHubFeatureMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKEHubFeatureMembershipSpec   `json:"spec,omitempty"`
	Status GKEHubFeatureMembershipStatus `json:"status,omitempty"`
}

// GKEHubFeatureMembershipSpec defines the desired state of GKEHubFeatureMembership.
type GKEHubFeatureMembershipSpec struct {
	// Config Management-specific spec.
	Configmanagement map[string]interface{} `json:"configmanagement,omitempty" yaml:"configmanagement,omitempty"`
	// Immutable.
	FeatureRef map[string]interface{} `json:"featureRef,omitempty" yaml:"featureRef,omitempty"`
	// Immutable. The location of the feature
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The location of the membership
	MembershipLocation string `json:"membershipLocation,omitempty" yaml:"membershipLocation,omitempty"`
	// Immutable.
	MembershipRef map[string]interface{} `json:"membershipRef,omitempty" yaml:"membershipRef,omitempty"`
	// Manage Mesh Features
	Mesh map[string]interface{} `json:"mesh,omitempty" yaml:"mesh,omitempty"`
	// Policy Controller-specific spec.
	Policycontroller map[string]interface{} `json:"policycontroller,omitempty" yaml:"policycontroller,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
}

// GKEHubFeatureMembershipStatus defines the observed state of GKEHubFeatureMembership.
type GKEHubFeatureMembershipStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
