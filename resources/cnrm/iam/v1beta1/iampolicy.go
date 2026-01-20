package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IAMPolicy is the Schema for the iampolicies API
type IAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMPolicySpec   `json:"spec,omitempty"`
	Status IAMPolicyStatus `json:"status,omitempty"`
}

// IAMPolicySpec defines the desired state of IAMPolicy.
type IAMPolicySpec struct {
	// Optional. The list of IAM audit configs.
	AuditConfigs []map[string]interface{} `json:"auditConfigs,omitempty" yaml:"auditConfigs,omitempty"`
	// Optional. The list of IAM bindings.
	Bindings []map[string]interface{} `json:"bindings,omitempty" yaml:"bindings,omitempty"`
	// Immutable. Required. The GCP resource to set the IAM policy on (e.g. organization, project...)
	ResourceRef map[string]interface{} `json:"resourceRef,omitempty" yaml:"resourceRef,omitempty"`
}

// IAMPolicyStatus defines the observed state of IAMPolicy.
type IAMPolicyStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
