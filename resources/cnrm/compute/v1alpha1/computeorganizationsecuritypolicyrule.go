package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeOrganizationSecurityPolicyRule represents a compute.cnrm.cloud.google.com ComputeOrganizationSecurityPolicyRule resource.
type ComputeOrganizationSecurityPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeOrganizationSecurityPolicyRuleSpec   `json:"spec,omitempty"`
	Status ComputeOrganizationSecurityPolicyRuleStatus `json:"status,omitempty"`
}

// ComputeOrganizationSecurityPolicyRuleSpec defines the desired state of ComputeOrganizationSecurityPolicyRule.
type ComputeOrganizationSecurityPolicyRuleSpec struct {
	// The Action to perform when the client connection triggers the rule. Can currently be either
	// "allow", "deny" or "goto_next".
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
	// A description of the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The direction in which this rule applies. If unspecified an INGRESS rule is created. Possible values: ["INGRESS", "EGRESS"].
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`
	// Denotes whether to enable logging for a particular rule.
	// If logging is enabled, logs will be exported to the
	// configured export destination in Stackdriver.
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match map[string]interface{} `json:"match,omitempty" yaml:"match,omitempty"`
	// Immutable. The ID of the OrganizationSecurityPolicy this rule applies to.
	PolicyID string `json:"policyId,omitempty" yaml:"policyId,omitempty"`
	// If set to true, the specified action is not enforced.
	Preview bool `json:"preview,omitempty" yaml:"preview,omitempty"`
	// Immutable. Optional. The priority of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A list of network resource URLs to which this rule applies.
	// This field allows you to control which network's VMs get
	// this rule. If this field is left blank, all VMs
	// within the organization will receive the rule.
	TargetResources []string `json:"targetResources,omitempty" yaml:"targetResources,omitempty"`
	// A list of service accounts indicating the sets of
	// instances that are applied with this rule.
	TargetServiceAccounts []string `json:"targetServiceAccounts,omitempty" yaml:"targetServiceAccounts,omitempty"`
}

// ComputeOrganizationSecurityPolicyRuleStatus defines the observed state of ComputeOrganizationSecurityPolicyRule.
type ComputeOrganizationSecurityPolicyRuleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
