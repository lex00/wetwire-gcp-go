package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MonitoringDashboard is the Schema for the monitoring API
type MonitoringDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringDashboardSpec   `json:"spec,omitempty"`
	Status MonitoringDashboardStatus `json:"status,omitempty"`
}

// MonitoringDashboardSpec defines the desired state of MonitoringDashboard.
type MonitoringDashboardSpec struct {
	// The content is divided into equally spaced columns and the widgets are arranged vertically.
	ColumnLayout map[string]interface{} `json:"columnLayout,omitempty" yaml:"columnLayout,omitempty"`
	// Filters to reduce the amount of data charted based on the filter criteria.
	DashboardFilters []map[string]interface{} `json:"dashboardFilters,omitempty" yaml:"dashboardFilters,omitempty"`
	// Required. The mutable, human-readable name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// Content is arranged with a basic layout that re-flows a simple list of informational elements like widgets or tiles.
	GridLayout map[string]interface{} `json:"gridLayout,omitempty" yaml:"gridLayout,omitempty"`
	// The content is arranged as a grid of tiles, with each content widget occupying one or more grid blocks.
	MosaicLayout map[string]interface{} `json:"mosaicLayout,omitempty" yaml:"mosaicLayout,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// The content is divided into equally spaced rows and the widgets are arranged horizontally.
	RowLayout map[string]interface{} `json:"rowLayout,omitempty" yaml:"rowLayout,omitempty"`
}

// MonitoringDashboardStatus defines the observed state of MonitoringDashboard.
type MonitoringDashboardStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
