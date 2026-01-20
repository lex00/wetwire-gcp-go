package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OrgPolicyPolicy is the Schema for the OrgPolicyPolicy API
type OrgPolicyPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrgPolicyPolicySpec   `json:"spec,omitempty"`
	Status OrgPolicyPolicyStatus `json:"status,omitempty"`
}

// OrgPolicyPolicySpec defines the desired state of OrgPolicyPolicy.
type OrgPolicyPolicySpec struct {
	// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it's enforced.
	DryRunSpec map[string]interface{} `json:"dryRunSpec,omitempty" yaml:"dryRunSpec,omitempty"`
	// Immutable. The Folder that this resource belongs to. One and only one of 'projectRef', 'folderRef', or 'organizationRef' must be set.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Immutable. The Organization that this resource belongs to. One and only one of 'projectRef', 'folderRef', or 'organizationRef' must be set.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. The Project that this resource belongs to. One and only one of 'projectRef', 'folderRef', or 'organizationRef' must be set.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The OrgPolicyPolicy name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Basic information about the Organization Policy.
	Spec map[string]interface{} `json:"spec,omitempty" yaml:"spec,omitempty"`
}

// OrgPolicyPolicyStatus defines the observed state of OrgPolicyPolicy.
type OrgPolicyPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
