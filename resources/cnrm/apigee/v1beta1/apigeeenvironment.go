package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ApigeeEnvironment is the Schema for the ApigeeEnvironment API
type ApigeeEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApigeeEnvironmentSpec   `json:"spec,omitempty"`
	Status ApigeeEnvironmentStatus `json:"status,omitempty"`
}

// ApigeeEnvironmentSpec defines the desired state of ApigeeEnvironment.
type ApigeeEnvironmentSpec struct {
	// Reference to parent Apigee Organization.
	ApigeeOrganizationRef map[string]interface{} `json:"apigeeOrganizationRef,omitempty" yaml:"apigeeOrganizationRef,omitempty"`
	// Optional. Description of the environment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Display name for this environment.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Optional. Key-value pairs that may be used for customizing the environment.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`
	// The ApigeeEnvironment name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// ApigeeEnvironmentStatus defines the observed state of ApigeeEnvironment.
type ApigeeEnvironmentStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
