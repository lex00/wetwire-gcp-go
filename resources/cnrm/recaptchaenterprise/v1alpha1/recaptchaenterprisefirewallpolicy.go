package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReCAPTCHAEnterpriseFirewallPolicy is the Schema for the ReCAPTCHAEnterpriseFirewallPolicy API
type ReCAPTCHAEnterpriseFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReCAPTCHAEnterpriseFirewallPolicySpec   `json:"spec,omitempty"`
	Status ReCAPTCHAEnterpriseFirewallPolicyStatus `json:"status,omitempty"`
}

// ReCAPTCHAEnterpriseFirewallPolicySpec defines the desired state of ReCAPTCHAEnterpriseFirewallPolicy.
type ReCAPTCHAEnterpriseFirewallPolicySpec struct {
	// Optional. The actions that the caller should take regarding user access. There should be at most one terminal action. A terminal action is any action that forces a response, such as `AllowAction`, `BlockAction` or `SubstituteAction`. Zero or more non-terminal actions such as `SetHeader` might be specified. A single policy can contain up to 16 actions.
	Actions []map[string]interface{} `json:"actions,omitempty" yaml:"actions,omitempty"`
	// Optional. A CEL (Common Expression Language) conditional expression that specifies if this policy applies to an incoming user request. If this condition evaluates to true and the requested path matched the path pattern, the associated actions should be executed by the caller. The condition string is checked for CEL syntax correctness on creation. For more information, see the [CEL spec](https://github.com/google/cel-spec) and its [language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md). A condition has a max length of 500 characters.
	Condition string `json:"condition,omitempty" yaml:"condition,omitempty"`
	// Optional. A description of what this policy aims to achieve, for convenience purposes. The description can at most include 256 UTF-8 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. The path for which this policy applies, specified as a glob pattern. For more information on glob, see the [manual page](https://man7.org/linux/man-pages/man7/glob.7.html). A path has a max length of 200 characters.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The ReCAPTCHAEnterpriseFirewallPolicy name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ReCAPTCHAEnterpriseFirewallPolicyStatus defines the observed state of ReCAPTCHAEnterpriseFirewallPolicy.
type ReCAPTCHAEnterpriseFirewallPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
