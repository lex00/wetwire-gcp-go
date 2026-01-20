package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LoggingLogSink represents a logging.cnrm.cloud.google.com LoggingLogSink resource.
type LoggingLogSink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingLogSinkSpec   `json:"spec,omitempty"`
	Status LoggingLogSinkStatus `json:"status,omitempty"`
}

// LoggingLogSinkSpec defines the desired state of LoggingLogSink.
type LoggingLogSinkSpec struct {
	// Options that affect sinks exporting data to BigQuery.
	BigqueryOptions map[string]interface{} `json:"bigqueryOptions,omitempty" yaml:"bigqueryOptions,omitempty"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Destination map[string]interface{} `json:"destination,omitempty" yaml:"destination,omitempty"`
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion's filters, it will not be exported.
	Exclusions []map[string]interface{} `json:"exclusions,omitempty" yaml:"exclusions,omitempty"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
	// The folder in which to create the sink. Only one of projectRef,
	// folderRef, or organizationRef may be specified.
	FolderRef map[string]interface{} `json:"folderRef,omitempty" yaml:"folderRef,omitempty"`
	// Immutable. Whether or not to include children organizations in the sink export. If true, logs associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	IncludeChildren bool `json:"includeChildren,omitempty" yaml:"includeChildren,omitempty"`
	// The organization in which to create the sink. Only one of projectRef,
	// folderRef, or organizationRef may be specified.
	OrganizationRef map[string]interface{} `json:"organizationRef,omitempty" yaml:"organizationRef,omitempty"`
	// The project in which to create the sink. Only one of projectRef,
	// folderRef, or organizationRef may be specified.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Whether or not to create a unique identity associated with this sink. If false (the default), then the writer_identity used is serviceAccount:cloud-logs@system.gserviceaccount.com. If true, then a unique service account is created and used for this sink. If you wish to publish logs across projects, you must set unique_writer_identity to true.
	UniqueWriterIdentity bool `json:"uniqueWriterIdentity,omitempty" yaml:"uniqueWriterIdentity,omitempty"`
}

// LoggingLogSinkStatus defines the observed state of LoggingLogSink.
type LoggingLogSinkStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
