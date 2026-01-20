package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeOrganizationSecurityPolicy represents a compute.cnrm.cloud.google.com ComputeOrganizationSecurityPolicy resource.
type ComputeOrganizationSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeOrganizationSecurityPolicySpec   `json:"spec,omitempty"`
	Status ComputeOrganizationSecurityPolicyStatus `json:"status,omitempty"`
}

// ComputeOrganizationSecurityPolicySpec defines the desired state of ComputeOrganizationSecurityPolicy.
type ComputeOrganizationSecurityPolicySpec struct {
	// A textual description for the organization security policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Immutable. A textual name of the security policy.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id} or folders/{folder_id}.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
	// Immutable. Optional. The policyId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Immutable. The type indicates the intended use of the security policy.
	// For organization security policies, the only supported type
	// is "FIREWALL". Default value: "FIREWALL" Possible values: ["FIREWALL"].
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// ComputeOrganizationSecurityPolicyStatus defines the observed state of ComputeOrganizationSecurityPolicy.
type ComputeOrganizationSecurityPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
