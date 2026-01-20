package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMDenyPolicy is the Schema for the IAMDenyPolicy API
type IAMDenyPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMDenyPolicySpec   `json:"spec,omitempty"`
	Status IAMDenyPolicyStatus `json:"status,omitempty"`
}

// IAMDenyPolicySpec defines the desired state of IAMDenyPolicy.
type IAMDenyPolicySpec struct {
	// A user-specified description of the `Policy`. This value can be up to 63 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The IAMDenyPolicy name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A list of rules that specify the behavior of the `Policy`. All of the rules should be of the `kind` specified in the `Policy`.
	Rules []map[string]interface{} `json:"rules,omitempty" yaml:"rules,omitempty"`
}

// IAMDenyPolicyStatus defines the observed state of IAMDenyPolicy.
type IAMDenyPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
