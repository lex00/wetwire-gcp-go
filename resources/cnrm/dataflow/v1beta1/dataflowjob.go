package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataflowJob represents a dataflow.cnrm.cloud.google.com DataflowJob resource.
type DataflowJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataflowJobSpec   `json:"spec,omitempty"`
	Status DataflowJobStatus `json:"status,omitempty"`
}

// DataflowJobSpec defines the desired state of DataflowJob.
type DataflowJobSpec struct {
	// List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].
	AdditionalExperiments []string `json:"additionalExperiments,omitempty" yaml:"additionalExperiments,omitempty"`
	// Indicates if the job should use the streaming engine feature.
	EnableStreamingEngine bool `json:"enableStreamingEngine,omitempty" yaml:"enableStreamingEngine,omitempty"`
	// The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	IPConfiguration string `json:"ipConfiguration,omitempty" yaml:"ipConfiguration,omitempty"`
	// The name for the Cloud KMS key for the job.
	KmsKeyRef map[string]interface{} `json:"kmsKeyRef,omitempty" yaml:"kmsKeyRef,omitempty"`
	// The machine type to use for the job.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`
	// Immutable. The number of workers permitted to work on the job. More workers may improve processing speed at additional cost.
	MaxWorkers int32 `json:"maxWorkers,omitempty" yaml:"maxWorkers,omitempty"`
	NetworkRef map[string]interface{} `json:"networkRef,omitempty" yaml:"networkRef,omitempty"`
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	// Immutable. The region in which the created job should run.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID string `json:"resourceID,omitempty" yaml:"resourceID,omitempty"`
	ServiceAccountRef map[string]interface{} `json:"serviceAccountRef,omitempty" yaml:"serviceAccountRef,omitempty"`
	SubnetworkRef map[string]interface{} `json:"subnetworkRef,omitempty" yaml:"subnetworkRef,omitempty"`
	// A writeable location on Google Cloud Storage for the Dataflow job to dump its temporary data.
	TempGcsLocation string `json:"tempGcsLocation,omitempty" yaml:"tempGcsLocation,omitempty"`
	// The Google Cloud Storage path to the Dataflow job template.
	TemplateGcsPath string `json:"templateGcsPath,omitempty" yaml:"templateGcsPath,omitempty"`
	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.
	TransformNameMapping map[string]interface{} `json:"transformNameMapping,omitempty" yaml:"transformNameMapping,omitempty"`
	// Immutable. The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

// DataflowJobStatus defines the observed state of DataflowJob.
type DataflowJobStatus struct {
	// Conditions represent the latest available observations of the resource's state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
