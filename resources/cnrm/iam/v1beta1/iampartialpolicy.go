package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMPartialPolicy is the Schema for the iampartialpolicy API
type IAMPartialPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMPartialPolicySpec   `json:"spec,omitempty"`
	Status IAMPartialPolicyStatus `json:"status,omitempty"`
}

// IAMPartialPolicySpec defines the desired state of IAMPartialPolicy.
type IAMPartialPolicySpec struct {
	// Optional. The list of IAM bindings managed by Config Connector.
	Bindings []map[string]interface{} `json:"bindings,omitempty" yaml:"bindings,omitempty"`
	// Immutable. Required. The GCP resource to set the IAM policy on (e.g. organization, project...)
	ResourceRef map[string]interface{} `json:"resourceRef,omitempty" yaml:"resourceRef,omitempty"`
}

// IAMPartialPolicyStatus defines the observed state of IAMPartialPolicy.
type IAMPartialPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
