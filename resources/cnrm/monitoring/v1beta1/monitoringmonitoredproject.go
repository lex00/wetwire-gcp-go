package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MonitoringMonitoredProject represents a monitoring.cnrm.cloud.google.com MonitoringMonitoredProject resource.
type MonitoringMonitoredProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringMonitoredProjectSpec   `json:"spec,omitempty"`
	Status MonitoringMonitoredProjectStatus `json:"status,omitempty"`
}

// MonitoringMonitoredProjectSpec defines the desired state of MonitoringMonitoredProject.
type MonitoringMonitoredProjectSpec struct {
	// Immutable. Required. The resource name of the existing Metrics Scope that will monitor this project. Example: locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}
	MetricsScope string `json:"metricsScope,omitempty" yaml:"metricsScope,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// MonitoringMonitoredProjectStatus defines the observed state of MonitoringMonitoredProject.
type MonitoringMonitoredProjectStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
