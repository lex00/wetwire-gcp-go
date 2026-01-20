package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataprocJob is the Schema for the DataprocJob API
type DataprocJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataprocJobSpec   `json:"spec,omitempty"`
	Status DataprocJobStatus `json:"status,omitempty"`
}

// DataprocJobSpec defines the desired state of DataprocJob.
type DataprocJobSpec struct {
	// Optional. Driver scheduling configuration.
	DriverSchedulingConfig map[string]interface{} `json:"driverSchedulingConfig,omitempty" yaml:"driverSchedulingConfig,omitempty"`
	// Optional. Job is a Flink job.
	FlinkJob map[string]interface{} `json:"flinkJob,omitempty" yaml:"flinkJob,omitempty"`
	// Optional. Job is a Hadoop job.
	HadoopJob map[string]interface{} `json:"hadoopJob,omitempty" yaml:"hadoopJob,omitempty"`
	// Optional. Job is a Hive job.
	HiveJob map[string]interface{} `json:"hiveJob,omitempty" yaml:"hiveJob,omitempty"`
	// Optional. The labels to associate with this job. Label **keys** must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** can be empty, but, if present, must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a job.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	// Required.
	Parent map[string]interface{} `json:"parent,omitempty" yaml:"parent,omitempty"`
	// Optional. Job is a Pig job.
	PigJob map[string]interface{} `json:"pigJob,omitempty" yaml:"pigJob,omitempty"`
	// Required. Job information, including how, when, and where to run the job.
	Placement map[string]interface{} `json:"placement,omitempty" yaml:"placement,omitempty"`
	// Optional. Job is a Presto job.
	PrestoJob map[string]interface{} `json:"prestoJob,omitempty" yaml:"prestoJob,omitempty"`
	// Optional. Job is a PySpark job.
	PysparkJob map[string]interface{} `json:"pysparkJob,omitempty" yaml:"pysparkJob,omitempty"`
	// Optional. The fully qualified reference to the job, which can be used to obtain the equivalent REST path of the job resource. If this property is not specified when a job is created, the server generates a <code>job_id</code>.
	Reference map[string]interface{} `json:"reference,omitempty" yaml:"reference,omitempty"`
	// The DataprocJob name. If not given, the metadata.name will be used.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	// Optional. Job scheduling configuration.
	Scheduling map[string]interface{} `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`
	// Optional. Job is a Spark job.
	SparkJob map[string]interface{} `json:"sparkJob,omitempty" yaml:"sparkJob,omitempty"`
	// Optional. Job is a SparkR job.
	SparkRJob map[string]interface{} `json:"sparkRJob,omitempty" yaml:"sparkRJob,omitempty"`
	// Optional. Job is a SparkSql job.
	SparkSQLJob map[string]interface{} `json:"sparkSQLJob,omitempty" yaml:"sparkSQLJob,omitempty"`
	// Optional. Job is a Trino job.
	TrinoJob map[string]interface{} `json:"trinoJob,omitempty" yaml:"trinoJob,omitempty"`
}

// DataprocJobStatus defines the observed state of DataprocJob.
type DataprocJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
