package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkSecurityAuthorizationPolicy is the Schema for the NetworkSecurityAuthorizationPolicy API
type NetworkSecurityAuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSecurityAuthorizationPolicySpec   `json:"spec,omitempty"`
	Status NetworkSecurityAuthorizationPolicyStatus `json:"status,omitempty"`
}

// NetworkSecurityAuthorizationPolicySpec defines the desired state of NetworkSecurityAuthorizationPolicy.
type NetworkSecurityAuthorizationPolicySpec struct {
	// Required. The action to take when a rule match is found. Possible values are "ALLOW" or "DENY".
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
	// Optional. Free-text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The NetworkSecurityAuthorizationPolicy name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. List of rules to match. Note that at least one of the rules must match in order for the action specified in the 'action' field to be taken. A rule is a match if there is a matching source and destination. If left blank, the action specified in the `action` field will be applied on every request.
	Rules []map[string]interface{} `json:"rules,omitempty" yaml:"rules,omitempty"`
}

// NetworkSecurityAuthorizationPolicyStatus defines the observed state of NetworkSecurityAuthorizationPolicy.
type NetworkSecurityAuthorizationPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
