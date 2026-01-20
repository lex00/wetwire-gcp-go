package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppHubApplication is the Schema for the AppHubApplication API
type AppHubApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppHubApplicationSpec   `json:"spec,omitempty"`
	Status AppHubApplicationStatus `json:"status,omitempty"`
}

// AppHubApplicationSpec defines the desired state of AppHubApplication.
type AppHubApplicationSpec struct {
	// Optional. Consumer provided attributes.
	Attributes map[string]interface{} `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	// Optional. User-defined description of an Application. Can have a maximum length of 2048 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. User-defined name for the Application. Can have a maximum length of 63 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Required. The location of the application.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. The host project of the application.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The AppHubApplication name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. Immutable. Defines what data can be included into this Application. Limits which Services and Workloads can be registered.
	Scope map[string]interface{} `json:"scope,omitempty" yaml:"scope,omitempty"`
}

// AppHubApplicationStatus defines the observed state of AppHubApplication.
type AppHubApplicationStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
