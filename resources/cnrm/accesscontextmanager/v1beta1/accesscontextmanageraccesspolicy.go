package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AccessContextManagerAccessPolicy is the Schema for the AccessContextManagerAccessPolicy API As per https://cloud.google.com/config-connector/docs/reference/resource-docs/accesscontextmanager/accesscontextmanageraccesspolicy#annotations the parent is organization which is stored in the cnrm.cloud.google.com/organization-id annotation.
type AccessContextManagerAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerAccessPolicySpec   `json:"spec,omitempty"`
	Status AccessContextManagerAccessPolicyStatus `json:"status,omitempty"`
}

// AccessContextManagerAccessPolicySpec defines the desired state of AccessContextManagerAccessPolicy.
type AccessContextManagerAccessPolicySpec struct {
	// The AccessContextManagerAccessPolicy name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Human readable title. Does not affect behavior.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}

// AccessContextManagerAccessPolicyStatus defines the observed state of AccessContextManagerAccessPolicy.
type AccessContextManagerAccessPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
