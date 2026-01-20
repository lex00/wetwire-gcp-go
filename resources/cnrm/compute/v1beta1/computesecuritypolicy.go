package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeSecurityPolicy represents a compute.cnrm.cloud.google.com ComputeSecurityPolicy resource.
type ComputeSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeSecurityPolicySpec   `json:"spec,omitempty"`
	Status ComputeSecurityPolicyStatus `json:"status,omitempty"`
}

// ComputeSecurityPolicySpec defines the desired state of ComputeSecurityPolicy.
type ComputeSecurityPolicySpec struct {
	// Adaptive Protection Config of this security policy.
	AdaptiveProtectionConfig map[string]interface{} `json:"adaptiveProtectionConfig,omitempty" yaml:"adaptiveProtectionConfig,omitempty"`
	// Advanced Options Config of this security policy.
	AdvancedOptionsConfig map[string]interface{} `json:"advancedOptionsConfig,omitempty" yaml:"advancedOptionsConfig,omitempty"`
	// An optional description of this security policy. Max size is 2048.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// reCAPTCHA configuration options to be applied for the security policy.
	RecaptchaOptionsConfig map[string]interface{} `json:"recaptchaOptionsConfig,omitempty" yaml:"recaptchaOptionsConfig,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The set of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rule []map[string]interface{} `json:"rule,omitempty" yaml:"rule,omitempty"`
	// The type indicates the intended use of the security policy. CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// ComputeSecurityPolicyStatus defines the observed state of ComputeSecurityPolicy.
type ComputeSecurityPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
