package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeployCustomTargetType is the Schema for the DeployCustomTargetType API
type DeployCustomTargetType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeployCustomTargetTypeSpec   `json:"spec,omitempty"`
	Status DeployCustomTargetTypeStatus `json:"status,omitempty"`
}

// DeployCustomTargetTypeSpec defines the desired state of DeployCustomTargetType.
type DeployCustomTargetTypeSpec struct {
	// Configures render and deploy for the `CustomTargetType` using Skaffold custom actions.
	CustomActions map[string]interface{} `json:"customActions,omitempty" yaml:"customActions,omitempty"`
	// Optional. Description of the `CustomTargetType`. Max length is 255 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The DeployCustomTargetType name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// DeployCustomTargetTypeStatus defines the observed state of DeployCustomTargetType.
type DeployCustomTargetTypeStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
