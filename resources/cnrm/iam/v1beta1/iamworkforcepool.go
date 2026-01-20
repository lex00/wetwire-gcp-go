package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMWorkforcePool represents a iam.cnrm.cloud.google.com IAMWorkforcePool resource.
type IAMWorkforcePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMWorkforcePoolSpec   `json:"spec,omitempty"`
	Status IAMWorkforcePoolStatus `json:"status,omitempty"`
}

// IAMWorkforcePoolSpec defines the desired state of IAMWorkforcePool.
type IAMWorkforcePoolSpec struct {
	// A user-specified description of the pool. Cannot exceed 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// A user-specified display name of the pool in Google Cloud Console. Cannot exceed 32 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Organization that this resource belongs to. Only one of [organizationRef] may be specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// How long the Google Cloud access tokens, console sign-in sessions, and gcloud sign-in sessions from this pool are valid. Must be greater than 15 minutes (900s) and less than 12 hours (43200s). If `session_duration` is not configured, minted credentials will have a default duration of one hour (3600s).
	SessionDuration string `json:"sessionDuration,omitempty" yaml:"sessionDuration,omitempty"`
}

// IAMWorkforcePoolStatus defines the observed state of IAMWorkforcePool.
type IAMWorkforcePoolStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
