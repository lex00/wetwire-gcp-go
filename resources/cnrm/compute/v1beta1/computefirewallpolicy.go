package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeFirewallPolicy represents a compute.cnrm.cloud.google.com ComputeFirewallPolicy resource.
type ComputeFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeFirewallPolicySpec   `json:"spec,omitempty"`
	Status ComputeFirewallPolicyStatus `json:"status,omitempty"`
}

// ComputeFirewallPolicySpec defines the desired state of ComputeFirewallPolicy.
type ComputeFirewallPolicySpec struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. The Folder that this resource belongs to. Only one of [folderRef, organizationRef] may be specified.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Immutable. The Organization that this resource belongs to. Only one of [folderRef, organizationRef] may be specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName string `json:"shortName,omitempty" yaml:"shortName,omitempty"`
}

// ComputeFirewallPolicyStatus defines the observed state of ComputeFirewallPolicy.
type ComputeFirewallPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
