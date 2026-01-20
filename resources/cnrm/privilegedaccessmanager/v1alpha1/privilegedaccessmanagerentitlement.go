package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrivilegedAccessManagerEntitlement is the Schema for the PrivilegedAccessManagerEntitlement API.
type PrivilegedAccessManagerEntitlement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivilegedAccessManagerEntitlementSpec   `json:"spec,omitempty"`
	Status PrivilegedAccessManagerEntitlementStatus `json:"status,omitempty"`
}

// PrivilegedAccessManagerEntitlementSpec defines the desired state of PrivilegedAccessManagerEntitlement.
type PrivilegedAccessManagerEntitlementSpec struct {
	// Optional. Additional email addresses to be notified based on actions taken.
	AdditionalNotificationTargets map[string]interface{} `json:"additionalNotificationTargets,omitempty" yaml:"additionalNotificationTargets,omitempty"`
	// Optional. The approvals needed before access are granted to a requester. No approvals are needed if this field is null.
	ApprovalWorkflow map[string]interface{} `json:"approvalWorkflow,omitempty" yaml:"approvalWorkflow,omitempty"`
	// Who can create grants using this entitlement. This list should contain at most one entry.
	EligibleUsers []map[string]interface{} `json:"eligibleUsers,omitempty" yaml:"eligibleUsers,omitempty"`
	// Immutable. The Folder that this resource belongs to. One and only one of 'projectRef', 'folderRef', or 'organizationRef' must be set.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Immutable. Location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. The maximum amount of time that access is granted for a request. A requester can ask for a duration less than this, but never more.
	MaxRequestDuration string `json:"maxRequestDuration,omitempty" yaml:"maxRequestDuration,omitempty"`
	// Immutable. The Organization that this resource belongs to. One and only one of 'projectRef', 'folderRef', or 'organizationRef' must be set.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The access granted to a requester on successful approval.
	PrivilegedAccess map[string]interface{} `json:"privilegedAccess,omitempty" yaml:"privilegedAccess,omitempty"`
	// Immutable. The Project that this resource belongs to. One and only one of 'projectRef', 'folderRef', or 'organizationRef' must be set.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Required. The manner in which the requester should provide a justification for requesting access.
	RequesterJustificationConfig map[string]interface{} `json:"requesterJustificationConfig,omitempty" yaml:"requesterJustificationConfig,omitempty"`
	// The PrivilegedAccessManagerEntitlement name. If not given, the 'metadata.name' will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// PrivilegedAccessManagerEntitlementStatus defines the observed state of PrivilegedAccessManagerEntitlement.
type PrivilegedAccessManagerEntitlementStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
