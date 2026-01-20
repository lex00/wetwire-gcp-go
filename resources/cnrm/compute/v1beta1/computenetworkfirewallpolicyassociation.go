package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputeNetworkFirewallPolicyAssociation represents a compute.cnrm.cloud.google.com ComputeNetworkFirewallPolicyAssociation resource.
type ComputeNetworkFirewallPolicyAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeNetworkFirewallPolicyAssociationSpec   `json:"spec,omitempty"`
	Status ComputeNetworkFirewallPolicyAssociationStatus `json:"status,omitempty"`
}

// ComputeNetworkFirewallPolicyAssociationSpec defines the desired state of ComputeNetworkFirewallPolicyAssociation.
type ComputeNetworkFirewallPolicyAssociationSpec struct {
	// The target that the firewall policy is attached to.
	AttachmentTargetRef map[string]interface{} `json:"attachmentTargetRef,omitempty" yaml:"attachmentTargetRef,omitempty"`
	// The firewall policy ID of the association.
	FirewallPolicyRef map[string]interface{} `json:"firewallPolicyRef,omitempty" yaml:"firewallPolicyRef,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ComputeNetworkFirewallPolicyAssociationStatus defines the observed state of ComputeNetworkFirewallPolicyAssociation.
type ComputeNetworkFirewallPolicyAssociationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
