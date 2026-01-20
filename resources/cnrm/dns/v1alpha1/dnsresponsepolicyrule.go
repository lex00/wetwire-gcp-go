package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DNSResponsePolicyRule represents a dns.cnrm.cloud.google.com DNSResponsePolicyRule resource.
type DNSResponsePolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSResponsePolicyRuleSpec   `json:"spec,omitempty"`
	Status DNSResponsePolicyRuleStatus `json:"status,omitempty"`
}

// DNSResponsePolicyRuleSpec defines the desired state of DNSResponsePolicyRule.
type DNSResponsePolicyRuleSpec struct {
	// Answer this query with a behavior rather than DNS data. Acceptable values are 'behaviorUnspecified', and 'bypassResponsePolicy'.
	Behavior string `json:"behavior,omitempty" yaml:"behavior,omitempty"`
	// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
	DNSName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`
	// Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
	// in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
	LocalData map[string]interface{} `json:"localData,omitempty" yaml:"localData,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The ruleName of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Identifies the response policy addressed by this request.
	ResponsePolicy string `json:"responsePolicy,omitempty" yaml:"responsePolicy,omitempty"`
}

// DNSResponsePolicyRuleStatus defines the observed state of DNSResponsePolicyRule.
type DNSResponsePolicyRuleStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
