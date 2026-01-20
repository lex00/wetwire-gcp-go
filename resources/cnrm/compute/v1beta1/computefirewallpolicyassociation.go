package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeFirewallPolicyAssociation represents a compute.cnrm.cloud.google.com ComputeFirewallPolicyAssociation resource.
type ComputeFirewallPolicyAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeFirewallPolicyAssociationSpec   `json:"spec,omitempty"`
	Status ComputeFirewallPolicyAssociationStatus `json:"status,omitempty"`
}

// ComputeFirewallPolicyAssociationSpec defines the desired state of ComputeFirewallPolicyAssociation.
type ComputeFirewallPolicyAssociationSpec struct {
	// Immutable.
	AttachmentTargetRef map[string]interface{} `json:"attachmentTargetRef,omitempty" yaml:"attachmentTargetRef,omitempty"`
	// Immutable.
	FirewallPolicyRef map[string]interface{} `json:"firewallPolicyRef,omitempty" yaml:"firewallPolicyRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeFirewallPolicyAssociationStatus defines the observed state of ComputeFirewallPolicyAssociation.
type ComputeFirewallPolicyAssociationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
