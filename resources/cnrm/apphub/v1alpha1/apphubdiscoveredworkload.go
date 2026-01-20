package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppHubDiscoveredWorkload is the Schema for the AppHubDiscoveredWorkload API
type AppHubDiscoveredWorkload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppHubDiscoveredWorkloadSpec   `json:"spec,omitempty"`
	Status AppHubDiscoveredWorkloadStatus `json:"status,omitempty"`
}

// AppHubDiscoveredWorkloadSpec defines the desired state of AppHubDiscoveredWorkload.
type AppHubDiscoveredWorkloadSpec struct {
	// Required. The location of the application.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required. The host project of the application.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// The AppHubDiscoveredWorkload name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// AppHubDiscoveredWorkloadStatus defines the observed state of AppHubDiscoveredWorkload.
type AppHubDiscoveredWorkloadStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
