package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNetworkFirewallPolicyRule represents a compute.cnrm.cloud.google.com ComputeNetworkFirewallPolicyRule resource.
type ComputeNetworkFirewallPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkFirewallPolicyRuleSpec   `json:"spec,omitempty"`
	Status ComputeNetworkFirewallPolicyRuleStatus `json:"status,omitempty"`
}

// ComputeNetworkFirewallPolicyRuleSpec defines the desired state of ComputeNetworkFirewallPolicyRule.
type ComputeNetworkFirewallPolicyRuleSpec struct {
	// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
	// An optional description for this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS.
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`
	// The firewall policy of the resource.
	FirewallPolicyRef map[string]interface{} `json:"firewallPolicyRef,omitempty" yaml:"firewallPolicyRef,omitempty"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match map[string]interface{} `json:"match,omitempty" yaml:"match,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The priority of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// An optional name for the rule. This field is not a unique identifier and can be updated.
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`
	// A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the target_secure_tag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
	TargetSecureTags []map[string]interface{} `json:"targetSecureTags,omitempty" yaml:"targetSecureTags,omitempty"`
	TargetServiceAccountRefs []map[string]interface{} `json:"targetServiceAccountRefs,omitempty" yaml:"targetServiceAccountRefs,omitempty"`
}

// ComputeNetworkFirewallPolicyRuleStatus defines the observed state of ComputeNetworkFirewallPolicyRule.
type ComputeNetworkFirewallPolicyRuleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
