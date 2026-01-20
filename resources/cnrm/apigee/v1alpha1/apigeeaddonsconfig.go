package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApigeeAddonsConfig represents a apigee.cnrm.cloud.google.com ApigeeAddonsConfig resource.
type ApigeeAddonsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeAddonsConfigSpec   `json:"spec,omitempty"`
	Status ApigeeAddonsConfigStatus `json:"status,omitempty"`
}

// ApigeeAddonsConfigSpec defines the desired state of ApigeeAddonsConfig.
type ApigeeAddonsConfigSpec struct {
	// Addon configurations of the Apigee organization.
	AddonsConfig map[string]interface{} `json:"addonsConfig,omitempty" yaml:"addonsConfig,omitempty"`
	// Immutable. Name of the Apigee organization.
	Org string `json:"org,omitempty" yaml:"org,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ApigeeAddonsConfigStatus defines the observed state of ApigeeAddonsConfig.
type ApigeeAddonsConfigStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
