package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryJob represents a bigquery.cnrm.cloud.google.com BigQueryJob resource.
type BigQueryJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryJobSpec   `json:"spec,omitempty"`
	Status BigQueryJobStatus `json:"status,omitempty"`
}

// BigQueryJobSpec defines the desired state of BigQueryJob.
type BigQueryJobSpec struct {
	// Immutable. Copies a table.
	Copy map[string]interface{} `json:"copy,omitempty" yaml:"copy,omitempty"`
	// Immutable. Configures an extract job.
	Extract map[string]interface{} `json:"extract,omitempty" yaml:"extract,omitempty"`
	// Immutable. Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs string `json:"jobTimeoutMs,omitempty" yaml:"jobTimeoutMs,omitempty"`
	// Immutable. Configures a load job.
	Load map[string]interface{} `json:"load,omitempty" yaml:"load,omitempty"`
	// Immutable. The geographic location of the job. The default value is US.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Immutable. Configures a query job.
	Query map[string]interface{} `json:"query,omitempty" yaml:"query,omitempty"`
	// Immutable. Optional. The jobId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
}

// BigQueryJobStatus defines the observed state of BigQueryJob.
type BigQueryJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
