package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMWorkloadIdentityPool represents a iam.cnrm.cloud.google.com IAMWorkloadIdentityPool resource.
type IAMWorkloadIdentityPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMWorkloadIdentityPoolSpec   `json:"spec,omitempty"`
	Status IAMWorkloadIdentityPoolStatus `json:"status,omitempty"`
}

// IAMWorkloadIdentityPoolSpec defines the desired state of IAMWorkloadIdentityPool.
type IAMWorkloadIdentityPoolSpec struct {
	// A description of the pool. Cannot exceed 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// A display name for the pool. Cannot exceed 32 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// IAMWorkloadIdentityPoolStatus defines the observed state of IAMWorkloadIdentityPool.
type IAMWorkloadIdentityPoolStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
