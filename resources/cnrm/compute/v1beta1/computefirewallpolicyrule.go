package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeFirewallPolicyRule is the Schema for the compute API
type ComputeFirewallPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeFirewallPolicyRuleSpec   `json:"spec,omitempty"`
	Status ComputeFirewallPolicyRuleStatus `json:"status,omitempty"`
}

// ComputeFirewallPolicyRuleSpec defines the desired state of ComputeFirewallPolicyRule.
type ComputeFirewallPolicyRuleSpec struct {
	// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
	// An optional description for this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`
	// Immutable.
	FirewallPolicyRef map[string]interface{} `json:"firewallPolicyRef,omitempty" yaml:"firewallPolicyRef,omitempty"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match map[string]interface{} `json:"match,omitempty" yaml:"match,omitempty"`
	// Immutable. An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority int64 `json:"priority,omitempty" yaml:"priority,omitempty"`
	TargetResources []map[string]interface{} `json:"targetResources,omitempty" yaml:"targetResources,omitempty"`
	TargetServiceAccounts []map[string]interface{} `json:"targetServiceAccounts,omitempty" yaml:"targetServiceAccounts,omitempty"`
}

// ComputeFirewallPolicyRuleStatus defines the observed state of ComputeFirewallPolicyRule.
type ComputeFirewallPolicyRuleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
