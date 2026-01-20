package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AccessContextManagerAccessLevel is the Schema for the AccessContextManagerAccessLevel API
type AccessContextManagerAccessLevel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerAccessLevelSpec   `json:"spec,omitempty"`
	Status AccessContextManagerAccessLevelStatus `json:"status,omitempty"`
}

// AccessContextManagerAccessLevelSpec defines the desired state of AccessContextManagerAccessLevel.
type AccessContextManagerAccessLevelSpec struct {
	// The AccessPolicy that this resource belongs to.
	AccessPolicyRef map[string]interface{} `json:"accessPolicyRef,omitempty" yaml:"accessPolicyRef,omitempty"`
	// A `BasicLevel` composed of `Conditions`.
	Basic map[string]interface{} `json:"basic,omitempty" yaml:"basic,omitempty"`
	// A `CustomLevel` written in the Common Expression Language.
	Custom map[string]interface{} `json:"custom,omitempty" yaml:"custom,omitempty"`
	// Description of the `AccessLevel` and its use. Does not affect behavior.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// The AccessContextManagerAccessLevel name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Human readable title. Must be unique within the Policy.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}

// AccessContextManagerAccessLevelStatus defines the observed state of AccessContextManagerAccessLevel.
type AccessContextManagerAccessLevelStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
