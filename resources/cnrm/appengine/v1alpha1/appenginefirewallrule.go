package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppEngineFirewallRule represents a appengine.cnrm.cloud.google.com AppEngineFirewallRule resource.
type AppEngineFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppEngineFirewallRuleSpec   `json:"spec,omitempty"`
	Status AppEngineFirewallRuleStatus `json:"status,omitempty"`
}

// AppEngineFirewallRuleSpec defines the desired state of AppEngineFirewallRule.
type AppEngineFirewallRuleSpec struct {
	// The action to take if this rule matches. Possible values: ["UNSPECIFIED_ACTION", "ALLOW", "DENY"].
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
	// An optional string description of this rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
	// Immutable. Optional. The priority of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange string `json:"sourceRange,omitempty" yaml:"sourceRange,omitempty"`
}

// AppEngineFirewallRuleStatus defines the observed state of AppEngineFirewallRule.
type AppEngineFirewallRuleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
