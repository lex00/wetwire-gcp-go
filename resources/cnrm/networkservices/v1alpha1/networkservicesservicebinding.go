package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkServicesServiceBinding is the Schema for the NetworkServicesServiceBinding API
type NetworkServicesServiceBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesServiceBindingSpec   `json:"spec,omitempty"`
	Status NetworkServicesServiceBindingStatus `json:"status,omitempty"`
}

// NetworkServicesServiceBindingSpec defines the desired state of NetworkServicesServiceBinding.
type NetworkServicesServiceBindingSpec struct {
	// Optional. A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Optional. Set of label tags associated with the ServiceBinding resource.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Required. The location of the application.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. The host project of the application.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The NetworkServicesServiceBinding name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Required. The full service directory service name of the format /projects/*/locations/*/namespaces/*/services/*
	ServiceRef map[string]interface{} `json:"serviceRef,omitempty" yaml:"serviceRef,omitempty"`
}

// NetworkServicesServiceBindingStatus defines the observed state of NetworkServicesServiceBinding.
type NetworkServicesServiceBindingStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
