package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MonitoringServiceLevelObjective represents a monitoring.cnrm.cloud.google.com MonitoringServiceLevelObjective resource.
type MonitoringServiceLevelObjective struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringServiceLevelObjectiveSpec   `json:"spec,omitempty"`
	Status MonitoringServiceLevelObjectiveStatus `json:"status,omitempty"`
}

// MonitoringServiceLevelObjectiveSpec defines the desired state of MonitoringServiceLevelObjective.
type MonitoringServiceLevelObjectiveSpec struct {
	// A calendar period, semantically "since the start of the current ``". At this time, only `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH` are supported. Possible values: CALENDAR_PERIOD_UNSPECIFIED, DAY, WEEK, FORTNIGHT, MONTH, QUARTER, HALF, YEAR
	CalendarPeriod string `json:"calendarPeriod,omitempty" yaml:"calendarPeriod,omitempty"`
	// Name used for UI elements listing this SLO.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	// The fraction of service that must be good in order for this objective to be met. `0 < goal <= 0.999`.
	Goal float64 `json:"goal,omitempty" yaml:"goal,omitempty"`
	// Immutable. The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// A rolling time period, semantically "in the past ``". Must be an integer multiple of 1 day no larger than 30 days.
	RollingPeriod string `json:"rollingPeriod,omitempty" yaml:"rollingPeriod,omitempty"`
	// The definition of good service, used to measure and calculate the quality of the `Service`'s performance with respect to a single aspect of service quality.
	ServiceLevelIndicator map[string]interface{} `json:"serviceLevelIndicator,omitempty" yaml:"serviceLevelIndicator,omitempty"`
	// Immutable.
	ServiceRef map[string]interface{} `json:"serviceRef,omitempty" yaml:"serviceRef,omitempty"`
}

// MonitoringServiceLevelObjectiveStatus defines the observed state of MonitoringServiceLevelObjective.
type MonitoringServiceLevelObjectiveStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
