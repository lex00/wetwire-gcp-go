package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BigQueryConnectionConnection is the Schema for the BigQueryConnectionConnection API
type BigQueryConnectionConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryConnectionConnectionSpec   `json:"spec,omitempty"`
	Status BigQueryConnectionConnectionStatus `json:"status,omitempty"`
}

// BigQueryConnectionConnectionSpec defines the desired state of BigQueryConnectionConnection.
type BigQueryConnectionConnectionSpec struct {
	// Amazon Web Services (AWS) properties.
	Aws map[string]interface{} `json:"aws,omitempty" yaml:"aws,omitempty"`
	// Azure properties.
	Azure map[string]interface{} `json:"azure,omitempty" yaml:"azure,omitempty"`
	// Use Cloud Resource properties.
	CloudResource map[string]interface{} `json:"cloudResource,omitempty" yaml:"cloudResource,omitempty"`
	// Cloud SQL properties.
	CloudSQL map[string]interface{} `json:"cloudSQL,omitempty" yaml:"cloudSQL,omitempty"`
	// Cloud Spanner properties.
	CloudSpanner map[string]interface{} `json:"cloudSpanner,omitempty" yaml:"cloudSpanner,omitempty"`
	// User provided description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// User provided display name for the connection.
	FriendlyName string `json:"friendlyName,omitempty" yaml:"friendlyName,omitempty"`
	// Immutable.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// The Project that this resource belongs to.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Immutable. Optional. The BigQuery Connection ID used for resource creation or acquisition. For creation: If specified, this value is used as the connection ID. If not provided, a UUID will be generated and assigned as the connection ID. For acquisition: This field must be provided to identify the connection resource to acquire.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Spark properties.
	Spark map[string]interface{} `json:"spark,omitempty" yaml:"spark,omitempty"`
}

// BigQueryConnectionConnectionStatus defines the observed state of BigQueryConnectionConnection.
type BigQueryConnectionConnectionStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
