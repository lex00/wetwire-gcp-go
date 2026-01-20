package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeOrganizationSecurityPolicyAssociation represents a compute.cnrm.cloud.google.com ComputeOrganizationSecurityPolicyAssociation resource.
type ComputeOrganizationSecurityPolicyAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeOrganizationSecurityPolicyAssociationSpec   `json:"spec,omitempty"`
	Status ComputeOrganizationSecurityPolicyAssociationStatus `json:"status,omitempty"`
}

// ComputeOrganizationSecurityPolicyAssociationSpec defines the desired state of ComputeOrganizationSecurityPolicyAssociation.
type ComputeOrganizationSecurityPolicyAssociationSpec struct {
	// Immutable. The resource that the security policy is attached to.
	AttachmentID string `json:"attachmentId,omitempty" yaml:"attachmentId,omitempty"`
	// Immutable. The security policy ID of the association.
	PolicyID string `json:"policyId,omitempty" yaml:"policyId,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeOrganizationSecurityPolicyAssociationStatus defines the observed state of ComputeOrganizationSecurityPolicyAssociation.
type ComputeOrganizationSecurityPolicyAssociationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
