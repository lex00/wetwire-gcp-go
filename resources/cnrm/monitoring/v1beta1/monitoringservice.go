package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MonitoringService represents a monitoring.cnrm.cloud.google.com MonitoringService resource.
type MonitoringService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringServiceSpec   `json:"spec,omitempty"`
	Status MonitoringServiceStatus `json:"status,omitempty"`
}

// MonitoringServiceSpec defines the desired state of MonitoringService.
type MonitoringServiceSpec struct {
	// Name used for UI elements listing this Service.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Configuration for how to query telemetry on a Service.
	Telemetry map[string]interface{} `json:"telemetry,omitempty" yaml:"telemetry,omitempty"`
}

// MonitoringServiceStatus defines the observed state of MonitoringService.
type MonitoringServiceStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
