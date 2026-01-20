package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppHubDiscoveredService is the Schema for the AppHubDiscoveredService API
type AppHubDiscoveredService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppHubDiscoveredServiceSpec   `json:"spec,omitempty"`
	Status AppHubDiscoveredServiceStatus `json:"status,omitempty"`
}

// AppHubDiscoveredServiceSpec defines the desired state of AppHubDiscoveredService.
type AppHubDiscoveredServiceSpec struct {
	// Required. The location of the application.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. The host project of the application.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The AppHubDiscoveredService name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// AppHubDiscoveredServiceStatus defines the observed state of AppHubDiscoveredService.
type AppHubDiscoveredServiceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
