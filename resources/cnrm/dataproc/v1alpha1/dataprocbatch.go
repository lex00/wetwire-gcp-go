package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataprocBatch is the Schema for the DataprocBatch API
type DataprocBatch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataprocBatchSpec   `json:"spec,omitempty"`
	Status DataprocBatchStatus `json:"status,omitempty"`
}

// DataprocBatchSpec defines the desired state of DataprocBatch.
type DataprocBatchSpec struct {
	// Optional. Environment configuration for the batch execution.
	EnvironmentConfig map[string]interface{} `json:"environmentConfig,omitempty" yaml:"environmentConfig,omitempty"`
	// Optional. The labels to associate with this batch. Label **keys** must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a batch.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Required.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
	// Required.
	ProjectRef map[string]interface{} `json:"projectRef,omitempty" yaml:"projectRef,omitempty"`
	// Optional. PySpark batch config.
	PysparkBatch map[string]interface{} `json:"pysparkBatch,omitempty" yaml:"pysparkBatch,omitempty"`
	// The DataprocBatch name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Runtime configuration for the batch execution.
	RuntimeConfig map[string]interface{} `json:"runtimeConfig,omitempty" yaml:"runtimeConfig,omitempty"`
	// Optional. Spark batch config.
	SparkBatch map[string]interface{} `json:"sparkBatch,omitempty" yaml:"sparkBatch,omitempty"`
	// Optional. SparkR batch config.
	SparkRBatch map[string]interface{} `json:"sparkRBatch,omitempty" yaml:"sparkRBatch,omitempty"`
	// Optional. SparkSql batch config.
	SparkSQLBatch map[string]interface{} `json:"sparkSQLBatch,omitempty" yaml:"sparkSQLBatch,omitempty"`
}

// DataprocBatchStatus defines the observed state of DataprocBatch.
type DataprocBatchStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
